package lint

import (
	"go/ast"
)

func init() {
	addChecker(&hugeParamChecker{}, attrExperimental)
}

type hugeParamChecker struct {
	checkerBase

	sizeThreshold int64
}

func (c *hugeParamChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects params that incur excessive amount of copying"
	d.Before = `func f(x [1024]int) {}`
	d.After = `func f(x *[1024]int) {}`
}

func (c *hugeParamChecker) Init() {
	c.sizeThreshold = int64(c.ctx.params.Int("sizeThreshold", 80))
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
	for _, p := range params {
		for _, id := range p.Names {
			typ := c.ctx.typesInfo.TypeOf(id)
			size := c.ctx.sizesInfo.Sizeof(typ)
			if size >= c.sizeThreshold {
				c.warn(id, size)
			}
		}
	}
}

func (c *hugeParamChecker) warn(cause *ast.Ident, size int64) {
	c.ctx.Warn(cause, "%s is heavy (%d bytes); consider passing it by pointer",
		cause, size)
}
