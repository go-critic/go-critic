package lint

//! Detects usage of `len` when result is obvious or doesn't make sense.
//
// @Before:
// len(arr) >= 0
// len(arr) <= 0
// len(arr) < 0
//
// @After:
// len(arr) > 0
// len(arr) == 0

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/astfmt"
)

func init() {
	addChecker(&uselessLenChecker{}, attrSyntaxOnly, attrExperimental)
}

type uselessLenChecker struct {
	checkerBase
}

func (c *uselessLenChecker) VisitExpr(x ast.Expr) {
	expr, ok := x.(*ast.BinaryExpr)
	if !ok {
		return
	}

	if expr.Op == token.LSS || expr.Op == token.GEQ || expr.Op == token.LEQ {
		if c.isLenCall(expr.X) && c.isZero(expr.Y) {
			c.warn(expr)
		}
	}
}

func (c *uselessLenChecker) isLenCall(x ast.Expr) bool {
	call, ok := x.(*ast.CallExpr)
	return ok && qualifiedName(call.Fun) == "len" && len(call.Args) == 1
}

func (c *uselessLenChecker) isZero(x ast.Expr) bool {
	value, ok := x.(*ast.BasicLit)
	return ok && value.Value == "0"
}

func (c *uselessLenChecker) warn(cause *ast.BinaryExpr) {
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
