package lint

//! Detects type switches that can benefit from type guard clause with variable.
//
// Before:
// switch v.(type) {
// case int:
// 	return v.(int)
// case point:
// 	return v.(point).x + v.(point).y
// default:
// 	return 0
// }
//
// After:
// switch v := v.(type) {
// case int:
// 	return v
// case point:
// 	return v.x + v.y
// default:
// 	return 0
// }

import (
	"go/ast"

	"github.com/go-toolsmith/astequal"
	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(typeSwitchVarChecker{})
}

type typeSwitchVarChecker struct {
	baseStmtChecker
}

func (c typeSwitchVarChecker) New(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&typeSwitchVarChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

func (c *typeSwitchVarChecker) CheckStmt(stmt ast.Stmt) {
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
	object := c.ctx.TypesInfo.ObjectOf(c.identOf(expr))
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
			if object == c.ctx.TypesInfo.ObjectOf(c.identOf(assert2)) {
				c.warn(root, i)
				break
			}
		}
	}
}

func (c *typeSwitchVarChecker) warn(node ast.Node, caseIndex int) {
	c.ctx.Warn(node, "case %d can benefit from type switch with assignment", caseIndex)
}

// identOf returns identifier for x that can be used to obtain associated types.Object.
// Returns nil for expressions that yield temporary results, like `f().field`.
//
// TODO: is this generally useful and can be placed in util.go?
func (c *typeSwitchVarChecker) identOf(x ast.Node) *ast.Ident {
	switch x := x.(type) {
	case *ast.Ident:
		// Found ident.
		return x
	case *ast.TypeAssertExpr:
		// x.(type) - x may contain ident.
		return c.identOf(x.X)
	case *ast.IndexExpr:
		// x[i] - x may contain ident.
		return c.identOf(x.X)
	case *ast.SelectorExpr:
		// x.y - x may contain ident.
		return c.identOf(x.X)

	default:
		// Note that this function is not comprehensive.
		return nil
	}
}
