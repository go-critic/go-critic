package lint

import (
	"go/ast"
	"go/types"

	"github.com/go-toolsmith/astcopy"
)

func init() {
	addChecker(&methodExprCallChecker{}, attrExperimental, attrVeryOpinionated)
}

type methodExprCallChecker struct {
	checkerBase
}

func (c *methodExprCallChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects method expression call that can be replaced with a method call"
	d.Before = `f := foo{}
foo.bar(f)`
	d.After = `f := foo{}
f.bar()`
}

func (c *methodExprCallChecker) VisitExpr(x ast.Expr) {
	call, ok := x.(*ast.CallExpr)
	if !ok {
		return
	}

	s, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	id, ok := s.X.(*ast.Ident)
	if !ok {
		return
	}

	obj := c.ctx.typesInfo.ObjectOf(id)
	if _, ok := obj.(*types.TypeName); ok {
		c.warn(call, s)
	}
}

func (c *methodExprCallChecker) warn(cause *ast.CallExpr, s *ast.SelectorExpr) {
	selector := astcopy.SelectorExpr(s)
	selector.X = cause.Args[0]

	c.ctx.Warn(cause, "consider to change `%s` to `%s`", cause.Fun, selector)
}
