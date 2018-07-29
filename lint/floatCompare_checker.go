package lint

import (
	"go/ast"
	"go/token"
	"go/types"
)

func init() {
	addChecker(&floatCompareChecker{}, attrExperimental)
}

type floatCompareChecker struct {
	checkerBase
}

func (c *floatCompareChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects fragile float variables comparisons"
	d.Before = `
// x and y are floats
return x == y`
	d.After = `
// x and y are floats
return math.Abs(x - y) < eps`
}

func (c *floatCompareChecker) VisitLocalExpr(expr ast.Expr) {
	binexpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return
	}
	if binexpr.Op == token.EQL || binexpr.Op == token.NEQ {
		typx, ok := c.ctx.typesInfo.TypeOf(binexpr.X).(*types.Basic)
		if ok && typx.Info()&types.IsFloat != 0 {
			c.warn(binexpr)
		}
	}
}

func (c *floatCompareChecker) warn(expr *ast.BinaryExpr) {
	var op string
	if expr.Op == token.EQL {
		op = "<"
	} else {
		op = ">="
	}
	c.ctx.Warn(expr, "change `%s` to `math.Abs(%s - %s) %s eps`",
		expr, expr.X, expr.Y, op)
}
