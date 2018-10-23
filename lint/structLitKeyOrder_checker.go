package lint

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/lint/internal/lintutil"
	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(&structLitKeyOrderChecker{}, attrVeryOpinionated, attrExperimental)
}

type structLitKeyOrderChecker struct {
	checkerBase
}

func (c *structLitKeyOrderChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects struct literal keys order that does not match declaration order"
	d.Before = `
type foo struct{ x, y int }
v := foo{y: y, x: x}`
	d.After = `
type foo struct{ x, y int }
v := foo{x: x, y: y}`
}

func (c *structLitKeyOrderChecker) VisitExpr(expr ast.Expr) {
	lit, ok := expr.(*ast.CompositeLit)
	if !ok || len(lit.Elts) <= 1 {
		// If there is 1 or less elements, consider the
		// list as sorted. This also makes addressing
		// the first element below safe.
		return
	}
	if !astp.IsKeyValueExpr(lit.Elts[0]) {
		// Skip literals without keyed field initializers.
		// If at least 1 field is initialized using keyed
		// syntax, all of them are guaranteed to be
		// initialized by a KeyValueExpr.
		return
	}
	typ, ok := c.ctx.typesInfo.TypeOf(lit).Underlying().(*types.Struct)
	if !ok {
		return
	}

	// Collect field name to declaration order mapping.
	// There are other solutions that do not involve allocated
	// map, but the current implementation is fast enough.
	fieldsOrder := make(map[string]int)
	for i := 0; i < typ.NumFields(); i++ {
		fieldsOrder[typ.Field(i).Name()] = i
	}

	// Check until we find a field that is misplaced.
	// Misplaced field is a field that comes after a field
	// that had higher declaration order value.
	maxOrder := -1
	for _, elt := range lit.Elts {
		kv := elt.(*ast.KeyValueExpr)
		order := fieldsOrder[lintutil.AsIdent(kv.Key).Name]
		if order > maxOrder {
			maxOrder = order
		} else {
			c.warn(expr)
			return
		}
	}
}

func (c *structLitKeyOrderChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "literal key order does not match declaration key order")
}
