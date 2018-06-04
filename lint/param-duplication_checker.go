package lint

import (
	"go/ast"

	"github.com/Quasilyte/astcmp"
)

func paramDuplicationCheck(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&paramDuplicationChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

type paramDuplicationChecker struct {
	baseFuncDeclChecker
}

func (c *paramDuplicationChecker) CheckFuncDecl(decl *ast.FuncDecl) {
	typ := c.optimizeFuncType(decl.Type)
	if !astcmp.EqualExpr(typ, decl.Type) {
		c.warn(decl.Type, typ)
	}
}

func (c *paramDuplicationChecker) optimizeFuncType(f *ast.FuncType) *ast.FuncType {
	return &ast.FuncType{
		Params:  c.optimizeParams(f.Params),
		Results: c.optimizeParams(f.Results),
	}
}
func (c *paramDuplicationChecker) optimizeParams(params *ast.FieldList) *ast.FieldList {
	if params == nil || len(params.List) < 2 {
		return params
	}

	list := []*ast.Field{}
	names := make([]*ast.Ident, len(params.List[0].Names))
	copy(names, params.List[0].Names)
	list = append(list, &ast.Field{
		Names: names,
		Type:  params.List[0].Type,
	})
	for i, p := range params.List[1:] {
		names = make([]*ast.Ident, len(p.Names))
		copy(names, p.Names)
		if astcmp.EqualExpr(p.Type, params.List[i].Type) {
			list[len(list)-1].Names =
				append(list[len(list)-1].Names, names...)
		} else {
			list = append(list, &ast.Field{
				Names: names,
				Type:  params.List[i+1].Type,
			})
		}
	}
	return &ast.FieldList{
		List: list,
	}
}

func (c *paramDuplicationChecker) warn(f1, f2 *ast.FuncType) {
	c.ctx.Warn(f1, "%s could be replaced with %s",
		nodeString(c.ctx.FileSet, f1),
		nodeString(c.ctx.FileSet, f2))
}
