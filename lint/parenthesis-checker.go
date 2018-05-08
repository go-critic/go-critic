package lint

import (
	"fmt"
	"go/ast"
)

// ParenthesisChecker detects some cases where parenthesis are unnecessary
type ParenthesisChecker struct {
	ctx *Context

	warnings []Warning
}

// NewParenthesisChecker returns initialized checker for type expressions.
func NewParenthesisChecker(ctx *Context) *ParenthesisChecker {
	return &ParenthesisChecker{
		ctx: ctx,
	}
}

// Check runs parenthesis checks for f.
//
// Features
//
// Detects parenthesis statements which could be simplified
// and offsers the way how to do it.
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
	// TODO improve suggestions for complex cases like (func([](func())))

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
