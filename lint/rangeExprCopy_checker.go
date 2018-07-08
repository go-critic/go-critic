package lint

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(&rangeExprCopyChecker{})
}

type rangeExprCopyChecker struct {
	checkerBase
}

func (c *rangeExprCopyChecker) InitDocs(d *Documentation) {
	d.Summary = "Detects expensive copies of `for` loop range expressions"
	d.Details = "Suggests to use pointer to array to avoid the copy using `&` on range expression."
	d.Before = `
var xs [256]byte
for _, x := range xs {
	// Loop body.
}`
	d.After = `
var xs [256]byte
for _, x := range &xs {
	// Loop body.
}`
}

func (c *rangeExprCopyChecker) EnterFunc(fn *ast.FuncDecl) bool {
	return fn.Body != nil && !c.ctx.IsUnitTestFuncDecl(fn)
}

func (c *rangeExprCopyChecker) VisitStmt(stmt ast.Stmt) {
	rng, ok := stmt.(*ast.RangeStmt)
	if !ok || rng.Key == nil || rng.Value == nil {
		return
	}
	tv := c.ctx.typesInfo.Types[rng.X]
	if !tv.Addressable() {
		return
	}
	if _, ok := tv.Type.(*types.Array); !ok {
		return
	}
	const sizeThreshold = 96 // Not recommended to set value lower than 64
	if size := c.ctx.sizesInfo.Sizeof(tv.Type); size >= sizeThreshold {
		c.warn(rng, size)
	}
}

func (c *rangeExprCopyChecker) warn(rng *ast.RangeStmt, size int64) {
	c.ctx.Warn(rng, "copy of %s (%d bytes) can be avoided with &%s",
		rng.X, size, rng.X)
}
