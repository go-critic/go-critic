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
	// TODO(quasilyte): do recursive analysis. Type switches
	// may have other type switches inside their clauses that
	// can trigger warnings as well.
	for _, decl := range collectFuncDecls(f) {
		if decl.Body == nil {
			continue
		}
		ast.Inspect(decl.Body, func(x ast.Node) bool {
			if stmt, ok := x.(*ast.TypeSwitchStmt); ok {
				c.check(stmt)
				return false
			}
			return true
		})
	}
	return c.warnings
}

func (c *TypeGuardChecker) check(root *ast.TypeSwitchStmt) {
	if _, ok := root.Assign.(*ast.ExprStmt); !ok {
		return
	}
	assert, ok := root.Assign.(*ast.ExprStmt).X.(*ast.TypeAssertExpr)
	if !ok {
		return
	}

	for i, clause := range root.Body.List {
		clause := clause.(*ast.CaseClause)
		// Multiple types in a list mean that assert.X will have
		// a type of interface{} inside body.
		// We are looking for precise type case.
		if len(clause.List) != 1 {
			continue
		}
		typ := clause.List[0]
		for _, stmt := range clause.Body {
			if c.findTypeAssert(stmt, assert.X, typ) {
				c.warn(root, i)
				break
			}
		}
	}
}

func (c *TypeGuardChecker) findTypeAssert(root ast.Stmt, expr, typ ast.Expr) bool {
	found := false
	// TODO(quasilyte): inspect does not end the traverse after false is returned.
	ast.Inspect(root, func(x ast.Node) bool {
		if assert, ok := x.(*ast.TypeAssertExpr); ok {
			found = astcmp.EqualExpr(expr, assert.X) &&
				astcmp.EqualExpr(typ, assert.Type)
			return !found
		}
		return true
	})
	return found
}

func (c *TypeGuardChecker) warn(stmt *ast.TypeSwitchStmt, caseIndex int) {
	s := "case %d can benefit from type switch with assignment"
	c.warnings = append(c.warnings, Warning{
		Node: stmt,
		Text: fmt.Sprintf(s, caseIndex),
	})
}
