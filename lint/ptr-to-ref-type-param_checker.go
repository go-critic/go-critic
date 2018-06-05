package lint

import (
	"go/ast"
	"go/types"
)

func ptrToRefTypeParamCheck(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&ptrToRefTypeParamChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

type ptrToRefTypeParamChecker struct {
	baseFuncDeclChecker
}

func (c *ptrToRefTypeParamChecker) CheckFuncDecl(fn *ast.FuncDecl) {
	c.checkParams(fn.Type.Params.List)
	if fn.Type.Results != nil {
		c.checkParams(fn.Type.Results.List)
	}
}

func (c *ptrToRefTypeParamChecker) checkParams(params []*ast.Field) {
	for _, param := range params {
		ptr, ok := c.ctx.TypesInfo.TypeOf(param.Type).(*types.Pointer)
		if !ok {
			continue
		}

		if isRefType(ptr.Elem().Underlying()) {
			if len(param.Names) == 0 {
				c.ctx.Warn(param, "consider to make non-pointer type for `%s`", ptr.String())
			} else {
				for i := range param.Names {
					c.warn(param.Names[i])
				}
			}
		}
	}
}

func (c *ptrToRefTypeParamChecker) warn(id *ast.Ident) {
	c.ctx.Warn(id, "consider `%s' to be of non-pointer type", id)
}

func isRefType(x types.Type) bool {
	switch x.(type) {
	case *types.Map, *types.Chan, *types.Slice:
		return true
	default:
		return false
	}
}
