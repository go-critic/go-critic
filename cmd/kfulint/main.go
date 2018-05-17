package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/types"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/PieselBois/kfulint/lint"
	"golang.org/x/tools/go/loader"
)

type checker struct {
	name string
	lint.Checker
}

type linter struct {
	ctx *lint.Context

	prog *loader.Program

	checkers []checker

	// Command line flags:

	packages        []string
	enabledCheckers []string
}

func main() {
	log.SetFlags(0)

	var l linter
	parseArgv(&l)
	l.LoadProgram()
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
	for _, name := range l.enabledCheckers {
		c, ok := lint.NewChecker(name, l.ctx)
		if !ok {
			log.Fatalf("%s: checker not found", name)
		}
		l.checkers = append(l.checkers, checker{name, c})
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

	for _, w := range l.ctx.Warnings {
		pos := l.ctx.FileSet.Position(w.Node.Pos())
		log.Printf("%s: %s: %v\n", pos, string(w.Kind), w.Text)
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
			c.Check(f)
		}(c)
	}
	wg.Wait()
}
