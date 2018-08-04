package lint

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/go-toolsmith/astequal"
	"github.com/go-toolsmith/astp"
	"golang.org/x/tools/go/ast/astutil"
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
		if c.isFloatExpr(binexpr) &&
			c.isMultiBinaryExpr(binexpr) &&
			!c.isNaNCheckExpr(binexpr) &&
			!c.isInfCheckExpr(binexpr) {
			c.warn(binexpr)
		}
	}
}

func (c *floatCompareChecker) isFloatExpr(binexpr *ast.BinaryExpr) bool {
	exprx := astutil.Unparen(binexpr.X)
	typx, ok := c.ctx.typesInfo.TypeOf(exprx).(*types.Basic)
	return ok && typx.Info()&types.IsFloat != 0
}

func (c *floatCompareChecker) isNaNCheckExpr(binexpr *ast.BinaryExpr) bool {
	exprx := astutil.Unparen(binexpr.X)
	expry := astutil.Unparen(binexpr.Y)
	return astequal.Expr(exprx, expry)
}

func (c *floatCompareChecker) isInfCheckExpr(binexpr *ast.BinaryExpr) bool {
	binx := astutil.Unparen(binexpr.X)
	biny := astutil.Unparen(binexpr.Y)
	expr, bin, ok := c.identExpr(binx, biny)
	if !ok {
		return false
	}
	x := astutil.Unparen(expr)
	y := astutil.Unparen(bin.X)
	z := astutil.Unparen(bin.Y)
	return astequal.Expr(x, y) && astequal.Expr(y, z)
}

func (c *floatCompareChecker) isMultiBinaryExpr(binexpr *ast.BinaryExpr) bool {
	exprx := astutil.Unparen(binexpr.X)
	expry := astutil.Unparen(binexpr.Y)
	return astp.IsBinaryExpr(exprx) || astp.IsBinaryExpr(expry)
}

func (c *floatCompareChecker) identExpr(x, y ast.Node) (ast.Expr, *ast.BinaryExpr, bool) {
	expr1, ok1 := x.(*ast.BinaryExpr)
	expr2, ok2 := y.(*ast.BinaryExpr)
	switch {
	case ok1 && !ok2:
		return y.(ast.Expr), expr1, true
	case !ok1 && ok2:
		return x.(ast.Expr), expr2, true

	default:
		return nil, nil, false
	}
}

func (c *floatCompareChecker) warn(expr *ast.BinaryExpr) {
	op := ">="
	if expr.Op == token.EQL {
		op = "<"
	}
	format := "change `%s` to `math.Abs(%s - %s) %s eps`"
	if astp.IsBinaryExpr(expr.Y) {
		format = "change `%s` to `math.Abs(%s - (%s)) %s eps`"
	}
	c.ctx.Warn(expr, format, expr, expr.X, expr.Y, op)
}
