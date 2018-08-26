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

	"github.com/go-critic/go-critic/cmd/internal/config"
	"github.com/go-critic/go-critic/cmd/internal/flagparser"
	"github.com/go-critic/go-critic/lint"
	"golang.org/x/tools/go/loader"
)

var generatedFileCommentRE = regexp.MustCompile("Code generated .* DO NOT EDIT.")

type linter struct {
	ctx *lint.Context

	prog *loader.Program

	checkers []*lint.Checker

	foundIssues bool // True if there any checker reported an issue

	// Command line flags:

	flags *flagparser.FlagParser

	packages []string

	configuration *config.Config
}

// parseEnabledCheckers processes enabled checkers, intersect them with disabled, etc.
func (l *linter) parseEnabledCheckers() {
	l.configuration = config.NewConfig()

	if l.flags.Disable == flagparser.DisableAll {
		return
	}

	disabledCheckers := make(map[string]bool)

	for _, d := range l.flags.DisabledCheckers() {
		disabledCheckers[d] = true
	}

	switch l.flags.Enable {
	case flagparser.EnableAll:
		for _, rule := range lint.RuleList() {
			if ok := disabledCheckers[rule.Name()]; ok {
				continue
			}
			if rule.Experimental && !l.flags.WithExperimental {
				continue
			}
			if rule.VeryOpinionated && !l.flags.WithOpinionated {
				continue
			}
			l.configuration.Checkers[rule.Name()] = &config.Checker{Type: rule.Name()}
		}
	default:
		// Comma-separated list of names.
		for _, checkerName := range l.flags.EnabledCheckers() {
			if ok := disabledCheckers[checkerName]; ok {
				continue
			}
			l.configuration.Checkers[checkerName] = &config.Checker{Type: checkerName}
		}
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
		log.Fatalf("cannot read config file %s, got error: %s", l.flags.ConfigFile, err)
		return
	}

	l.configuration = config.NewConfig()

	if err := json.Unmarshal(raw, l.configuration); err != nil {
		log.Fatalf("cannot parse config file, got error: %s", err)
		return
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
	var notFoundCheckers []string

	available := make(map[string]*lint.Rule)

	for _, rule := range lint.RuleList() {
		available[rule.Name()] = rule
	}

	for checkerName, checkerConfig := range l.configuration.Checkers {
		if rule, ok := available[checkerConfig.Type]; ok {
			l.checkers = append(l.checkers, lint.NewChecker(
				rule,
				l.ctx,
				checkerConfig.Parameters,
			))
		} else {
			notFoundCheckers = append(notFoundCheckers, checkerName)
		}
	}

	if len(notFoundCheckers) != 0 {
		for name := range notFoundCheckers {
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
		if l.flags.CheckGenerated || !isGenerated(f) {
			l.ctx.SetFileInfo(l.getFilename(f))
			l.checkFile(f)
		}
	}
}

func isGenerated(f *ast.File) bool {
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
				log.Printf("%s: %s: %v\n", loc, c.Rule, warn.Text)
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
