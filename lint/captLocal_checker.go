package lint

import (
	"go/ast"
	"strings"
)

func init() {
	addChecker(captLocalChecker{}, &ruleInfo{
		SyntaxOnly: true,
	})
}

type captLocalChecker struct {
	baseParamListChecker

	loudNames map[string]bool
}

func (c captLocalChecker) New(ctx *context) func(*ast.File) {
	return wrapParamListChecker(&captLocalChecker{
		baseParamListChecker: baseParamListChecker{ctx: ctx},

		loudNames: map[string]bool{
			"IN":    true,
			"OUT":   true,
			"INOUT": true,

			// TODO: add common acronyms like HTTP and URL?
		},
	})
}

func (c *captLocalChecker) CheckParamList(params []*ast.Field) {
	for _, p := range params {
		for _, id := range p.Names {
			switch {
			case c.loudNames[id.Name]:
				c.warnLoud(id)
			case ast.IsExported(id.Name):
				c.warnCapitalized(id)
			}
		}
	}
}

func (c *captLocalChecker) warnCapitalized(id ast.Node) {
	c.ctx.Warn(id, "`%s' should not be capitalized", id)
}

func (c *captLocalChecker) warnLoud(id *ast.Ident) {
	c.ctx.Warn(id, "consider `%s' name instead of `%s'",
		strings.ToLower(id.Name), id)
}
