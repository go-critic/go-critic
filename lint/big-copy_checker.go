package lint

import (
	"go/ast"
	"go/types"
	"strings"
)

func bigCopyCheck(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&bigCopyChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

type bigCopyChecker struct {
	baseStmtChecker
}

func (c *bigCopyChecker) PerFuncInit(fn *ast.FuncDecl) bool {
	return fn.Body != nil && !c.isUnitTestFunc(fn)
}

func (c *bigCopyChecker) CheckStmt(stmt ast.Stmt) {
	if rng, ok := stmt.(*ast.RangeStmt); ok {
		c.checkRangeStmt(rng)
	}
}

func (c *bigCopyChecker) checkRangeStmt(rng *ast.RangeStmt) {
	if rng.Value == nil {
		return
	}
	const sizeThreshold = 48
	typ := c.ctx.TypesInfo.TypeOf(rng.Value)
	if typ == nil {
		return
	}
	size := c.ctx.SizesInfo.Sizeof(typ)
	if size > sizeThreshold {
		c.warnRangeValue(rng, size)
	}
}

// isUnitTestFunc reports whether FuncDecl declares testing function.
func (c *bigCopyChecker) isUnitTestFunc(fn *ast.FuncDecl) bool {
	if !strings.HasPrefix(fn.Name.Name, "Test") {
		return false
	}
	typ := c.ctx.TypesInfo.TypeOf(fn.Name)
	if sig, ok := typ.(*types.Signature); ok {
		return sig.Results().Len() == 0 &&
			sig.Params().Len() == 1
	}
	return false
}

func (c *bigCopyChecker) warnRangeValue(node ast.Node, size int64) {
	c.ctx.Warn(node, "each iteration copies %d bytes (consider pointers or indexing)", size)
}
