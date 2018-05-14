package lint

import (
	"fmt"
	"go/ast"
	"go/types"
)

// UnderefChecker detects expressions, with C style field selection.
type UnderefChecker struct {
	ctx *Context
}

// NewUnderefChecker returns initialized checker for deref (ast.Star) expressions.
func newUnderefChecker(ctx *Context) Checker {
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
func (c *UnderefChecker) Check(f *ast.File) {
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
				if expr, ok := expr.X.(*ast.StarExpr); ok {
					if c.checkStarExpr(expr) {
						c.warnSelect(n)
						return false
					}
				}
			case *ast.IndexExpr:
				expr, ok := n.X.(*ast.ParenExpr)
				if !ok {
					return true
				}
				if expr, ok := expr.X.(*ast.StarExpr); ok {
					if !c.checkStarExpr(expr) {
						return true
					}
					if c.checkArray(expr) {
						c.warnArray(n)
					}
				}
			}

			return true
		})
	}
}

// TODO add () to function output.
func (c *UnderefChecker) warnSelect(expr *ast.SelectorExpr) {
	c.ctx.addWarning(Warning{
		Kind: "underef",
		Node: expr,
		Text: fmt.Sprintf("could simplify %s to %s.%s",
			nodeString(c.ctx.FileSet, expr),
			nodeString(c.ctx.FileSet, expr.X.(*ast.ParenExpr).X.(*ast.StarExpr).X),
			expr.Sel.Name),
	})
}

func (c *UnderefChecker) warnArray(expr *ast.IndexExpr) {
	c.ctx.addWarning(Warning{
		Kind: "underef",
		Node: expr,
		Text: fmt.Sprintf("could simplify %s to %s[%s]",
			nodeString(c.ctx.FileSet, expr),
			nodeString(c.ctx.FileSet, expr.X.(*ast.ParenExpr).X.(*ast.StarExpr).X),
			nodeString(c.ctx.FileSet, expr.Index)),
	})
}

// checkStarExpr checks if ast.StarExpr could be simplified
func (c *UnderefChecker) checkStarExpr(expr *ast.StarExpr) bool {
	// checks if expr is pointer type
	typ, ok := c.ctx.TypesInfo.TypeOf(expr.X).(*types.Pointer)
	if !ok {
		return false
	}

	// checks if dereference of typ is pointer
	if _, ok := typ.Elem().(*types.Pointer); ok {
		return false
	}

	return true
}

func (c *UnderefChecker) checkArray(expr *ast.StarExpr) bool {
	typ, ok := c.ctx.TypesInfo.TypeOf(expr.X).(*types.Pointer)
	if !ok {
		return false
	}
	_, ok = typ.Elem().(*types.Array)
	return ok
}
