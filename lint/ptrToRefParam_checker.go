package lint

//! Detects input and output parameters that have a type of pointer to referential type.
//
// Before:
// func f(m *map[string]int) (ch *chan *int)
//
// After:
// func f(m map[string]int) (ch chan *int)
//
// Note:
// > Slices are not as referential as maps or channels, but it's usually
// > better to return them by value rather than modyfing them by pointer.

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(&ptrToRefParamChecker{})
}

type ptrToRefParamChecker struct {
	baseFuncDeclChecker
}

func (c *ptrToRefParamChecker) New(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&ptrToRefParamChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

func (c *ptrToRefParamChecker) CheckFuncDecl(fn *ast.FuncDecl) {
	c.checkParams(fn.Type.Params.List)
	if fn.Type.Results != nil {
		c.checkParams(fn.Type.Results.List)
	}
}

func (c *ptrToRefParamChecker) checkParams(params []*ast.Field) {
	for _, param := range params {
		ptr, ok := c.ctx.TypesInfo.TypeOf(param.Type).(*types.Pointer)
		if !ok {
			continue
		}

		if c.isRefType(ptr.Elem().Underlying()) {
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

func (c *ptrToRefParamChecker) isRefType(x types.Type) bool {
	switch x.(type) {
	case *types.Map, *types.Chan, *types.Slice:
		return true
	default:
		return false
	}
}

func (c *ptrToRefParamChecker) warn(id *ast.Ident) {
	c.ctx.Warn(id, "consider `%s' to be of non-pointer type", id)
}
