package criticize

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/go-critic/go-critic/lint"
	"golang.org/x/tools/go/loader"
)

type linter struct {
	ctx *lint.Context

	prog *loader.Program

	checkers []*lint.Checker

	foundIssues bool // True if there any checker reported an issue

	// Command line flags:

	packages        []string
	enabledCheckers []string
	failureExitCode int
}

// Main implements gocritic sub-command entry point.
func Main() {
	var l linter
	parseArgv(&l)
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

	enable := flag.String("enable", "all", "comma-separated list of enabled checkers")
	flag.IntVar(&l.failureExitCode, "failcode", 1, "exit code to be used when lint issues are found")

	flag.Parse()

	l.packages = flag.Args()
	if len(l.packages) == 0 {
		blame("no packages specified\n")
	}

	switch *enable {
	case "all":
		// Special case. l.enabledCheckers remains nil.
	case "":
		// Empty slice. Semantically "disable-all".
		// Can be used to run all pipelines without actual checkers.
		l.enabledCheckers = []string{}
	default:
		// Comma-separated list of names.
		l.enabledCheckers = strings.Split(*enable, ",")
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

	l.ctx = &lint.Context{
		FileSet:   prog.Fset,
		SizesInfo: sizes,
	}
}

func (l *linter) InitCheckers() {
	requested := make(map[string]bool)
	available := lint.RuleList()

	if l.enabledCheckers == nil {
		for _, rule := range available {
			// Exclude experimental checkers from default list.
			if !rule.Experimental {
				requested[rule.Name()] = true
			}
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
		l.checkers = append(l.checkers, lint.NewChecker(rule, l.ctx))
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

	l.ctx.TypesInfo = &pkgInfo.Info
	l.ctx.Package = pkgInfo.Pkg
	for _, f := range pkgInfo.Files {
		l.ctx.Filename = l.getFilename(f)
		l.checkFile(f)
	}
}

func (l *linter) getFilename(f *ast.File) string {
	// see https://github.com/golang/go/issues/24498
	fname := l.prog.Fset.Position(f.Pos()).String() // ex: /usr/go/src/pkg/main.go:1:1
	fname = filepath.Base(fname)                    // ex: main.go:1:1
	return fname[:len(fname)-4]                     // ex: main.go
}

// ExitCode returns status code that should be used as an argument to os.Exit.
func (l *linter) ExitCode() int {
	if l.foundIssues {
		return l.failureExitCode
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
				pos := l.ctx.FileSet.Position(warn.Node.Pos())
				log.Printf("%s: %s: %v\n", pos, c.Rule, warn.Text)
			}
		}(c)
	}
	wg.Wait()
}
