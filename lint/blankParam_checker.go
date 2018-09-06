package lint

import (
	"go/ast"
)

func init() {
	addChecker(&blankParamChecker{}, attrExperimental, attrVeryOpinionated)
}

type blankParamChecker struct {
	checkerBase
}

func (c *blankParamChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects blank params and suggests to name them as `_` (underscore)"
	d.Before = `func f(a int, b float64) // b isn't used inside function body`
	d.After = `func f(a int, _ float64) // everything is cool`
}

func (c *blankParamChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	params := decl.Type.Params
	if decl.Body == nil || params == nil || params.NumFields() == 0 {
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
	for id := range c.ctx.typesInfo.Uses {
		if _, ok := objToIdent[id.Obj]; ok {
			delete(objToIdent, id.Obj)
			if len(objToIdent) == 0 {
				return
			}
		}
	}

	// all params that are left are blank
	for _, id := range objToIdent {
		c.warn(id)
	}
}

func (c *blankParamChecker) warn(param *ast.Ident) {
	c.ctx.Warn(param, "rename `%s` to `_`", param)
}

func (c *blankParamChecker) warnUnnamed(n ast.Node) {
	c.ctx.Warn(n, "consider to name parameters as `_`")
}
