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

	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			if decl.Type.Results == nil {
				continue
			}
			for _, res := range decl.Type.Results.List {
				c.validateExpr(res.Type)
			}
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				spec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				c.validateExpr(spec.Type)
			}
		}
	}
	return c.warnings
}

func (c *ParenthesisChecker) validateExpr(n ast.Node) {
	// TODO improve suggestions for complex cases like (func([](func())))
	// TODO improve linter output to write full type, not just place
	// where it could be simplified

	ast.Inspect(n, func(n ast.Node) bool {
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
