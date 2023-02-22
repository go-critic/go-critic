package linttest

import (
	"go/ast"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/go-critic/go-critic/linter"

	"github.com/go-toolsmith/pkgload"
	"golang.org/x/tools/go/packages"
)

var sizes = types.SizesFor("gc", runtime.GOARCH)

func saneCheckersList(t *testing.T, checkers []*linter.CheckerInfo) []*linter.CheckerInfo {
	var saneList []*linter.CheckerInfo

	for _, info := range checkers {
		pkgPath := "github.com/go-critic/go-critic/framework/linttest/testdata/sanity"
		t.Run(info.Name+"/sanity", func(t *testing.T) {
			defer func() {
				r := recover()
				if r != nil {
					t.Errorf("unexpected panic: %v\n%s", r, debug.Stack())
				} else {
					saneList = append(saneList, info)
				}
			}()

			fset := token.NewFileSet()
			pkgs := newPackages(t, pkgPath, fset)
			for _, pkg := range pkgs {
				ctx := &linter.Context{
					SizesInfo: sizes,
					FileSet:   fset,
					TypesInfo: pkg.TypesInfo,
					Pkg:       pkg.Types,
				}
				c, err := linter.NewChecker(ctx, info)
				if err != nil {
					t.Errorf("Unexpected error: %v\n%s", err, debug.Stack())
				}
				for _, f := range pkg.Syntax {
					ctx.SetFileInfo(getFilename(fset, f), f)
					_ = c.Check(f)
				}
			}
		})
	}

	return saneList
}

// IntegrationTest specifies integration test options.
type IntegrationTest struct {
	Main string

	// Dir specifies a path to integration tests.
	Dir string
}

type CheckersTest struct {
	// IgnoreErrors is a checker names list those tests ignore parse/typecheck errors.
	IgnoreErrors []string
}

// Run executes every registered checker tests.
//
// TODO(quasilyte): document default options.
func (cfg *CheckersTest) Run(t *testing.T) {
	checkers := linter.GetCheckersInfo()

	ignoreErrors := make(map[string]bool, len(cfg.IgnoreErrors))
	for _, checkerName := range cfg.IgnoreErrors {
		ignoreErrors[checkerName] = true
	}

	// See #980.
	for i := range checkers {
		info := checkers[i]
		t.Run(info.Name+"/debug", func(t *testing.T) {
			debugFile := filepath.Join("testdata", info.Name, "debug.go")
			target := lintTarget{
				pattern:      debugFile,
				ignoreErrors: ignoreErrors[info.Name],
			}
			checkTarget(t, target, info)
		})
	}

	for _, info := range saneCheckersList(t, checkers) {
		t.Run(info.Name, func(t *testing.T) {
			pkgPath := "./testdata/" + info.Name
			target := lintTarget{
				pattern:      pkgPath,
				ignoreErrors: ignoreErrors[info.Name],
			}
			checkTarget(t, target, info)
		})
	}
}

type lintTarget struct {
	pattern      string
	ignoreErrors bool
}

func checkTarget(t *testing.T, target lintTarget, info *linter.CheckerInfo) {
	t.Helper()
	fset := token.NewFileSet()
	pkgs := newPackages(t, target.pattern, fset)
	for _, pkg := range pkgs {
		if len(pkg.Errors) != 0 && !target.ignoreErrors {
			for _, err := range pkg.Errors {
				t.Error(err)
			}
			return
		}
		ctx := &linter.Context{
			SizesInfo: sizes,
			FileSet:   fset,
			TypesInfo: pkg.TypesInfo,
			Pkg:       pkg.Types,
		}
		c, err := linter.NewChecker(ctx, info)
		if err != nil {
			t.Errorf("Unexpected error: %v\n%s", err, debug.Stack())
		}
		for _, f := range pkg.Syntax {
			checkFile(t, c, ctx, f)
		}
	}
}

func checkFile(t *testing.T, c *linter.Checker, ctx *linter.Context, f *ast.File) {
	filename := getFilename(ctx.FileSet, f)
	testFilename := filepath.Join("testdata", c.Info.Name, filename)

	rc, err := os.Open(testFilename)
	if err != nil {
		t.Fatalf("read file %q: %v", testFilename, err)
	}
	defer rc.Close()

	ws, err := newWarnings(rc)
	if err != nil {
		t.Fatal(err)
	}

	stripDirectives(f)
	ctx.SetFileInfo(filename, f)

	matched := make(map[*string]struct{})
	for _, warn := range c.Check(f) {
		line := ctx.FileSet.Position(warn.Pos).Line

		if w := ws.find(line, warn.Text); w != nil {
			if _, seen := matched[w]; seen {
				t.Errorf("%s:%d: multiple matches for %s",
					testFilename, line, *w)
			}
			matched[w] = struct{}{}
		} else {
			t.Errorf("%s:%d: unexpected warn: %s",
				testFilename, line, warn.Text)
		}
	}

	checkUnmatched(ws, matched, t, testFilename)
}

// stripDirectives replaces "///" comments with empty single-line
// comments, so the checkers that inspect comments see ordinary
// comment groups (with extra newlines, but that's not important).
func stripDirectives(f *ast.File) {
	for _, cg := range f.Comments {
		for _, c := range cg.List {
			if strings.HasPrefix(c.Text, "/// ") {
				c.Text = "//"
			}
		}
	}
}

func getFilename(fset *token.FileSet, f *ast.File) string {
	// see https://github.com/golang/go/issues/24498
	return filepath.Base(fset.Position(f.Pos()).Filename)
}

func checkUnmatched(ws warnings, matched map[*string]struct{}, t *testing.T, testFilename string) {
	for line, sl := range ws {
		for i, w := range sl {
			if _, ok := matched[&sl[i]]; !ok {
				t.Errorf("%s:%d: unmatched `%s`", testFilename, line, w)
			}
		}
	}
}

func newPackages(t *testing.T, pattern string, fset *token.FileSet) []*packages.Package {
	mode := packages.NeedName |
		packages.NeedFiles |
		packages.NeedCompiledGoFiles |
		packages.NeedImports |
		packages.NeedTypes |
		packages.NeedSyntax |
		packages.NeedTypesInfo |
		packages.NeedTypesSizes
	cfg := packages.Config{
		Mode:  mode,
		Tests: true,
		Fset:  fset,
	}
	pkgs, err := pkgload.LoadPackages(&cfg, []string{pattern})
	if err != nil {
		t.Fatalf("load package: %v", err)
	}
	return pkgs
}
