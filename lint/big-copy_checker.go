package lint

import (
	"go/ast"
)

// bigCopyCheck finds places where big value copy could be unexpected.
// Detects large value copies in non-testing functions.
//
// Rationale: performance.
func bigCopyCheck(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&bigCopyChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

type bigCopyChecker struct {
	baseStmtChecker
}

func (c *bigCopyChecker) PerFuncInit(fn *ast.FuncDecl) bool {
	return fn.Body != nil && !c.ctx.IsUnitTestFuncDecl(fn)
}

func (c *bigCopyChecker) CheckStmt(stmt ast.Stmt) {
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

func (c *bigCopyChecker) warn(node ast.Node, size int64) {
	c.ctx.Warn(node, "each iteration copies %d bytes (consider pointers or indexing)", size)
}
