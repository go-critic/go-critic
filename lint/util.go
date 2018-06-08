package lint

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"strings"
)

func nodeString(fset *token.FileSet, x ast.Node) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}

// IsUnitTestFunc reports whether FuncDecl declares testing function.
func (ctx *context) IsUnitTestFuncDecl(fn *ast.FuncDecl) bool {
	if !strings.HasPrefix(fn.Name.Name, "Test") {
		return false
	}
	typ := ctx.TypesInfo.TypeOf(fn.Name)
	if sig, ok := typ.(*types.Signature); ok {
		return sig.Results().Len() == 0 &&
			sig.Params().Len() == 1 &&
			sig.Params().At(0).Type().String() == "*testing.T"
	}
	return false
}

// functionName returns called function fully-quallified name.
//
// It works for simple function calls like f() => "f" and functions
// from other package like pkg.f() => "pkg.f".
//
// For all unexpected expressions returns empty string.
func functionName(x *ast.CallExpr) string {
	return qualifiedName(x.Fun)
}

func qualifiedName(x ast.Expr) string {
	switch x := x.(type) {
	case *ast.SelectorExpr:
		pkg, ok := x.X.(*ast.Ident)
		if !ok {
			return ""
		}
		return pkg.Name + "." + x.Sel.Name
	case *ast.Ident:
		return x.Name
	default:
		return ""
	}
}
