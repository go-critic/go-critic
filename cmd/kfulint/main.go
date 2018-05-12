package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/PieselBois/kfulint/lint"
	"golang.org/x/tools/go/loader"
)

func main() {
	log.SetFlags(0)

	var l linter
	parseArgv(&l)
	l.LoadProgram()
	l.InitContext()
	l.InitCheckers()
	for _, pkgPath := range l.packages {
		l.CheckPackage(pkgPath)
	}
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
		fmt.Fprintln(os.Stderr, "usage: kfulint [flags] [package ...]")
		flag.PrintDefaults()
	}

	enable := flag.String("enable", "all", "comma-separated list of enabled checkers")

	flag.Parse()

	l.packages = flag.Args()
	if len(l.packages) == 0 {
		blame("no packages specified\n")
	}

	switch *enable {
	case "all":
		// Special case. l.enableSet remains nil.
	case "":
		// Empty set. Semantically "disable-all".
		// Can be used to run all pipelines without actual checkers.
		l.enabledSet = map[string]bool{}
	default:
		// Comma-separated list of names.
		l.enabledSet = make(map[string]bool)
		nameRE := regexp.MustCompile(`[a-z][a-z0-9_\-]*`)
		for i, s := range strings.Split(*enable, ",") {
			s = strings.TrimSpace(s)
			l.enabledSet[s] = true
			if !nameRE.MatchString(s) {
				log.Fatalf("-enable element #%d: invalid %q checker name", i, s)
			}
		}
	}

}

type checker struct {
	name string
	lint.Checker
}

type linter struct {
	ctx *lint.Context

	prog *loader.Program

	checkers []checker

	// Command line flags:

	packages   []string
	enabledSet map[string]bool
}

func (l *linter) LoadProgram() {
	conf := loader.Config{
		ParserMode: parser.ParseComments,
	}

	if _, err := conf.FromArgs(l.packages, true); err != nil {
		log.Fatalf("resolve packages: %v", err)
	}
	prog, err := conf.Load()
	if err != nil {
		log.Fatalf("load program: %v", err)
	}

	l.prog = prog
}

func (l *linter) InitContext() {
	l.ctx = &lint.Context{
		FileSet: l.prog.Fset,
	}
}

func (l *linter) InitCheckers() {
	checkers := []checker{
		{"param-name", lint.NewParamNameChecker(l.ctx)},
		{"type-guard", lint.NewTypeGuardChecker(l.ctx)},
		{"parenthesis", lint.NewParenthesisChecker(l.ctx)},
		{"underef", lint.NewUnderefChecker(l.ctx)},
		{"param-duplication", lint.NewParamDuplicationChecker(l.ctx)},
		{"stddef", lint.NewStddefChecker(l.ctx)},
	}

	for _, c := range checkers {
		// Nil enabledSet means "all checkers are enabled".
		if l.enabledSet == nil || l.enabledSet[c.name] {
			l.checkers = append(l.checkers, c)
		}
	}
}

func (l *linter) CheckPackage(pkgPath string) {
	pkgInfo := l.prog.Imported[pkgPath]
	if pkgInfo == nil || !pkgInfo.TransitivelyErrorFree {
		log.Fatalf("%s package is not properly loaded", pkgPath)
	}

	l.ctx.TypesInfo = &pkgInfo.Info
	for _, f := range pkgInfo.Files {
		l.checkFile(f)
	}
}

func (l *linter) checkFile(f *ast.File) {
	var wg sync.WaitGroup
	wg.Add(len(l.checkers))
	for _, c := range l.checkers {
		// All checkers are expected to use *lint.Context
		// as read-only structure, so no copying is required.
		go func(c checker) {
			defer func() {
				wg.Done()
				// Checker signals unexpected error with panic(error).
				r := recover()
				if r == nil {
					return // There were no panic
				}
				if err, ok := r.(error); ok {
					log.Printf("%s: error: %v\n", c.name, err)
					panic(err)
				} else {
					// Some other kind of run-time panic.
					// Undo the recover and resume panic.
					panic(r)
				}
			}()
			for _, w := range c.Check(f) {
				pos := l.ctx.FileSet.Position(w.Node.Pos())
				name := c.name
				if w.Kind != "" {
					name += "/" + string(w.Kind)
				}
				log.Printf("%s: %s: %v\n", pos, name, w.Text)
			}
		}(c)
	}
	wg.Wait()
}
