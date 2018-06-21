package lint

//! Detects redundant x!=nil before len(x).
//
// @Before:
// if myMap != nil && len(myMap) == smth {
//
// @After:
// if len(myMap) == smth {

import (
	"go/ast"
	"go/token"
	"go/types"
)

func init() {
	addChecker(&lenNilChecker{})
}

type lenNilChecker struct {
	checkerBase
}

func (c *lenNilChecker) VisitStmt(stmt ast.Stmt) {
	ifstmt, ok := stmt.(*ast.IfStmt)
	if !ok {
		return
	}

	expr, ok := ifstmt.Cond.(*ast.BinaryExpr)
	if !ok {
		return
	}

	if (c.isNilCheck(expr.X) && c.hasLenCall(expr.Y)) ||
		(c.isNilCheck(expr.Y) && c.hasLenCall(expr.X)) {
		c.warn(ifstmt)
	}
}

func (c *lenNilChecker) isNilCheck(expr ast.Expr) bool {
	binexpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return false
	}
	return binexpr.Op == token.NEQ &&
		(binexpr.Y.(*ast.Ident).Name == "nil" && c.isRefType(binexpr.X) ||
			binexpr.X.(*ast.Ident).Name == "nil" && c.isRefType(binexpr.Y))
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

func (c *lenNilChecker) hasLenCall(x ast.Expr) bool {
	return containsNode(x, func(x ast.Node) bool {
		return callQualifiedName(x) == "len"
	})
}

func (c *lenNilChecker) warn(stmt ast.Stmt) {
	c.ctx.Warn(stmt, "consider to remove redundant nil check")
}
