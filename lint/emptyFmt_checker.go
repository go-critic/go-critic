package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astcast"
)

func init() {
	addChecker(&emptyFmtChecker{}, attrExperimental)
}

type emptyFmtChecker struct {
	checkerBase
}

func (c *emptyFmtChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects usages of formatting functions without formatting arguments"
	d.Before = `
fmt.Sprintf("whatever")
fmt.Errorf("wherever")`
	d.After = `
fmt.Sprint("whatever")
errors.New("wherever")`
}

func (c *emptyFmtChecker) VisitExpr(expr ast.Expr) {
	call := astcast.ToCallExpr(expr)
	name := qualifiedName(call.Fun)

	switch len(call.Args) {
	case 1:
		switch name {
		case "fmt.Sprintf":
			c.warn(call, "fmt.Sprint")
		case "log.Printf":
			c.warn(call, "log.Print")
		case "log.Panicf":
			c.warn(call, "log.Panic")
		case "log.Fatalf":
			c.warn(call, "log.Fatal")
		case "fmt.Errorf":
			c.warn(call, "errors.New")
		}
	case 2:
		if name == "fmt.Fprintf" {
			c.warn(call, "fmt.Fprint")
		}
	}
}

func (c *emptyFmtChecker) warn(cause *ast.CallExpr, suggestion string) {
	c.ctx.Warn(cause, "consider to change function to %s", suggestion)
}
