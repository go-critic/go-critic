package lint

//! Detects loops that copy big objects during each iteration.
// Suggests to use index access or take address and make use pointer instead.
//
// Before:
// xs := make([][1024]byte, length)
// for _, x := range xs {
// 	// Loop body.
// }
//
// After:
// xs := make([][1024]byte, length)
// for i := range xs {
// 	x := &xs[i]
// 	// Loop body.
// }

import (
	"go/ast"
)

func init() {
	addChecker(&rangeValCopyChecker{})
}

type rangeValCopyChecker struct {
	checkerBase
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
	const sizeThreshold = 48
	if size := c.ctx.sizesInfo.Sizeof(typ); size >= sizeThreshold {
		c.warn(rng, size)
	}
}

func (c *rangeValCopyChecker) warn(node ast.Node, size int64) {
	c.ctx.Warn(node, "each iteration copies %d bytes (consider pointers or indexing)", size)
}
