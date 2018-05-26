package lint

import (
	"go/ast"
	"strings"

	"github.com/Quasilyte/astcmp"
)

// paramDuplicationCheck detects if function parameters could be combined by type
// and suggest the way to do it.
//
// Rationale: better godoc; code readability.
//
// Example: func f(a int, b int) could be simplified as func f(a, b int)
func paramDuplicationCheck(ctx *context) func(*ast.File) {
	return wrapParamListChecker(&paramDuplicationChecker{
		baseParamListChecker: baseParamListChecker{ctx: ctx},
	})
}

type paramDuplicationChecker struct {
	baseParamListChecker

	cause *ast.FuncDecl
}

func (c *paramDuplicationChecker) PerFuncInit(fn *ast.FuncDecl) bool {
	c.cause = fn
	return true
}

func (c *paramDuplicationChecker) CheckParamList(params []*ast.Field) {
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
		c.warn(equalRange)
	}
}

func (c *paramDuplicationChecker) warn(fields [][]*ast.Field) {
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
	c.ctx.Warn(c.cause, "%s could be replaced with %s",
		strings.Join(paramsBefore, ", "),
		strings.Join(paramsAfter, ", "))
}

func (c *paramDuplicationChecker) paramNamesStr(idents []*ast.Ident) string {
	if idents == nil {
		return "_"
	}
	names := []string{}
	for _, id := range idents {
		names = append(names, id.Name)
	}
	return strings.Join(names, ", ")
}
