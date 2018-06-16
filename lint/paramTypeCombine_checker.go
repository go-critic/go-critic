package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astequal"
)

func init() {
	addChecker(paramTypeCombineChecker{}, &ruleInfo{})
}

type paramTypeCombineChecker struct {
	baseFuncDeclChecker
}

func (c paramTypeCombineChecker) New(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&paramTypeCombineChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

func (c *paramTypeCombineChecker) CheckFuncDecl(decl *ast.FuncDecl) {
	typ := c.optimizeFuncType(decl.Type)
	if !astequal.Expr(typ, decl.Type) {
		c.warn(decl.Type, typ)
	}
}

func (c *paramTypeCombineChecker) optimizeFuncType(f *ast.FuncType) *ast.FuncType {
	return &ast.FuncType{
		Params:  c.optimizeParams(f.Params),
		Results: c.optimizeParams(f.Results),
	}
}
func (c *paramTypeCombineChecker) optimizeParams(params *ast.FieldList) *ast.FieldList {
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
		if astequal.Expr(p.Type, params.List[i].Type) {
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

func (c *paramTypeCombineChecker) warn(f1, f2 *ast.FuncType) {
	c.ctx.Warn(f1, "%s could be replaced with %s", f1, f2)
}
