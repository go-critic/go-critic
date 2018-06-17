package lint

//! Detects expensive copies of `for` loop range expressions.
// Suggests to use pointer to array to avoid the copy using `&` on range expression.
//
// Before:
// var xs [256]byte
// for _, x := range xs {
// 	// Loop body.
// }
//
// After:
// var xs [256]byte
// for _, x := range &xs {
// 	// Loop body.
// }

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(rangeExprCopyChecker{})
}

type rangeExprCopyChecker struct {
	baseStmtChecker
}

func (c rangeExprCopyChecker) New(ctx *context) func(*ast.File) {
	// TODO(quasilyte): there is some annoying code duplication with other
	// range statement checker. We should consider refactoring if
	// more checkers that inspect range statements will appear.
	return wrapStmtChecker(&rangeExprCopyChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

func (c *rangeExprCopyChecker) PerFuncInit(fn *ast.FuncDecl) bool {
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
