package lint

import (
	"fmt"
	"go/ast"
	"strings"

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
func (c *ParamDuplicationChecker) checkParamDuplication(decl *ast.FuncDecl) {
	params := decl.Type.Params.List
	if len(params) < 2 {
		return
	}
	for i, p := range params[1:] {
		if astcmp.EqualExpr(p.Type, params[i].Type) {
			c.warn(params[i], p, decl)
		}
	}
}

func (c *ParamDuplicationChecker) warn(a, b *ast.Field, decl *ast.FuncDecl) {
	var winfo string
	winfo += c.paramNamesStr(a.Names) + " "
	winfo += nodeString(c.ctx.FileSet, a.Type) + ", "

	winfo += c.paramNamesStr(b.Names) + " "
	winfo += nodeString(c.ctx.FileSet, b.Type)

	winfo += " could be replaced with "

	winfo += c.paramNamesStr(a.Names) + ", "
	winfo += c.paramNamesStr(b.Names) + " "
	winfo += nodeString(c.ctx.FileSet, b.Type)

	c.warnings = append(c.warnings, Warning{
		Kind: "Duplication",
		Node: decl,
		Text: fmt.Sprint(winfo),
	})

}

func (c *ParamDuplicationChecker) paramNamesStr(idents []*ast.Ident) string {
	if idents == nil {
		return "_"
	}
	names := []string{}
	for _, id := range idents {
		names = append(names, id.Name)
	}
	return strings.Join(names, " ,")
}
