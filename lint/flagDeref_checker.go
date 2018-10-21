package lint

import (
	"go/ast"
)

func init() {
	addChecker(&flagDerefChecker{}, attrSyntaxOnly)
}

type flagDerefChecker struct {
	checkerBase

	flagPtrFuncs map[string]bool
}

func (c *flagDerefChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects immediate dereferencing of `flag` package pointers"
	d.Before = `b := *flag.Bool("b", false, "b docs")`
	d.After = `
var b bool
flag.BoolVar(&b, "b", false, "b docs")`
	d.Note = `
Dereferencing returned pointers will lead to hard to find errors
where flag values are not updated after flag.Parse().`
}

func (c *flagDerefChecker) Init() {
	c.flagPtrFuncs = map[string]bool{
		"flag.Bool":     true,
		"flag.Duration": true,
		"flag.Float64":  true,
		"flag.Int":      true,
		"flag.Int64":    true,
		"flag.String":   true,
		"flag.Uint":     true,
		"flag.Uint64":   true,
	}
}

func (c *flagDerefChecker) VisitExpr(expr ast.Expr) {
	if expr, ok := expr.(*ast.StarExpr); ok {
		call, ok := expr.X.(*ast.CallExpr)
		if !ok {
			return
		}
		called := qualifiedName(call.Fun)
		if c.flagPtrFuncs[called] {
			c.warn(expr, called+"Var")
		}
	}
}

func (c *flagDerefChecker) warn(x ast.Node, suggestion string) {
	c.ctx.Warn(x, "immediate deref in %s is most likely an error; consider using %s",
		x, suggestion)
}
