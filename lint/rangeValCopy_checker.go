package lint

import (
	"go/ast"
)

func init() {
	addChecker(&rangeValCopyChecker{}, attrPerformance)
}

type rangeValCopyChecker struct {
	checkerBase

	sizeThreshold int64
}

func (c *rangeValCopyChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects loops that copy big objects during each iteration"
	d.Details = "Suggests to use index access or take address and make use pointer instead."
	d.Before = `
xs := make([][1024]byte, length)
for _, x := range xs {
	// Loop body.
}`
	d.After = `
xs := make([][1024]byte, length)
for i := range xs {
	x := &xs[i]
	// Loop body.
}`
}

func (c *rangeValCopyChecker) Init() {
	c.sizeThreshold = int64(c.ctx.params.Int("sizeThreshold", 128))
}

func (c *rangeValCopyChecker) EnterFunc(fn *ast.FuncDecl) bool {
	return fn.Body != nil && !c.ctx.IsUnitTestFuncDecl(fn)
}

func (c *rangeValCopyChecker) VisitStmt(stmt ast.Stmt) {
	rng, ok := stmt.(*ast.RangeStmt)
	if !ok || rng.Value == nil {
		return
	}
	typ := c.ctx.typesInfo.TypeOf(rng.Value)
	if typ == nil {
		return
	}
	if size := c.ctx.sizesInfo.Sizeof(typ); size >= c.sizeThreshold {
		c.warn(rng, size)
	}
}

func (c *rangeValCopyChecker) warn(cause ast.Node, size int64) {
	c.ctx.Warn(cause, "each iteration copies %d bytes (consider pointers or indexing)", size)
}
