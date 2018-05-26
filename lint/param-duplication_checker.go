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
}

// NewParamDuplicationChecker returns initialized ParamDuplicationChecker.
func newParamDuplicationChecker(ctx *Context) Checker {
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
func (c *ParamDuplicationChecker) Check(f *ast.File) {
	for _, decl := range collectFuncDecls(f) {
		c.checkParamDuplication(decl)
	}
}

// TODO(fexolm) don't create multiple warnings on the same function.
func (c *ParamDuplicationChecker) checkParamDuplication(decl *ast.FuncDecl) {
	params := decl.Type.Params.List
	if len(params) < 2 {
		return
	}
	equalRange := [][]*ast.Field{}
	equalRange = append(equalRange, []*ast.Field{params[0]})
	warn := false
	for i, p := range params[1:] {
		if astcmp.EqualExpr(p.Type, params[i].Type) {
			equalRange[len(equalRange)-1] =
				append(equalRange[len(equalRange)-1], p)
			warn = true
		} else {
			equalRange = append(equalRange, []*ast.Field{p})
		}
	}
	if warn {
		c.warn(equalRange, decl)
	}
}
func (c *ParamDuplicationChecker) warn(fields [][]*ast.Field, decl *ast.FuncDecl) {
	paramsBefore := []string{}
	paramsAfter := []string{}
	for _, flist := range fields {
		typeBefore := []string{}
		typeAfter := []string{}

		for _, f := range flist {
			typeBefore = append(typeBefore,
				c.paramNamesStr(f.Names)+" "+nodeString(c.ctx.FileSet, f.Type))
			typeAfter = append(typeAfter,
				c.paramNamesStr(f.Names))
		}

		paramsBefore = append(paramsBefore, strings.Join(typeBefore, ", "))
		paramsAfter = append(paramsAfter,
			strings.Join(typeAfter, ", ")+" "+nodeString(c.ctx.FileSet, flist[0].Type))
	}
	c.ctx.addWarning(Warning{
		Kind: "param-duplication/Duplication",
		Node: decl,
		Text: fmt.Sprintf("%s could be replaced with %s",
			strings.Join(paramsBefore, ", "),
			strings.Join(paramsAfter, ", ")),
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
	return strings.Join(names, ", ")
}
