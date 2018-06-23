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

	"github.com/cristaloleg/astp"
	"github.com/go-toolsmith/astcopy"
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
		if c.isConstExpr(binexpr.X) && !c.isConstExpr(binexpr.Y) { //c.isVar(binexpr.Y) {
			c.warn(binexpr)
		}
	}
}

func (c *yodaStyleExprChecker) isConstExpr(expr ast.Expr) bool {
	return qualifiedName(expr) == "nil" || astp.IsBasicLit(expr)
}

func (c *yodaStyleExprChecker) isVar(expr ast.Expr) bool {
	return astp.IsIdent(expr) || astp.IsSelectorExpr(expr)
}

func (c *yodaStyleExprChecker) warn(expr *ast.BinaryExpr) {
	e := astp.AsBinaryExpr(astcopy.Expr(expr))
	e.X, e.Y = e.Y, e.X
	c.ctx.Warn(expr, "consider to change order in expression to %s", e)
}
