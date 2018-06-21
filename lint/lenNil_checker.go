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
)

func init() {
	addChecker(&lenNilChecker{}, attrExperimental)
}

type lenNilChecker struct {
	checkerBase
}

func (c *lenNilChecker) VisitStmt(stmt ast.Stmt) {
	ifstmt, ok := stmt.(*ast.IfStmt)
	if !ok {
		return
	}

	visited := make(map[ast.Node]struct{})

	findNode(ifstmt, func(node ast.Node) bool {
		expr, ok := ifstmt.Cond.(*ast.BinaryExpr)
		if !ok {
			return false
		}

		if _, ok := visited[expr]; ok {
			return false
		}
		visited[expr] = struct{}{}

		if expr.Op == token.LAND || expr.Op == token.LOR {
			if c.isNilCheck(expr.X) && c.hasLenCall(expr.Y) {
				c.warn(expr.X)
			}
			if c.isNilCheck(expr.Y) && c.hasLenCall(expr.X) {
				c.warn(expr.Y)
			}
		}
		return false
	})
}

func (c *lenNilChecker) isNilCheck(expr ast.Expr) bool {
	binexpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return false
	}
	return (binexpr.Op == token.NEQ || binexpr.Op == token.EQL) &&
		(c.isRefType(binexpr.X) && c.isNilKeyword(binexpr.Y) ||
			c.isRefType(binexpr.Y) && c.isNilKeyword(binexpr.X))
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

func (c *lenNilChecker) isNilKeyword(x ast.Expr) bool {
	id, ok := x.(*ast.Ident)
	return ok && id.Name == "nil"
}

func (c *lenNilChecker) hasLenCall(x ast.Expr) bool {
	return containsNode(x, func(x ast.Node) bool {
		return callQualifiedName(x) == "len"
	})
}

func (c *lenNilChecker) warn(expr ast.Expr) {
	c.ctx.Warn(expr, "%s check is redundant", expr)
}
