package lint

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/Quasilyte/astcmp"
)

// ParamNameChecker detects potential issues in function parameter names.
//
// Rationale: makes godoc output better.
type ParamNameChecker struct {
	ctx *Context

	loudNames map[string]bool

	warnings []Warning
}

// NewParamNameChecker returns initialized checker for Go functions param names.
func NewParamNameChecker(ctx *Context) *ParamNameChecker {
	return &ParamNameChecker{
		ctx: ctx,
		loudNames: map[string]bool{
			"IN":    true,
			"OUT":   true,
			"INOUT": true,
		},
	}
}

// Check runs parameter names checks for f.
//
// Features
//
// 1. Detects somewhat common "loud" identifiers, like IN or OUT.
//    Suggests to replace them with down-case versions (in, out).
//
// 2. If capitalized name is not recognized as "loud",
//    treat it as "redundantly exported".
//    Suggests to use non-capitalized identifier.
func (c *ParamNameChecker) Check(f *ast.File) []Warning {
	c.warnings = c.warnings[:0]
	for _, decl := range collectFuncDecls(f) {
		c.checkParamDuplication(decl)
		for _, param := range c.collectFuncParams(decl) {
			for _, id := range param.Names {
				switch {
				case c.loudNames[id.Name]:
					c.warnLoud(id)
				case ast.IsExported(id.Name):
					c.warnCapitalized(id)
				}
			}
		}
	}
	return c.warnings
}

func (c *ParamNameChecker) collectFuncParams(decl *ast.FuncDecl) []*ast.Field {
	var params []*ast.Field
	if decl.Recv != nil {
		recv := decl.Recv.List[0]
		params = append(params, recv)
	}
	params = append(params, decl.Type.Params.List...)
	if decl.Type.Results != nil {
		params = append(params, decl.Type.Results.List...)
	}
	return params
}

// TODO(fexolm) don't create multiple warnings on the same function
// TODO(fexolm) create warning in other function
func (c *ParamNameChecker) checkParamDuplication(decl *ast.FuncDecl) {
	params := decl.Type.Params.List
	if len(params) < 2 {
		return
	}
	for i, p := range params[1:] {
		if astcmp.EqualExpr(p.Type, params[i].Type) {
			var winfo string
			winfo += paramNamesStr(params[i].Names) + " "
			winfo += nodeString(c.ctx.FileSet, params[i].Type) + ", "

			winfo += paramNamesStr(p.Names) + " "
			winfo += nodeString(c.ctx.FileSet, p.Type)

			winfo += " could be replaced with "

			winfo += paramNamesStr(params[i].Names) + ", "
			winfo += paramNamesStr(p.Names) + " "
			winfo += nodeString(c.ctx.FileSet, p.Type)

			c.warnings = append(c.warnings, Warning{
				Kind: "Duplication",
				Node: decl,
				Text: fmt.Sprint(winfo),
			})
		}
	}
}

func (c *ParamNameChecker) warnCapitalized(id ast.Node) {
	c.warnings = append(c.warnings, Warning{
		Kind: "Capitalized",
		Node: id,
		Text: fmt.Sprintf("`%s' should not be capitalized", id),
	})
}

func (c *ParamNameChecker) warnLoud(id *ast.Ident) {
	c.warnings = append(c.warnings, Warning{
		Kind: "Loud",
		Node: id,
		Text: fmt.Sprintf("consider `%s' name instead of `%s'",
			strings.ToLower(id.Name), id),
	})
}
