package lint

import (
	"fmt"
	"go/ast"
)

// UnderefChecker detects expressions, with C style field selection.
type UnderefChecker struct {
	ctx *Context

	warnings []Warning
}

// NewUnderefChecker returns initialized checker for deref (ast.Star) expressions.
func NewUnderefChecker(ctx *Context) *UnderefChecker {
	return &UnderefChecker{
		ctx: ctx,
	}
}

// Check runs underef checker for f.
//
// Features
//
// Detects expressions with C style field selection and
// suggest Go style correction.
func (c *UnderefChecker) Check(f *ast.File) []Warning {
	c.warnings = c.warnings[:0]

	for _, decl := range collectFuncDecls(f) {
		if decl.Body == nil {
			continue
		}
		ast.Inspect(decl.Body, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.SelectorExpr:
				expr, ok := n.X.(*ast.ParenExpr)
				if !ok {
					return true
				}
				if _, ok := expr.X.(*ast.StarExpr); ok {
					c.warn(n)
					return false
				}
			}
			return true
		})
	}
	return c.warnings
}

// TODO add () to function output.
func (c *UnderefChecker) warn(expr *ast.SelectorExpr) {
	c.warnings = append(c.warnings, Warning{
		Node: expr,
		Text: fmt.Sprintf("could simplify %s to %s.%s",
			nodeString(c.ctx.FileSet, expr),
			nodeString(c.ctx.FileSet, expr.X.(*ast.ParenExpr).X.(*ast.StarExpr).X),
			expr.Sel.Name),
	})
}
