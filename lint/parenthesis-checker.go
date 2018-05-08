package lint

import (
	"fmt"
	"go/ast"
)

// ParenthesisChecker ...
type ParenthesisChecker struct {
	ctx *Context

	warnings []Warning
}

func NewParenthesisChecker(ctx *Context) *ParenthesisChecker {
	return &ParenthesisChecker{
		ctx: ctx,
	}
}

func (c *ParenthesisChecker) Check(f *ast.File) []Warning {
	c.warnings = c.warnings[:0]
	for _, decl := range collectFuncDecls(f) {
		if decl.Type.Results == nil {
			continue
		}
		for _, res := range decl.Type.Results.List {
			c.validateResultDecl(res)
		}
	}
	return c.warnings
}

func (c *ParenthesisChecker) validateResultDecl(f *ast.Field) {
	ast.Inspect(f.Type, func(n ast.Node) bool {
		expr, ok := n.(*ast.ParenExpr)
		if !ok {
			return true
		}
		c.warnings = append(c.warnings, Warning{
			Node: expr,
			Text: fmt.Sprintf("could simplify %s to %s", nodeString(c.ctx.FileSet, expr), nodeString(c.ctx.FileSet, expr.X)),
		})
		return false
	})
}
