package lint

import (
	"go/ast"
	"go/types"
)

// unsliceChecker finds slice expressions that can be
// simplified to sliced expression.
func unsliceCheck(ctx *context) func(*ast.File) {
	return wrapLocalExprChecker(&unsliceChecker{
		baseLocalExprChecker: baseLocalExprChecker{ctx: ctx},
	})
}

type unsliceChecker struct {
	baseLocalExprChecker
}

func (c *unsliceChecker) CheckLocalExpr(expr ast.Expr) {
	if expr, ok := expr.(*ast.SliceExpr); ok {
		if expr.Low == nil && expr.High == nil {
			typ := c.ctx.TypesInfo.TypeOf(expr)
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
	c.ctx.Warn(expr, "could simplify %s to %s",
		nodeString(c.ctx.FileSet, expr),
		nodeString(c.ctx.FileSet, expr.X))
}
