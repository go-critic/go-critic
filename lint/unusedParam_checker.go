package lint

import (
	"go/ast"
)

func init() {
	addChecker(&unusedParamChecker{}, attrExperimental)
}

type unusedParamChecker struct {
	checkerBase
}

func (c *unusedParamChecker) CheckFuncDecl(decl *ast.FuncDecl) {
	params := decl.Type.Params
	if params == nil || params.NumFields() == 0 {
		return
	}

	// collect all params to map
	objToIdent := make(map[*ast.Object]*ast.Ident)
	for _, p := range params.List {
		if len(p.Names) == 0 {
			c.warnUnnamed(p)
			return
		}
		for _, id := range p.Names {
			if id.Name != "_" {
				objToIdent[id.Obj] = id
			}
		}
	}

	// remove used params
	// TODO(cristaloleg): we might want to have less iterations here.
	for id := range c.ctx.TypesInfo.Uses {
		if _, ok := objToIdent[id.Obj]; ok {
			delete(objToIdent, id.Obj)
		}
	}

	// all params that are left are unused
	for _, id := range objToIdent {
		c.warn(id)
	}
}

func (c *unusedParamChecker) warn(param *ast.Ident) {
	c.ctx.Warn(param, "parameter `%s` isn't used, consider to name it as `_`", param)
}

func (c *unusedParamChecker) warnUnnamed(n ast.Node) {
	c.ctx.Warn(n, "consider to name parameters as `_`")
}
