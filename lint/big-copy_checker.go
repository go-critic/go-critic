package lint

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
)

type bigCopyChecker struct {
	ctx *Context
}

func newBigCopyChecker(ctx *Context) Checker {
	return &bigCopyChecker{ctx: ctx}
}

// Check finds places where big value copy could be unexpected.
//
// Features
//
// Detects large value copies in non-testing functions.
func (c *bigCopyChecker) Check(f *ast.File) {
	for _, decl := range collectFuncDecls(f) {
		if c.isUnitTestFunc(decl) {
			continue
		}
		ast.Inspect(decl, func(x ast.Node) bool {
			if rng, ok := x.(*ast.RangeStmt); ok {
				c.checkRangeStmt(rng)
			}
			return true
		})
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

// isUnitTestFunc reports whether decl declares testing function.
func (c *bigCopyChecker) isUnitTestFunc(decl *ast.FuncDecl) bool {
	if !strings.HasPrefix(decl.Name.Name, "Test") {
		return false
	}
	typ := c.ctx.TypesInfo.TypeOf(decl.Name)
	if sig, ok := typ.(*types.Signature); ok {
		return sig.Results().Len() == 0 &&
			sig.Params().Len() == 1
	}
	return false
}

func (c *bigCopyChecker) warnRangeValue(node ast.Node, size int64) {
	c.ctx.addWarning(Warning{
		Kind: "big-copy/RangeValue",
		Node: node,
		Text: fmt.Sprintf("each iteration copies %d bits (consider pointers or indexing)", size),
	})
}
