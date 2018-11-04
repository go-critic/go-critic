package lint

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

func (c *yodaStyleExprChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects Yoda style expressions and suggests to replace them"
	d.Before = `return nil != ptr`
	d.After = `return ptr != nil`
}

func (c *yodaStyleExprChecker) VisitLocalExpr(expr ast.Expr) {
	binexpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return
	}
	if binexpr.Op == token.EQL || binexpr.Op == token.NEQ {
		if c.isConstExpr(binexpr.X) && !c.isConstExpr(binexpr.Y) {
			c.warn(binexpr)
		}
	}
}

func (c *yodaStyleExprChecker) isConstExpr(expr ast.Expr) bool {
	return qualifiedName(expr) == "nil" || astp.IsBasicLit(expr)
}

func (c *yodaStyleExprChecker) warn(expr *ast.BinaryExpr) {
	e := astcopy.BinaryExpr(expr)
	e.X, e.Y = e.Y, e.X
	c.ctx.Warn(expr, "consider to change order in expression to %s", e)
}
