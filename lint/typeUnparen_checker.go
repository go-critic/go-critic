package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/astp"
	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	addChecker(typeUnparenChecker{}, &ruleInfo{
		SyntaxOnly: true,
	})
}

type typeUnparenChecker struct {
	baseTypeExprChecker
}

func (c typeUnparenChecker) New(ctx *context) func(*ast.File) {
	return wrapTypeExprChecker(&typeUnparenChecker{
		baseTypeExprChecker: baseTypeExprChecker{ctx: ctx},
	})
}

func (c *typeUnparenChecker) CheckTypeExpr(expr ast.Expr) {
	// Arrays require extra care: we don't want to unparen
	// length expression as they are not type expressions.
	if arr, ok := expr.(*ast.ArrayType); ok {
		if !c.hasParens(arr.Elt) {
			return
		}
		noParens := astcopy.Expr(arr).(*ast.ArrayType)
		noParens.Elt = c.unparenExpr(noParens.Elt)
		c.warn(expr, noParens)
		return
	}
	if !c.hasParens(expr) {
		return
	}
	c.warn(expr, c.unparenExpr(astcopy.Expr(expr)))
}

func (c *typeUnparenChecker) hasParens(x ast.Expr) bool {
	return findNode(x, astp.IsParenExpr) != nil
}

func (c *typeUnparenChecker) unparenExpr(x ast.Expr) ast.Expr {
	// Replace every paren expr with expression it encloses.
	return astutil.Apply(x, nil, func(cur *astutil.Cursor) bool {
		if paren, ok := cur.Node().(*ast.ParenExpr); ok {
			cur.Replace(paren.X)
		}
		return true
	}).(ast.Expr)
}

func (c *typeUnparenChecker) warn(cause, noParens ast.Expr) {
	c.ctx.Warn(cause, "could simplify %s to %s", cause, noParens)
}
