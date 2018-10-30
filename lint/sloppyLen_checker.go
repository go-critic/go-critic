package lint

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcast"

	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/astfmt"
)

func init() {
	addChecker(&sloppyLenChecker{}, attrSyntaxOnly, attrExperimental)
}

type sloppyLenChecker struct {
	checkerBase
}

func (c *sloppyLenChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects usage of `len` when result is obvious or doesn't make sense"
	d.Before = `
len(arr) >= 0 // Sloppy
len(arr) <= 0 // Sloppy
len(arr) < 0  // Doesn't make sense at all`
	d.After = `
len(arr) > 0
len(arr) == 0`
}

func (c *sloppyLenChecker) VisitExpr(x ast.Expr) {
	expr := astcast.ToBinaryExpr(x)
	if expr.Op == token.LSS || expr.Op == token.GEQ || expr.Op == token.LEQ {
		if c.isLenCall(expr.X) && c.isZero(expr.Y) {
			c.warn(expr)
		}
	}
}

func (c *sloppyLenChecker) isLenCall(x ast.Expr) bool {
	call, ok := x.(*ast.CallExpr)
	return ok && qualifiedName(call.Fun) == "len" && len(call.Args) == 1
}

func (c *sloppyLenChecker) isZero(x ast.Expr) bool {
	value, ok := x.(*ast.BasicLit)
	return ok && value.Value == "0"
}

func (c *sloppyLenChecker) warn(cause *ast.BinaryExpr) {
	info := ""
	switch cause.Op {
	case token.LSS:
		info = "is always false"
	case token.GEQ:
		info = "is always true"
	case token.LEQ:
		expr := astcopy.BinaryExpr(cause)
		expr.Op = token.EQL
		info = astfmt.Sprintf("can be %s", expr)
	}
	c.ctx.Warn(cause, "%s %s", cause, info)
}
