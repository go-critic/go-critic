package main

import (
	"flag"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/PieselBois/kfulint/lint"
)

func main() {
	log.SetFlags(0)

	l := linter{
		ctx: &lint.Context{
			FileSet: token.NewFileSet(),
		},
		typesChecker: &types.Config{
			Importer: importer.Default(),
		},
	}

	parseArgv(&l)
	pkgs := parseDir(l.ctx.FileSet, l.targetDir)

	l.InitCheckers()
	for _, pkg := range pkgs {
		// TODO: if linters require ast.Package, assign it here.
		files := sortedPackageFiles(pkg)
		if err := l.CollectTypesInfo(files); err != nil {
			log.Printf("skip %s: type check error: %v", pkg.Name, err)
			continue
		}
		for _, f := range files {
			l.CheckFile(f)
		}
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
	flag.StringVar(&l.targetDir, "dir", "", "target package(s) directory")
	enable := flag.String("enable", "all", "comma-separated list of enabled checkers")

	flag.Parse()

	if l.targetDir == "" {
		blame("Illegal empty -dir argument\n")
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

func parseDir(fset *token.FileSet, path string) []*ast.Package {
	pkgMap, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("parse dir error: %v", err)
	}

	var pkgList []*ast.Package
	for _, pkg := range pkgMap {
		pkgList = append(pkgList, pkg)
	}

	sort.Slice(pkgList, func(i, j int) bool {
		return pkgList[i].Name < pkgList[j].Name
	})
	return pkgList
}

func sortedPackageFiles(pkg *ast.Package) []*ast.File {
	var files []*ast.File
	for _, f := range pkg.Files {
		files = append(files, f)
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name.Name < files[j].Name.Name
	})
	return files
}

type checker struct {
	name string
	lint.Checker
}

type linter struct {
	ctx *lint.Context

	typesChecker *types.Config

	checkers []checker

	// Command line flags:

	targetDir  string
	enabledSet map[string]bool
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

func (l *linter) CheckFile(f *ast.File) {
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

func (l *linter) CollectTypesInfo(files []*ast.File) error {
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}
	l.ctx.TypesInfo = info
	// TODO: if lint.Context needs types.Scope or types.Package, assign it here.
	_, err := l.typesChecker.Check(l.targetDir, l.ctx.FileSet, files, info)
	return err
}
