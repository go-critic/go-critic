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
	d.Summary = "Detects bool expressions that compare float variables"
	d.Before = `
var a, b float64 = 10.0, 20.0
return a == b`
	d.After = `
var a, b float64 = 10.0, 20.0
return math.Abs(a - b) < eps`
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
	c.ctx.Warn(expr, "consider to change way to compare floats in expression"+
		" to math.Abs(%s - %s) %s eps", expr.X, expr.Y, op)
}
