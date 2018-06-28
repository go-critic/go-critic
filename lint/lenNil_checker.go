package lint

//! Detects redundant x!=nil before len(x).
//
// @Before:
// if myMap != nil && len(myMap) == smth {}
//
// @After:
// if len(myMap) == smth {}

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/go-toolsmith/astequal"
)

func init() {
	addChecker(&lenNilChecker{}, attrExperimental)
}

type lenNilChecker struct {
	checkerBase
}

func (c *lenNilChecker) VisitExpr(expr ast.Expr) {
	binexpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return
	}

	if binexpr.Op == token.LAND || binexpr.Op == token.LOR {
		if c.isNilCheck(binexpr.X) && c.hasLenCall(binexpr.Y, binexpr.X) {
			c.warn(binexpr.X)
		}
		if c.isNilCheck(binexpr.Y) && c.hasLenCall(binexpr.X, binexpr.Y) {
			c.warn(binexpr.Y)
		}
	}
}

func (c *lenNilChecker) isNilCheck(expr ast.Expr) bool {
	binexpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return false
	}
	return (binexpr.Op == token.NEQ || binexpr.Op == token.EQL) &&
		(c.isRefType(binexpr.X) && qualifiedName(binexpr.Y) == "nil" ||
			c.isRefType(binexpr.Y) && qualifiedName(binexpr.X) == "nil")
}

func (c *lenNilChecker) isRefType(expr ast.Expr) bool {
	typ := c.ctx.typesInfo.TypeOf(expr)

	switch typ.(type) {
	case *types.Map, *types.Chan, *types.Slice:
		return true
	default:
		return false
	}
}

func (c *lenNilChecker) hasLenCall(x ast.Expr, v ast.Expr) bool {
	binexpr, ok := v.(*ast.BinaryExpr)
	if !ok {
		return false
	}
	return containsNode(x, func(x ast.Node) bool {
		call, ok := x.(*ast.CallExpr)
		return ok &&
			qualifiedName(call.Fun) == "len" &&
			len(call.Args) == 1 &&
			(astequal.Expr(call.Args[0], binexpr.X) || astequal.Expr(call.Args[0], binexpr.Y))
	})
}

func (c *lenNilChecker) warn(expr ast.Expr) {
	c.ctx.Warn(expr, "%s check is redundant", expr)
}
