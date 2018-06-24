package lint

//! Detects usages of formatting functions without formatting arguments.
//
// @Before:
// fmt.Sprintf("whatever")
// fmt.Errorf("wherever")
//
// @After:
// fmt.Sprint("whatever")
// errors.New("wherever")

import (
	"go/ast"
	"strings"
)

func init() {
	addChecker(&emptyFmtChecker{}, attrExperimental)
}

type emptyFmtChecker struct {
	checkerBase
}

func (c *emptyFmtChecker) VisitExpr(expr ast.Expr) {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return
	}

	name := qualifiedName(call.Fun)

	switch len(call.Args) {
	case 1:
		switch name {
		case "fmt.Sprintf":
			c.warn(call, name, "fmt.Sprint")
		case "log.Printf":
			c.warn(call, name, "log.Print")
		case "log.Panicf":
			c.warn(call, name, "log.Panic")
		case "log.Fatalf":
			c.warn(call, name, "log.Fatal")
		case "fmt.Errorf":
			c.warn(call, name, "errors.New")
		}
	case 2:
		if name == "fmt.Fprintf" {
			c.warn(call, name, "fmt.Fprint")
		}
	}
}

func (c *emptyFmtChecker) warn(cause *ast.CallExpr, original, suggestion string) {
	c.ctx.Warn(cause, "consider to change function to %s", strings.Replace(original, original, suggestion, 1))
}
