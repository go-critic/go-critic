package lint

import (
	"go/ast"
	"strings"

	"github.com/go-toolsmith/astp"
	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	addChecker(&regexpMustChecker{})
}

type regexpMustChecker struct {
	checkerBase
}

func (c *regexpMustChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`"
	d.Before = `re, _ := regexp.Compile("const pattern")`
	d.After = `re := regexp.MustCompile("const pattern")`
}

func (c *regexpMustChecker) VisitExpr(x ast.Expr) {
	if x, ok := x.(*ast.CallExpr); ok {
		switch name := qualifiedName(x.Fun); name {
		case "regexp.Compile", "regexp.CompilePOSIX":
			// Only check for trivial string args, permit parenthesis.
			if !astp.IsBasicLit(astutil.Unparen(x.Args[0])) {
				return
			}
			c.warn(x, strings.Replace(name, "Compile", "MustCompile", 1))
		}
	}
}

func (c *regexpMustChecker) warn(cause *ast.CallExpr, suggestion string) {
	c.ctx.Warn(cause, "for const patterns like %s, use %s",
		cause.Args[0], suggestion)
}
