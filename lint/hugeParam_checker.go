package lint

import (
	"go/ast"
)

//! Detects params that incur excessive amount of copying.
//
// @Before:
// func f(x [1024]int) {}
//
// @After:
// func f(x *[1024]int) {}

func init() {
	addChecker(&hugeParamChecker{}, attrExperimental)
}

type hugeParamChecker struct {
	checkerBase
}

func (c *hugeParamChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	// TODO(quasilyte): maybe it's worthwhile to permit skipping
	// test files for this checker?
	if decl.Recv != nil {
		c.checkParams(decl, decl.Recv.List)
	}
	c.checkParams(decl, decl.Type.Params.List)
}

func (c *hugeParamChecker) checkParams(decl *ast.FuncDecl, params []*ast.Field) {
	const sizeThreshold = 80 // TODO(quasilyte): should be configurable
	for _, p := range params {
		for _, id := range p.Names {
			typ := c.ctx.typesInfo.TypeOf(id)
			size := c.ctx.sizesInfo.Sizeof(typ)
			if size >= sizeThreshold {
				c.warn(id, size)
			}
		}
	}
}

func (c *hugeParamChecker) warn(cause *ast.Ident, size int64) {
	c.ctx.Warn(cause, "%s is heavy (%d bytes); consider passing it by pointer",
		cause, size)
}
