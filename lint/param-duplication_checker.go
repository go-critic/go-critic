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
	opt := c.optimiseFuncType(decl.Type)
	if !astcmp.EqualExpr(opt, decl.Type) {
		c.warn(decl.Type, opt)
	}
}

func (c *paramDuplicationChecker) optimiseFuncType(f *ast.FuncType) *ast.FuncType {
	res := c.optimiseParams(f.Results)
	par := c.optimiseParams(f.Params)
	return &ast.FuncType{
		Params:  par,
		Results: res,
	}
}
func (c *paramDuplicationChecker) optimiseParams(params *ast.FieldList) *ast.FieldList {
	if params == nil || len(params.List) < 2 {
		return params
	}

	newParams := &ast.FieldList{
		List: []*ast.Field{},
	}
	names := make([]*ast.Ident, len(params.List[0].Names))
	copy(names, params.List[0].Names)
	newParams.List = append(newParams.List, &ast.Field{
		Names: names,
		Type:  params.List[0].Type,
	})
	for i, p := range params.List[1:] {
		names = make([]*ast.Ident, len(p.Names))
		copy(names, p.Names)
		if astcmp.EqualExpr(p.Type, params.List[i].Type) {
			newParams.List[len(newParams.List)-1].Names =
				append(newParams.List[len(newParams.List)-1].Names, names...)
		} else {
			newParams.List = append(newParams.List, &ast.Field{
				Names: names,
				Type:  params.List[i+1].Type,
			})
		}
	}
	return newParams
}

func (c *paramDuplicationChecker) warn(f1, f2 *ast.FuncType) {
	c.ctx.Warn(f1, "%s could be replaced with %s",
		nodeString(c.ctx.FileSet, f1),
		nodeString(c.ctx.FileSet, f2))
}
