package lint

//! Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`.
//
// @Before:
// re, _ := regexp.Compile(`const pattern`)
//
// @After:
// re := regexp.MustCompile(`const pattern`)

import (
	"go/ast"
	"strings"

	"github.com/go-toolsmith/astp"
	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	addChecker(&regexpMustChecker{}, attrExperimental)
}

type regexpMustChecker struct {
	checkerBase
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
