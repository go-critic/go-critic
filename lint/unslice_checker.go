package lint

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

func (c *unsliceChecker) InitDocs(d *Documentation) {
	d.Summary = "Detects slice expressions that can be simplified to sliced expression itself"
	d.Before = `
f(s[:])               // s is string
copy(b[:], values...) // b is []byte`
	d.After = `
f(s)
copy(b, values...)`
}

func (c *unsliceChecker) VisitLocalExpr(expr ast.Expr) {
	if expr, ok := expr.(*ast.SliceExpr); ok {
		// No need to worry about 3-index slicing,
		// because it's only permitted if expr.High is not nil.
		if expr.Low != nil || expr.High != nil {
			return
		}
		switch c.ctx.typesInfo.TypeOf(expr.X).(type) {
		case *types.Slice, *types.Basic:
			// Basic kind catches strings, Slice cathes everything else.
			c.warn(expr)
		}
	}
}

func (c *unsliceChecker) warn(expr *ast.SliceExpr) {
	c.ctx.Warn(expr, "could simplify %s to %s", expr, expr.X)
}
