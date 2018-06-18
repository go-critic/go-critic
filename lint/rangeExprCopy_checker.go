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

func (c *rangeExprCopyChecker) VisitFunc(fn *ast.FuncDecl) bool {
	return fn.Body != nil && !c.ctx.IsUnitTestFuncDecl(fn)
}

func (c *rangeExprCopyChecker) CheckStmt(stmt ast.Stmt) {
	rng, ok := stmt.(*ast.RangeStmt)
	if !ok || rng.Key == nil || rng.Value == nil {
		return
	}
	tv := c.ctx.TypesInfo.Types[rng.X]
	if !tv.Addressable() {
		return
	}
	if _, ok := tv.Type.(*types.Array); !ok {
		return
	}
	const sizeThreshold = 96 // Not recommended to set value lower than 64
	if size := c.ctx.SizesInfo.Sizeof(tv.Type); size >= sizeThreshold {
		c.warn(rng, size)
	}
}

func (c *rangeExprCopyChecker) warn(rng *ast.RangeStmt, size int64) {
	c.ctx.Warn(rng, "copy of %s (%d bytes) can be avoided with &%s",
		rng.X, size, rng.X)
}
