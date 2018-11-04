package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astequal"
	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(&typeSwitchVarChecker{})
}

type typeSwitchVarChecker struct {
	checkerBase
}

func (c *typeSwitchVarChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects type switches that can benefit from type guard clause with variable"
	d.Before = `
switch v.(type) {
case int:
	return v.(int)
case point:
	return v.(point).x + v.(point).y
default:
	return 0
}`
	d.After = `
switch v := v.(type) {
case int:
	return v
case point:
	return v.x + v.y
default:
	return 0
}`
}

func (c *typeSwitchVarChecker) VisitStmt(stmt ast.Stmt) {
	if stmt, ok := stmt.(*ast.TypeSwitchStmt); ok {
		c.checkTypeSwitch(stmt)
	}
}

func (c *typeSwitchVarChecker) checkTypeSwitch(root *ast.TypeSwitchStmt) {
	if astp.IsAssignStmt(root.Assign) {
		return // Already with type guard
	}
	// Must be a *ast.ExprStmt then.
	expr := root.Assign.(*ast.ExprStmt).X.(*ast.TypeAssertExpr).X
	object := c.ctx.typesInfo.ObjectOf(identOf(expr))
	if object == nil {
		return // Give up: can't handle shadowing without object
	}

	for i, clause := range root.Body.List {
		clause := clause.(*ast.CaseClause)
		// Multiple types in a list mean that assert.X will have
		// a type of interface{} inside clause body.
		// We are looking for precise type case.
		if len(clause.List) != 1 {
			continue
		}
		// Create artificial node just for matching.
		assert1 := ast.TypeAssertExpr{X: expr, Type: clause.List[0]}
		for _, stmt := range clause.Body {
			assert2 := findNode(stmt, func(x ast.Node) bool {
				return astequal.Node(&assert1, x)
			})
			if object == c.ctx.typesInfo.ObjectOf(identOf(assert2)) {
				c.warn(root, i)
				break
			}
		}
	}
}

func (c *typeSwitchVarChecker) warn(cause ast.Node, caseIndex int) {
	c.ctx.Warn(cause, "case %d can benefit from type switch with assignment", caseIndex)
}
