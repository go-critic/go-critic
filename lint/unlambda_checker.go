package lint

import (
	"go/ast"
	"go/types"

	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/astequal"
)

func init() {
	addChecker(&unlambdaChecker{}, attrExperimental)
}

type unlambdaChecker struct {
	checkerBase
}

func (c *unlambdaChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects function literals that can be simplified"
	d.Before = "func(x int) int { return fn(x) }"
	d.After = "fn"
}

func (c *unlambdaChecker) VisitExpr(x ast.Expr) {
	fn, ok := x.(*ast.FuncLit)
	if !ok || len(fn.Body.List) != 1 {
		return
	}

	ret, ok := fn.Body.List[0].(*ast.ReturnStmt)
	if !ok || len(ret.Results) != 1 {
		return
	}

	result := astcast.ToCallExpr(ret.Results[0])
	callable := qualifiedName(result.Fun)
	if callable == "" {
		return // Skip tricky cases; only handle simple calls
	}
	fnType := c.ctx.typesInfo.TypeOf(fn)
	resultType := c.ctx.typesInfo.TypeOf(result.Fun)
	if !types.Identical(fnType, resultType) {
		return
	}
	// Now check that all arguments match the parameters.
	n := 0
	for _, params := range fn.Type.Params.List {
		for _, id := range params.Names {
			if !astequal.Expr(id, result.Args[n]) {
				return
			}
			n++
		}
	}

	c.warn(fn, callable)
}

func (c *unlambdaChecker) warn(cause ast.Node, suggestion string) {
	c.ctx.Warn(cause, "replace `%s` with `%s`", cause, suggestion)
}
