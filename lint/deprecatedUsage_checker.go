package lint

//! Detects usage of deprecated methods and warns user about that.
//
// @Before:
// Foo() // has a comment that it's deprecated
//
// @After:
// Bar() // no comment about deprecation. good.

import (
	"go/ast"
	"strings"
)

func init() {
	addChecker(&deprecatedUsageChecker{}, attrExperimental)
}

type deprecatedUsageChecker struct {
	checkerBase
}

func (c *deprecatedUsageChecker) VisitExpr(x ast.Expr) {
	call, ok := x.(*ast.CallExpr)
	if !ok {
		return
	}

	id, ok := call.Fun.(*ast.Ident)
	if !ok || id.Obj == nil {
		return
	}

	decl, ok := id.Obj.Decl.(*ast.FuncDecl)
	if !ok {
		return
	}

	if text := c.deprecatedParagraph(decl.Doc); text != "" {
		c.warn(call, text)
	}
}

func (c *deprecatedUsageChecker) deprecatedParagraph(cs *ast.CommentGroup) string {
	if cs == nil {
		return ""
	}
	prefix := "// Deprecated: "
	for _, c := range cs.List {
		if strings.HasPrefix(c.Text, prefix) {
			return strings.TrimPrefix(c.Text, prefix)
		}
	}
	return ""
}

func (c *deprecatedUsageChecker) warn(cause *ast.CallExpr, suggestion string) {
	c.ctx.Warn(cause, "%s is deprected, see doc: %s", cause, suggestion)
}
