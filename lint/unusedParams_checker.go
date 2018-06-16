package lint

import (
	"fmt"
	"go/ast"
)

func init() {
	addChecker(unusedParamsChecker{}, &ruleInfo{})
}

type unusedParamsChecker struct {
	baseFuncDeclChecker
}

func (c unusedParamsChecker) New(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&unusedParamsChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

func (c *unusedParamsChecker) CheckFuncDecl(decl *ast.FuncDecl) {
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
		for _, name := range p.Names {
			if name.Name == "_" {
				continue
			}
			if _, ok := objUsed[name.Obj]; !ok {
				c.warn(name, name.Name)
			}
		}
	}
}

func (c *unusedParamsChecker) warn(n ast.Node, param string) {
	c.ctx.Warn(n, fmt.Sprintf("parameter `%s` isn't used, consider to name it as `_`", param))
}

func (c *unusedParamsChecker) warnUnnamed(n ast.Node) {
	c.ctx.Warn(n, "consider to name parameters as `_`")
}
