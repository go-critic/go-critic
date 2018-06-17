package lint

import (
	"go/ast"
	"strings"
)

func init() {
	addChecker(captLocalChecker{}, attrSyntaxOnly)
}

type captLocalChecker struct {
	baseLocalNameChecker

	loudNames map[string]bool
}

func (c captLocalChecker) New(ctx *context) func(*ast.File) {
	return wrapLocalNameChecker(&captLocalChecker{
		baseLocalNameChecker: baseLocalNameChecker{ctx: ctx},

		loudNames: map[string]bool{
			"IN":    true,
			"OUT":   true,
			"INOUT": true,

			// TODO: add common acronyms like HTTP and URL?
		},
	})
}

func (c *captLocalChecker) CheckLocalName(id *ast.Ident) {
	switch {
	case c.loudNames[id.Name]:
		c.warnLoud(id)
	case ast.IsExported(id.Name):
		c.warnCapitalized(id)
	}
}

func (c *captLocalChecker) warnCapitalized(id ast.Node) {
	c.ctx.Warn(id, "`%s' should not be capitalized", id)
}

func (c *captLocalChecker) warnLoud(id *ast.Ident) {
	c.ctx.Warn(id, "consider `%s' name instead of `%s'",
		strings.ToLower(id.Name), id)
}
