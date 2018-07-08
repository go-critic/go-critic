package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/astp"
	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	addChecker(&typeUnparenChecker{}, attrSyntaxOnly)
}

type typeUnparenChecker struct {
	checkerBase

	cause ast.Node // Last warning cause
}

func (c *typeUnparenChecker) InitDocs(d *Documentation) {
	d.Summary = "Detects unneded parenthesis inside type expressions and suggests to remove them"
	d.Before = "type foo [](func([](func())))"
	d.After = "type foo []func([]func())"
}

func (c *typeUnparenChecker) EnterChilds(x ast.Node) bool { return c.cause != x }

func (c *typeUnparenChecker) checkTypeExpr(x ast.Expr) {
	switch x := x.(type) {
	case *ast.ArrayType:
		// Arrays require extra care: we don't want to unparen
		// length expression as they are not type expressions.
		if !c.hasParens(x.Elt) {
			return
		}
		noParens := astcopy.ArrayType(x)
		noParens.Elt = c.unparenExpr(noParens.Elt)
		c.warn(x, noParens)
	case *ast.StructType, *ast.InterfaceType:
		// Only nested fields are to be reported.
	default:
		if !c.hasParens(x) {
			return
		}
		// fmt.Printf("=> %#v\n", x)
		c.warn(x, c.unparenExpr(astcopy.Expr(x)))
	}
}

func (c *typeUnparenChecker) VisitTypeExpr(x ast.Expr) {
	switch x := x.(type) {
	case *ast.ParenExpr:
		switch x.X.(type) {
		case *ast.StructType:
			c.ctx.Warn(x, "could simplify (struct{...}) to struct{...}")
		case *ast.InterfaceType:
			c.ctx.Warn(x, "could simplify (interface{...}) to interface{...}")
		default:
			c.warn(x, c.unparenExpr(astcopy.Expr(x)))
		}
	default:
		c.checkTypeExpr(x)
	}
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
	c.cause = cause
	c.ctx.Warn(cause, "could simplify %s to %s", c.cause, noParens)
}
