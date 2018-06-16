package lint

import (
	"go/ast"
)

func init() {
	addChecker(unusedParamChecker{}, &ruleInfo{})
}

type unusedParamChecker struct {
	baseFuncDeclChecker
}

func (c unusedParamChecker) New(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&unusedParamChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

func (c *unusedParamChecker) CheckFuncDecl(decl *ast.FuncDecl) {
	params := decl.Type.Params
	if params == nil || params.NumFields() == 0 {
		return
	}

	objUsed := make(map[*ast.Object]struct{})
	for id := range c.ctx.TypesInfo.Uses {
		objUsed[id.Obj] = struct{}{}
	}

	for _, p := range params.List {
		if len(p.Names) == 0 {
			c.warnUnnamed(p)
			break
		}
		for _, id := range p.Names {
			if id.Name == "_" {
				continue
			}
			if _, ok := objUsed[id.Obj]; !ok {
				c.warn(id)
			}
		}
	}
}

func (c *unusedParamChecker) warn(param *ast.Ident) {
	c.ctx.Warn(param, "parameter `%s` isn't used, consider to name it as `_`", param)
}

func (c *unusedParamChecker) warnUnnamed(n ast.Node) {
	c.ctx.Warn(n, "consider to name parameters as `_`")
}
