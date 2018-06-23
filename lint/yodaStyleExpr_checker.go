package lint

//! Detects Yoda style expressions that suggest to replace them.
//
// @Before:
// if nil != ptr {}
//
// @After:
// if ptr != nil {}

import (
	"go/ast"
	"go/token"
)

func init() {
	addChecker(&yodaStyleExprChecker{})
}

type yodaStyleExprChecker struct {
	checkerBase
}

func (c *yodaStyleExprChecker) VisitLocalExpr(expr ast.Expr) {
	binexpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return
	}
	if binexpr.Op == token.EQL || binexpr.Op == token.NEQ {
		if c.isConstExpr(binexpr.X) && c.isVar(binexpr.Y) {
			c.warn(binexpr)
		}
	}
}

func (c *yodaStyleExprChecker) isConstExpr(expr ast.Expr) bool {
	if qualifiedName(expr) == "nil" {
		return true
	}
	_, ok := expr.(*ast.BasicLit)
	return ok
}

func (c *yodaStyleExprChecker) isVar(expr ast.Expr) bool {
	switch expr.(type) {
	case *ast.Ident, *ast.SelectorExpr:
		return true
	default:
		return false
	}
}

func (c *yodaStyleExprChecker) warn(expr ast.Expr) {
	c.ctx.Warn(expr, "consider to change order of expression in %s", expr)
}
