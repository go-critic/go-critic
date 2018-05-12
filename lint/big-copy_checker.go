package lint

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
)

// BigCopyChecker finds places where big value copy could be unexpected.
//
// Rationale: performance.
type BigCopyChecker struct {
	ctx *Context

	warnings []Warning
}

// NewBigCopyChecker returns initialized checker for range statements.
func NewBigCopyChecker(ctx *Context) *BigCopyChecker {
	return &BigCopyChecker{ctx: ctx}
}

// Check runs range statement inspections for f.
//
// Features
//
// Detects large value copies in non-testing functions.
func (c *BigCopyChecker) Check(f *ast.File) []Warning {
	c.warnings = c.warnings[:0]
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

	return c.warnings
}

func (c *BigCopyChecker) checkRangeStmt(rng *ast.RangeStmt) {
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
func (c *BigCopyChecker) isUnitTestFunc(decl *ast.FuncDecl) bool {
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

func (c *BigCopyChecker) warnRangeValue(node ast.Node, size int64) {
	c.warnings = append(c.warnings, Warning{
		Kind: "RangeValue",
		Node: node,
		Text: fmt.Sprintf("each iteration copies %d bits (consider pointers or indexing)", size),
	})
}
