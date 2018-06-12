package lint

import (
	"go/ast"

	"github.com/cristaloleg/astp"
	"github.com/go-toolsmith/astcopy"
	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	addChecker(parenthesisChecker{}, &ruleInfo{})
}

type parenthesisChecker struct {
	baseTypeExprChecker
}

func (c parenthesisChecker) New(ctx *context) func(*ast.File) {
	return wrapTypeExprChecker(&parenthesisChecker{
		baseTypeExprChecker: baseTypeExprChecker{ctx: ctx},
	})
}

func (c *parenthesisChecker) CheckTypeExpr(expr ast.Expr) {
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

func (c *parenthesisChecker) hasParens(x ast.Expr) bool {
	return nil != findNode(x, func(x ast.Node) bool {
		return astp.IsParenExpr(x)
	})
}

func (c *parenthesisChecker) unparenExpr(x ast.Expr) ast.Expr {
	// Replace every paren expr with expression it encloses.
	return astutil.Apply(x, nil, func(cur *astutil.Cursor) bool {
		if paren, ok := cur.Node().(*ast.ParenExpr); ok {
			cur.Replace(paren.X)
		}
		return true
	}).(ast.Expr)
}

func (c *parenthesisChecker) warn(cause, noParens ast.Expr) {
	c.ctx.Warn(cause, "could simplify %s to %s",
		nodeString(c.ctx.FileSet, cause),
		nodeString(c.ctx.FileSet, noParens))
}
