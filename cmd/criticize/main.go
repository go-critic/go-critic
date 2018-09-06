package criticize

import (
	"encoding/json"
	"flag"
	"go/ast"
	"go/build"
	"go/parser"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/go-critic/go-critic/cmd/internal/flagparser"
	"github.com/go-critic/go-critic/lint"
	"github.com/logrusorgru/aurora"
	"golang.org/x/tools/go/loader"
)

var generatedFileCommentRE = regexp.MustCompile("Code generated .* DO NOT EDIT.")

type linter struct {
	ctx *lint.Context

	prog *loader.Program

	checkers []*lint.Checker

	foundIssues bool // True if there any checker reported an issue

	flags *flagparser.FlagParser

	packages        []string
	enabledCheckers []string
	checkerParams   map[string]map[string]interface{}
}

// parseEnabledCheckers processes enabled checkers, intersect them with disabled, etc.
func (l *linter) parseEnabledCheckers() {
	switch l.flags.Enable {
	case flagparser.EnableAll:
		for _, rule := range lint.RuleList() {
			if rule.Experimental && !l.flags.WithExperimental {
				continue
			}
			if rule.VeryOpinionated && !l.flags.WithOpinionated {
				continue
			}
			l.enabledCheckers = append(l.enabledCheckers, rule.Name())
		}
	case "":
		// Empty slice. Semantically "disable-all".
		// Can be used to run all pipelines without actual checkers.
		l.enabledCheckers = []string{}
	default:
		// Comma-separated list of names.
		l.enabledCheckers = l.flags.EnabledCheckers()
	}

	switch l.flags.Disable {
	case flagparser.DisableAll:
		l.enabledCheckers = l.enabledCheckers[:0]
	case "":
	// nothing to disable, skip
	default:
		filtred := l.enabledCheckers[:0]

		for _, e := range l.enabledCheckers {
			found := false
			for _, d := range l.flags.DisabledCheckers() {
				if e == d {
					found = true
				}
			}
			if !found {
				filtred = append(filtred, e)
			}
		}
		l.enabledCheckers = filtred
	}
}

// Main implements gocritic sub-command entry point.
func Main() {
	var l linter
	parseArgv(&l)
	if l.flags.ConfigFile != "" {
		l.LoadConfig()
	} else {
		l.parseEnabledCheckers()
	}
	l.LoadProgram()
	l.InitCheckers()

	for _, pkgPath := range l.packages {
		l.CheckPackage(pkgPath)
	}

	os.Exit(l.ExitCode())
}

func blame(format string, args ...interface{}) {
	log.Printf(format, args...)
	flag.Usage()
	os.Exit(1)
}

// parseArgv processes command-line arguments and fills ctx argument with them.
// Terminates program on error.
func parseArgv(l *linter) {
	flag.Usage = func() {
		log.Printf("usage: [flags] package...")
		flag.PrintDefaults()
	}

	l.flags = flagparser.NewFlagParser(flag.CommandLine)

	if err := l.flags.Parse(); err != nil {
		blame(err.Error())
	}

	l.packages = l.flags.ParsedArgs()

	if len(l.packages) == 0 {
		blame("no packages specified\n")
	}
}

func (l *linter) LoadConfig() {
	raw, err := ioutil.ReadFile(l.flags.ConfigFile)
	if err != nil {
		log.Printf("cannot read config file %s, got error: %s", l.flags.ConfigFile, err)
		return
	}

	var params map[string]interface{}
	if err := json.Unmarshal(raw, &params); err != nil {
		log.Fatalf("cannot parse config file, got error: %s", err)
		return
	}

	l.enabledCheckers = []string{}
	l.checkerParams = make(map[string]map[string]interface{})
	for k, v := range params {
		l.enabledCheckers = append(l.enabledCheckers, k)
		if v, ok := v.(map[string]interface{}); ok {
			l.checkerParams[k] = v
		} else {
			log.Printf("cannot parse value for %v", k)
		}
	}
}

func (l *linter) LoadProgram() {
	sizes := types.SizesFor("gc", runtime.GOARCH)
	if sizes == nil {
		log.Fatalf("can't find sizes info for %s", runtime.GOARCH)
	}

	conf := loader.Config{
		ParserMode: parser.ParseComments,
		TypeChecker: types.Config{
			Sizes: sizes,
		},
	}

	if _, err := conf.FromArgs(l.packages, true); err != nil {
		log.Fatalf("resolve packages: %v", err)
	}
	prog, err := conf.Load()
	if err != nil {
		log.Fatalf("load program: %v", err)
	}

	l.prog = prog
	l.ctx = lint.NewContext(prog.Fset, sizes)
}

func (l *linter) InitCheckers() {
	requested := make(map[string]bool)
	available := lint.RuleList()

	if l.enabledCheckers == nil {
		// Fill default checkers set.
		for _, rule := range available {
			if rule.Experimental && !l.flags.WithExperimental {
				continue
			}
			if rule.VeryOpinionated && !l.flags.WithOpinionated {
				continue
			}
			requested[rule.Name()] = true
		}
	} else {
		for _, name := range l.enabledCheckers {
			requested[name] = true
		}
	}

	for _, rule := range available {
		if !requested[rule.Name()] {
			continue
		}
		l.checkers = append(l.checkers, lint.NewChecker(
			rule,
			l.ctx,
			l.checkerParams[rule.Name()],
		))
		delete(requested, rule.Name())
	}

	if len(requested) != 0 {
		for name := range requested {
			log.Printf("%s: checker not found", name)
		}
		log.Fatalf("exiting due to initialization failure")
	}
}

func (l *linter) CheckPackage(pkgPath string) {
	pkgInfo := l.prog.Imported[pkgPath]
	if pkgInfo == nil || !pkgInfo.TransitivelyErrorFree {
		log.Fatalf("%s package is not properly loaded", pkgPath)
	}

	l.ctx.SetPackageInfo(&pkgInfo.Info, pkgInfo.Pkg)
	for _, f := range pkgInfo.Files {
		if l.flags.IgnoreTests && l.isTestFile(f) {
			continue
		}
		if !l.flags.CheckGenerated && l.isGenerated(f) {
			continue
		}
		l.ctx.SetFileInfo(l.getFilename(f), f)
		l.checkFile(f)
	}
}

func (l *linter) isTestFile(f *ast.File) bool {
	return strings.HasSuffix(l.getFilename(f), "_test.go")
}

func (l *linter) isGenerated(f *ast.File) bool {
	return len(f.Comments) != 0 && generatedFileCommentRE.MatchString(f.Comments[0].Text())
}

func (l *linter) getFilename(f *ast.File) string {
	// see https://github.com/golang/go/issues/24498
	return filepath.Base(l.prog.Fset.Position(f.Pos()).Filename)
}

// ExitCode returns status code that should be used as an argument to os.Exit.
func (l *linter) ExitCode() int {
	if l.foundIssues {
		return l.flags.FailureExitCode
	}
	return 0
}

func (l *linter) checkFile(f *ast.File) {
	var wg sync.WaitGroup
	wg.Add(len(l.checkers))
	for _, c := range l.checkers {
		// All checkers are expected to use *lint.Context
		// as read-only structure, so no copying is required.
		go func(c *lint.Checker) {
			defer func() {
				wg.Done()
				// Checker signals unexpected error with panic(error).
				r := recover()
				if r == nil {
					return // There were no panic
				}
				if err, ok := r.(error); ok {
					log.Printf("%s: error: %v\n", c.Rule, err)
					panic(err)
				} else {
					// Some other kind of run-time panic.
					// Undo the recover and resume panic.
					panic(r)
				}
			}()

			for _, warn := range c.Check(f) {
				l.foundIssues = true
				loc := l.ctx.FileSet().Position(warn.Node.Pos()).String()
				if l.flags.ShorterErrLocation {
					loc = shortenLocation(loc)
				}

				printWarning(l, c.Rule.Name(), loc, warn.Text)
			}
		}(c)
	}
	wg.Wait()
}

func shortenLocation(loc string) string {
	switch {
	case strings.HasPrefix(loc, build.Default.GOPATH):
		return strings.Replace(loc, build.Default.GOPATH, "$GOPATH", 1)
	case strings.HasPrefix(loc, build.Default.GOROOT):
		return strings.Replace(loc, build.Default.GOROOT, "$GOROOT", 1)
	default:
		return loc
	}
}

func printWarning(l *linter, rule, loc, warn string) {
	switch {
	case l.flags.JSONOutput:
		// due to sort capabilities,
		// using struct instead of map
		b, err := json.MarshalIndent(struct {
			Rule     string `json:"rule"`
			Location string `json:"location"`
			Warning  string `json:"warning"`
		}{
			Rule:     rule,
			Location: loc,
			Warning:  warn,
		}, "", "    ")
		if err != nil {
			panic(err)
		}
		log.Println(string(b))

	case l.flags.ColoredOutput:
		log.Printf("%v: %v: %v\n",
			aurora.Magenta(aurora.Bold(loc)),
			aurora.Red(rule),
			warn)

	default:
		log.Printf("%s: %s: %s\n", loc, rule, warn)
	}
}
