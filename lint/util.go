package lint

import (
	"go/ast"
	"go/token"
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

// containsNode reports whether findNode(root, pred) returned non-nil node.
func containsNode(root ast.Node, pred func(ast.Node) bool) bool {
	return findNode(root, pred) != nil
}

// identOf returns identifier for x that can be used to obtain associated types.Object.
// Returns nil for expressions that yield temporary results, like `f().field`.
func identOf(x ast.Node) *ast.Ident {
	switch x := x.(type) {
	case *ast.Ident:
		// Found ident.
		return x
	case *ast.TypeAssertExpr:
		// x.(type) - x may contain ident.
		return identOf(x.X)
	case *ast.IndexExpr:
		// x[i] - x may contain ident.
		return identOf(x.X)
	case *ast.SelectorExpr:
		// x.y - x may contain ident.
		return identOf(x.X)
	case *ast.StarExpr:
		// *x - x may contain ident.
		return identOf(x.X)
	case *ast.SliceExpr:
		// x[:] - x may contain ident.
		return identOf(x.X)

	default:
		// Note that this function is not comprehensive.
		return nil
	}
}

// typeIsPointer reports whether typ has type of *types.Pointer.
func typeIsPointer(typ types.Type) bool {
	_, ok := typ.(*types.Pointer)
	return ok
}

// isSafeExpr reports whether expr is softly safe expression and contains
// no significant side-effects. As opposed to strictly safe expressions,
// soft safe expressions permit some forms of side-effects, like
// panic possibility during indexing.
func isSafeExpr(expr ast.Expr) bool {
	// This list switch is not comprehensive and uses
	// whitelist to be on the conservative side.
	// Can be extended as needed.
	//
	// Note that it is not very strict "safe" as
	// index expressions are permitted even though they
	// may cause panics.
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		return isSafeExpr(expr.X) && isSafeExpr(expr.Y)
	case *ast.UnaryExpr:
		return expr.Op != token.ARROW && isSafeExpr(expr.X)
	case *ast.BasicLit, *ast.Ident:
		return true
	case *ast.IndexExpr:
		return isSafeExpr(expr.X) && isSafeExpr(expr.Index)
	case *ast.SelectorExpr:
		return isSafeExpr(expr.X)
	case *ast.ParenExpr:
		return isSafeExpr(expr.X)
	default:
		return false
	}
}
