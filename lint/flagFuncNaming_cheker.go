package lint

import (
	"go/ast"
	"go/types"
	"strings"
)

func init() {
	addChecker(flagFuncNamingChecker{}, &ruleInfo{
		Experimental: true,
	})
}

type flagFuncNamingChecker struct {
	baseFuncDeclChecker
}

func (c flagFuncNamingChecker) New(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&flagFuncNamingChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

func (c *flagFuncNamingChecker) CheckFuncDecl(decl *ast.FuncDecl) {
	params := decl.Type.Params
	results := decl.Type.Results

	if params.NumFields() > 0 ||
		results.NumFields() != 1 ||
		!c.isBoolType(results.List[0].Type) ||
		c.hasProperPrefix(decl.Name.Name) {
		return
	}
	c.warn(decl)
}

func (c *flagFuncNamingChecker) warn(fn *ast.FuncDecl) {
	c.ctx.Warn(fn, "consider to add Is/Has/Contains prefix to function name")
}

func (c *flagFuncNamingChecker) isBoolType(expr ast.Expr) bool {
	switch typ := c.ctx.TypesInfo.TypeOf(expr).(type) {
	case *types.Basic:
		return typ.Kind() == types.Bool
	default:
		return false
	}
}

func (c *flagFuncNamingChecker) hasProperPrefix(name string) bool {
	name = strings.ToLower(name)
	return strings.HasPrefix(name, "is") ||
		strings.HasPrefix(name, "has") ||
		strings.HasPrefix(name, "contains")
}
