package lint

//! For functions with multiple return values, detects unnamed results
//  that do not match `(T, error)` or `(T, bool)` pattern.
//
// @Before:
// func f() (float64, float64)
//
// @After:
// func f() (x, y float64)

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(&unnamedResultChecker{}, attrExperimental)
}

type unnamedResultChecker struct {
	checkerBase
}

func (c *unnamedResultChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	results := decl.Type.Results
	switch {
	case results == nil:
		return // Function has no results
	case len(results.List) > 0 && results.List[0].Names != nil:
		return // Skip named results
	}

	typeName := func(x ast.Expr) string { return c.typeName(c.ctx.typesInfo.TypeOf(x)) }
	isError := func(x ast.Expr) bool { return qualifiedName(x) == "error" }
	isBool := func(x ast.Expr) bool { return qualifiedName(x) == "bool" }

	// Main difference with case of len=2 is that we permit any
	// typ1 as long as second type is either error or bool.
	if results.NumFields() == 2 {
		typ1, typ2 := results.List[0].Type, results.List[1].Type
		name1, name2 := typeName(typ1), typeName(typ2)
		cond := (name1 != name2 && name2 != "") ||
			(!isError(typ1) && isError(typ2)) ||
			(!isBool(typ1) && isBool(typ2))
		if !cond {
			c.warn(decl)
		}
		return
	}

	seen := make(map[string]bool, len(results.List))
	for i := range results.List {
		typ := results.List[i].Type
		name := typeName(typ)
		isLast := i == len(results.List)-1

		cond := !seen[name] ||
			(isLast && (isError(typ) || isBool(typ)))
		if !cond {
			c.warn(decl)
			return
		}

		seen[name] = true
	}
}

func (c *unnamedResultChecker) typeName(typ types.Type) string {
	switch typ := typ.(type) {
	case *types.Array:
		return c.typeName(typ.Elem())
	case *types.Pointer:
		return c.typeName(typ.Elem())
	case *types.Slice:
		return c.typeName(typ.Elem())
	case *types.Named:
		return typ.Obj().Name()
	default:
		return ""
	}
}

func (c *unnamedResultChecker) warn(n ast.Node) {
	c.ctx.Warn(n, "consider giving a name to these results")
}
