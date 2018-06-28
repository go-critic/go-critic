package lint

//! Detects Yoda style expressions that suggest to replace them.
//
// @Before:
// if nil != ptr {}
// return 10 > a
//
// @After:
// if ptr != nil {}
// return a < 10

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(&yodaStyleExprChecker{}, attrExperimental, attrVeryOpinionated)
}

type yodaStyleExprChecker struct {
	checkerBase
}

func (c *yodaStyleExprChecker) VisitLocalExpr(expr ast.Expr) {
	binexpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return
	}
	switch binexpr.Op {
	case token.EQL, token.NEQ, token.LSS, token.GTR, token.LEQ, token.GEQ:
		if c.isConstExpr(binexpr.X) && !c.isConstExpr(binexpr.Y) {
			c.warn(binexpr)
		}
	default:
		// we are't interested in another operations
	}
}

func (c *yodaStyleExprChecker) isConstExpr(expr ast.Expr) bool {
	return qualifiedName(expr) == "nil" || astp.IsBasicLit(expr)
}

func (c *yodaStyleExprChecker) warn(expr *ast.BinaryExpr) {
	e := astcopy.BinaryExpr(expr)
	switch expr.Op {
	case token.LSS:
		e.Op = token.GTR
	case token.GTR:
		e.Op = token.LSS
	case token.LEQ:
		e.Op = token.GEQ
	case token.GEQ:
		e.Op = token.LEQ
	}
	e.X, e.Y = e.Y, e.X
	c.ctx.Warn(expr, "consider to change order in expression to %s", e)
}
