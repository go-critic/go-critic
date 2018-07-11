package lint

//! Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`.
//
// @Before:
// re, _ := regexp.Compile(`const pattern`)
//
// @After:
// re := regexp.MustCompile(`const pattern`)

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcopy"

	"github.com/go-toolsmith/astfmt"
)

func init() {
	addChecker(&dummyLenChecker{}, attrSyntaxOnly, attrExperimental)
}

type dummyLenChecker struct {
	checkerBase
}

func (c *dummyLenChecker) VisitExpr(x ast.Expr) {
	expr, ok := x.(*ast.BinaryExpr)
	if !ok {
		return
	}

	if expr.Op == token.LSS || expr.Op == token.LEQ || expr.Op == token.GEQ {
		if c.isLenCall(expr.X) && c.isZero(expr.Y) {
			c.warn(expr)
		}
	}
}

func (c *dummyLenChecker) isLenCall(x ast.Expr) bool {
	call, ok := x.(*ast.CallExpr)
	return ok && qualifiedName(call.Fun) == "len" && len(call.Args) == 1
}

func (c *dummyLenChecker) isZero(x ast.Expr) bool {
	value, ok := x.(*ast.BasicLit)
	return ok && value.Value == "0"
}

func (c *dummyLenChecker) warn(cause *ast.BinaryExpr) {
	info := ""
	switch cause.Op {
	case token.LSS:
		info = "always false"
	case token.LEQ:
		expr := astcopy.BinaryExpr(cause)
		expr.Op = token.EQL
		info = astfmt.Sprintf("can be %s", expr)
	case token.GEQ:
		info = "always true"
	}
	c.ctx.Warn(cause, "useless len comparison %s, %s", cause, info)
}
