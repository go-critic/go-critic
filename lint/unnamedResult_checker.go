package lint

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(unnamedResultChecker{}, &ruleInfo{})
}

type unnamedResultChecker struct {
	baseFuncDeclChecker
}

func (c unnamedResultChecker) New(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&unnamedResultChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

func (c *unnamedResultChecker) CheckFuncDecl(decl *ast.FuncDecl) {
	results := decl.Type.Results
	switch {
	case results == nil || c.resultsNum(results.List) < 2:
		return

	case c.resultsNum(results.List) == 2:
		if !c.isBoolOrError(results.List[0].Type) &&
			c.isBoolOrError(results.List[1].Type) {
			// no need to name results for (T, error) or (T, bool)
		} else {
			c.warn(decl)
		}

	default:
		for _, res := range results.List {
			if len(res.Names) == 0 {
				c.warn(decl)
				break
			}
		}
	}
}

func (c *unnamedResultChecker) warn(n ast.Node) {
	c.ctx.Warn(n, "consider to give name to results")
}

// resultsNum will return number of results.
// If they're unnamed than we meed to return number of fields.
// Else we need to sum all amount of names in each field.
func (c *unnamedResultChecker) resultsNum(fields []*ast.Field) int {
	res := 0
	for _, f := range fields {
		if f != nil {
			res += len(f.Names)
		}
	}
	if len(fields) > res {
		return len(fields)
	}
	return res
}

func (c *unnamedResultChecker) isBoolOrError(expr ast.Expr) bool {
	switch typ := c.ctx.TypesInfo.TypeOf(expr).(type) {
	case *types.Named:
		return typ.Obj().Name() == "error"

	case *types.Basic:
		return typ.Kind() == types.Bool

	default:
		return false
	}
}
