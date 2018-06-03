package lint

import (
	"go/ast"

	"github.com/Quasilyte/astcmp"
	"golang.org/x/tools/go/ast/astutil"
)

func typeGuardCheck(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&typeGuardChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

type typeGuardChecker struct {
	baseStmtChecker
}

func (c *typeGuardChecker) CheckStmt(stmt ast.Stmt) {
	if stmt, ok := stmt.(*ast.TypeSwitchStmt); ok {
		c.checkTypeSwitch(stmt)
	}
}

func (c *typeGuardChecker) checkTypeSwitch(root *ast.TypeSwitchStmt) {
	if _, ok := root.Assign.(*ast.AssignStmt); ok {
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
			assert2 := c.find(stmt, func(x ast.Node) bool {
				return astcmp.Equal(&assert1, x)
			})
			if object == c.ctx.TypesInfo.ObjectOf(c.identOf(assert2)) {
				c.warn(root, i)
				break
			}
		}
	}
}

func (c *typeGuardChecker) warn(node ast.Node, caseIndex int) {
	c.ctx.Warn(node, "case %d can benefit from type switch with assignment", caseIndex)
}

// find applies pred for root and all it's childs until it returns true.
// Matched node is returned.
// If none of the nodes matched predicate, nil is returned.
//
// TODO: is this generally useful and can be placed in util.go?
func (c *typeGuardChecker) find(root ast.Node, pred func(ast.Node) bool) ast.Node {
	var found ast.Node
	astutil.Apply(root, nil, func(cur *astutil.Cursor) bool {
		if pred(cur.Node()) {
			found = cur.Node()
			return false
		}
		return true
	})
	return found
}

// identOf returns identifier for x that can be used to obtain associated types.Object.
// Returns nil for expressions that yield temporary results, like `f().field`.
//
// TODO: is this generally useful and can be placed in util.go?
func (c *typeGuardChecker) identOf(x ast.Node) *ast.Ident {
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
