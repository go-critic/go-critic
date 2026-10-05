package checkers

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"runtime"
	"testing"

	"github.com/go-critic/go-critic/linter"
)

// TestUnnamedResultCheckExported checks that checkExported=true reports both
// exported and unexported functions with unnamed same-typed results.
//
// The checker reads the flag from the registered CheckerInfo when it is
// constructed. NewChecker has no params override, so the shared value is
// set for this test and restored afterwards. TestCheckers expects the default.
func TestUnnamedResultCheckExported(t *testing.T) {
	const src = `package p

func ExportedPair() (int, int) { return 0, 0 }

func unexportedPair() (int, int) { return 0, 0 }
`
	const wantText = "consider giving a name to these results"

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "unnamed_result.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	typeInfo := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	var conf types.Config
	pkg, err := conf.Check("p", fset, []*ast.File{f}, typeInfo)
	if err != nil {
		t.Fatalf("typecheck: %v", err)
	}

	info := unnamedResultCheckerInfo(t)
	param := info.Params["checkExported"]
	if param == nil {
		t.Fatal("checkExported param not found")
	}
	old := param.Value
	param.Value = true
	t.Cleanup(func() { param.Value = old })

	ctx := linter.NewContext(fset, types.SizesFor("gc", runtime.GOARCH))
	ctx.SetPackageInfo(typeInfo, pkg)
	ctx.SetFileInfo("unnamed_result.go", f)

	c, err := linter.NewChecker(ctx, info)
	if err != nil {
		t.Fatalf("NewChecker: %v", err)
	}

	got := map[string]string{}
	for _, warn := range c.Check(f) {
		name := funcDeclNameAt(f, warn.Pos)
		if name == "" {
			t.Errorf("warning %q at unknown position", warn.Text)
			continue
		}
		got[name] = warn.Text
	}

	for _, name := range []string{"ExportedPair", "unexportedPair"} {
		if got[name] != wantText {
			t.Errorf("%s: got warning %q, want %q", name, got[name], wantText)
		}
	}
	if len(got) != 2 {
		t.Errorf("got warnings %#v, want only ExportedPair and unexportedPair", got)
	}
}

func unnamedResultCheckerInfo(t *testing.T) *linter.CheckerInfo {
	t.Helper()
	for _, info := range linter.GetCheckersInfo() {
		if info.Name == "unnamedResult" {
			return info
		}
	}
	t.Fatal("unnamedResult checker is not registered")
	return nil
}

func funcDeclNameAt(f *ast.File, pos token.Pos) string {
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if ok && fn.Pos() == pos {
			return fn.Name.Name
		}
	}
	return ""
}
