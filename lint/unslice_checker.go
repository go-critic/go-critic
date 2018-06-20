package lint

//! Detects slice expressions that can be simplified to sliced expression itself.
//
// Before:
// f(s[:]) // s is string
// copy(b[:], values...) // b is []byte
//
// After:
// f(s)
// copy(b, values...)

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(&unsliceChecker{})
}

type unsliceChecker struct {
	checkerBase
}

func (c *unsliceChecker) VisitLocalExpr(expr ast.Expr) {
	if expr, ok := expr.(*ast.SliceExpr); ok {
		if expr.Low == nil && expr.High == nil {
			typ := c.ctx.typesInfo.TypeOf(expr)
			switch typ := typ.(type) {
			case *types.Basic:
				if typ.Kind() == types.String {
					c.warn(expr)
				}
			}
		}
	}
}

func (c *unsliceChecker) warn(expr *ast.SliceExpr) {
	c.ctx.Warn(expr, "could simplify %s to %s", expr, expr.X)
}
