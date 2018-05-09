package lint

import (
	"fmt"
	"go/ast"

	"github.com/Quasilyte/astcmp"
)

// ParamDuplicationChecker detects functions where parameters declaration
// could be simplified by combine arguments with same type.
//
// Example: func f(a int, b int) could be simplified as func f(a, b int)
type ParamDuplicationChecker struct {
	ctx *Context

	warnings []Warning
}

// NewParamDuplicationChecker returns initialized ParamDuplicationChecker.
func NewParamDuplicationChecker(ctx *Context) *ParamDuplicationChecker {
	return &ParamDuplicationChecker{
		ctx: ctx,
	}
}

// Check runs parameter types duplication checks for f.
//
// Features
//
// Detects if function parameters could be combined by type
// and suggest the way to do it.
func (c *ParamDuplicationChecker) Check(f *ast.File) []Warning {
	c.warnings = c.warnings[:0]
	for _, decl := range collectFuncDecls(f) {
		c.checkParamDuplication(decl)
	}
	return c.warnings
}

// TODO(fexolm) don't create multiple warnings on the same function.
// TODO(fexolm) create warning in other function.
func (c *ParamDuplicationChecker) checkParamDuplication(decl *ast.FuncDecl) {
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
