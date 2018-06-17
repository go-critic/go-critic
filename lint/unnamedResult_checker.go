package lint

//! For functions with multiple return values, detects unnamed results
//  that do not match `(T, error)` or `(T, bool)` pattern.
//
// Before:
// func f() (float64, float64)
//
// After:
// func f() (x, y float64)

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(&unnamedResultChecker{})
}

type unnamedResultChecker struct {
	baseFuncDeclChecker
}

func (c *unnamedResultChecker) New(ctx *context) func(*ast.File) {
	return wrapFuncDeclChecker(&unnamedResultChecker{
		baseFuncDeclChecker: baseFuncDeclChecker{ctx: ctx},
	})
}

func (c *unnamedResultChecker) CheckFuncDecl(decl *ast.FuncDecl) {
	results := decl.Type.Results
	switch {
	case results == nil || results.NumFields() < 2:
		return

	case results.NumFields() == 2:
		typ1, typ2 := c.getResultsTypes(results.List)
		if len(results.List[0].Names) == 2 ||
			(!c.isBoolOrError(typ1) && c.isBoolOrError(typ2)) {
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

func (c *unnamedResultChecker) getResultsTypes(fields []*ast.Field) (res1, res2 ast.Expr) {
	if len(fields) == 2 {
		return fields[0].Type, fields[1].Type
	}
	return fields[0].Type, fields[0].Type
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
