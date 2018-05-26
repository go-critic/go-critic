package lint

import (
	"go/ast"
)

// parenthesisCheck detects unneded parenthesis inside type expressions
// and suggests to remove them.
//
// Rationale: code readability.
func parenthesisCheck(ctx *context) func(*ast.File) {
	return wrapTypeExprChecker(&parenthesisChecker{
		baseTypeExprChecker: baseTypeExprChecker{ctx: ctx},
	})
}

type parenthesisChecker struct {
	baseTypeExprChecker
}

func (c *parenthesisChecker) CheckTypeExpr(expr ast.Expr) {
	// TODO: improve suggestions for complex cases like (func([](func()))).
	// TODO: print outermost cause instead of innermost.
	ast.Inspect(expr, func(n ast.Node) bool {
		expr, ok := n.(*ast.ParenExpr)
		if !ok {
			return true
		}
		c.ctx.Warn(expr, "could simplify %s to %s",
			nodeString(c.ctx.FileSet, expr),
			nodeString(c.ctx.FileSet, expr.X))
		return false
	})
}
