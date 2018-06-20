package lint

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

// IsUnitTestFunc reports whether FuncDecl declares testing function.
func (ctx *context) IsUnitTestFuncDecl(fn *ast.FuncDecl) bool {
	if !strings.HasPrefix(fn.Name.Name, "Test") {
		return false
	}
	typ := ctx.typesInfo.TypeOf(fn.Name)
	if sig, ok := typ.(*types.Signature); ok {
		return sig.Results().Len() == 0 &&
			sig.Params().Len() == 1 &&
			sig.Params().At(0).Type().String() == "*testing.T"
	}
	return false
}

// qualifiedName returns called expr fully-quallified name.
//
// It works for simple identifiers like f => "f" and identifiers
// from other package like pkg.f => "pkg.f".
//
// For all unexpected expressions returns empty string.
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

// findNode applies pred for root and all it's childs until it returns true.
// Matched node is returned.
// If none of the nodes matched predicate, nil is returned.
func findNode(root ast.Node, pred func(ast.Node) bool) ast.Node {
	var found ast.Node
	astutil.Apply(root, nil, func(cur *astutil.Cursor) bool {
		if pred(cur.Node()) {
			found = cur.Node()
			return false
		}
		return true
	})
	return found
}
