package lint

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

// unparen returns expr without surrounding parenthesis.
// Only 1 level of ParenExpr is removed.
// In other words, for ((x)) it returns (x),
// second invocation will return just x.
func unparen(expr ast.Expr) ast.Expr {
	if expr, ok := expr.(*ast.ParenExpr); ok {
		return expr.X
	}
	return expr
}

func nodeString(fset *token.FileSet, x ast.Node) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
