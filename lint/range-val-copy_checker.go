package lint

import (
	"go/ast"
)

func rangeValCopyCheck(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&rangeValCopyChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

type rangeValCopyChecker struct {
	baseStmtChecker
}

func (c *rangeValCopyChecker) PerFuncInit(fn *ast.FuncDecl) bool {
	return fn.Body != nil && !c.ctx.IsUnitTestFuncDecl(fn)
}

func (c *rangeValCopyChecker) CheckStmt(stmt ast.Stmt) {
	rng, ok := stmt.(*ast.RangeStmt)
	if !ok || rng.Value == nil {
		return
	}
	typ := c.ctx.TypesInfo.TypeOf(rng.Value)
	if typ == nil {
		return
	}
	const sizeThreshold = 48
	if size := c.ctx.SizesInfo.Sizeof(typ); size >= sizeThreshold {
		c.warn(rng, size)
	}
}

func (c *rangeValCopyChecker) warn(node ast.Node, size int64) {
	c.ctx.Warn(node, "each iteration copies %d bytes (consider pointers or indexing)", size)
}
