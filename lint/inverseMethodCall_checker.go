package lint

import (
	"go/ast"
	"go/types"
	"strings"
)

func init() {
	addChecker(&inverseMethodCallChecker{}, attrExperimental)
}

type inverseMethodCallChecker struct {
	checkerBase
}

func (c *inverseMethodCallChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects inverse method execution"
	d.Before = `type foo struct{}
f := foo{}
foo.bar(f)`
	d.After = `type foo struct{}
f := foo{}
f.bar()`
}

func (c *inverseMethodCallChecker) VisitExpr(x ast.Expr) {
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

	obj := c.ctx.typesInfo.Uses[id]
	if c.isObjectStruct(obj) {
		c.warn(call, "XXX")
	}

}

func (c *inverseMethodCallChecker) isObjectStruct(t types.Object) bool {
	s := t.String()
	return strings.HasPrefix(s, "type ") && strings.Contains(s, " struct{")
}

func (c *inverseMethodCallChecker) warn(cause *ast.CallExpr, suggestion string) {
	c.ctx.Warn(cause, "consider to change to `%s`", suggestion)
}
