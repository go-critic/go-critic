package lint

import (
	"fmt"
	"go/ast"

	"github.com/Quasilyte/astcmp"
)

// TypeGuardChecker finds type switches that may benefit from type guard clause.
//
// Rationale: code readability.
type TypeGuardChecker struct {
	ctx *Context

	warnings []Warning
}

// NewTypeGuardChecker returns initialized checker for Go type switch statements.
func NewTypeGuardChecker(ctx *Context) *TypeGuardChecker {
	return &TypeGuardChecker{ctx: ctx}
}

// Check runs type switch checks for f.
func (c *TypeGuardChecker) Check(f *ast.File) []Warning {
	c.warnings = c.warnings[:0]
	inspectFuncBodies(f, c.inspectNode)
	return c.warnings
}

func (c *TypeGuardChecker) inspectNode(x ast.Node) bool {
	if _, ok := x.(ast.Stmt); !ok {
		return false
	}
	if stmt, ok := x.(*ast.TypeSwitchStmt); ok {
		c.checkTypeSwitch(stmt)
	}
	return true
}

func (c *TypeGuardChecker) checkTypeSwitch(root *ast.TypeSwitchStmt) {
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
		// Create artifical node just for matching.
		assert1 := ast.TypeAssertExpr{X: expr, Type: clause.List[0]}
		for _, stmt := range clause.Body {
			assert2 := c.findEqual(stmt, &assert1)
			if object == c.ctx.TypesInfo.ObjectOf(c.identOf(assert2)) {
				c.warn(root, i)
				break
			}
		}
	}
}

// findEqual walks tree root trying to find node that is equal to x.
// Matched node is returned.
func (c *TypeGuardChecker) findEqual(root, x ast.Node) (found ast.Node) {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

	ast.Inspect(root, func(y ast.Node) bool {
		if x == nil {
			return false
		}
		if astcmp.Equal(x, y) {
			found = y
			panic(nil)
		}
		return true
	})

	return nil // Non-nil return only happens if there was a panic
}

// identOf returns identifier for x that can be used to obtain associated types.Object.
// Returns nil for expressions that yield temporary results, like `f().field`.
func (c *TypeGuardChecker) identOf(x ast.Node) *ast.Ident {
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

func (c *TypeGuardChecker) warn(stmt *ast.TypeSwitchStmt, caseIndex int) {
	s := "case %d can benefit from type switch with assignment"
	c.warnings = append(c.warnings, Warning{
		Node: stmt,
		Text: fmt.Sprintf(s, caseIndex),
	})
}
