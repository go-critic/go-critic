package lint

import (
	"go/ast"
	"strings"
)

// paramNameCheck detects potential issues in function parameter names.
//
// Rationale: better godoc; code readability.
//
// Detects somewhat common "loud" identifiers, like IN or OUT.
// Suggests to replace them with down-case versions (in, out).
//
// If capitalized name is not recognized as "loud",
// treat it as "redundantly exported".
// Suggests to use non-capitalized identifier.
func paramNameCheck(ctx *context) func(*ast.File) {
	return wrapParamListChecker(&paramNameChecker{
		baseParamListChecker: baseParamListChecker{ctx: ctx},

		loudNames: map[string]bool{
			"IN":    true,
			"OUT":   true,
			"INOUT": true,

			// TODO: add common acronyms like HTTP and URL?
		},
	})
}

type paramNameChecker struct {
	baseParamListChecker

	loudNames map[string]bool
}

func (c *paramNameChecker) CheckParamList(params []*ast.Field) {
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

func (c *paramNameChecker) warnCapitalized(id ast.Node) {
	c.ctx.Warn(id, "`%s' should not be capitalized", id)
}

func (c *paramNameChecker) warnLoud(id *ast.Ident) {
	c.ctx.Warn(id, "consider `%s' name instead of `%s'",
		strings.ToLower(id.Name), id)
}
