package lint

//! Detects bool expressions that can be simplified.
//
// @Before:
// a := !(elapsed >= expectElapsedMin)
// b := !(x) == !(y)
//
// @After:
// a := elapsed < expectElapsedMin
// b := (x) == (y)

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/astequal"
	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	addChecker(&boolOptChecker{}, attrExperimental)
}

type boolOptChecker struct {
	checkerBase

	cause ast.Node // Last warning cause

	// nil sentinels are used as a replacements for
	// bare nil to avoid a need to perform nil checks
	// when doing type-assertion that may return nil.
	nilUnaryExpr  *ast.UnaryExpr
	nilBinaryExpr *ast.BinaryExpr
}

func (c *boolOptChecker) Init() {
	c.nilUnaryExpr = &ast.UnaryExpr{}
	c.nilBinaryExpr = &ast.BinaryExpr{}
}

func (c *boolOptChecker) EnterChilds(x ast.Node) bool { return c.cause != x }

func (c *boolOptChecker) VisitExpr(x ast.Expr) {
	// TODO: avoid eager copy?
	// Can't be stable until wasted copying is fixed.
	opt := c.optimizeBool(astcopy.Expr(x))
	if !astequal.Expr(x, opt) {
		c.warn(x, opt)
	}
}

func (c *boolOptChecker) optimizeBool(x ast.Expr) ast.Expr {
	return astutil.Apply(x, nil, func(cur *astutil.Cursor) bool {
		return c.doubleNegation(cur) ||
			c.negatedEquals(cur) ||
			c.invertComparison(cur) ||
			true
	}).(ast.Expr)
}

func (c *boolOptChecker) doubleNegation(cur *astutil.Cursor) bool {
	neg1 := c.unaryNot(cur.Node())
	neg2 := c.unaryNot(astutil.Unparen(neg1.X))
	if neg1 != c.nilUnaryExpr && neg2 != c.nilUnaryExpr {
		cur.Replace(astutil.Unparen(neg2.X))
		return true
	}
	return false
}

func (c *boolOptChecker) negatedEquals(cur *astutil.Cursor) bool {
	x, ok := cur.Node().(*ast.BinaryExpr)
	if !ok || x.Op != token.EQL {
		return false
	}
	neg1 := c.unaryNot(x.X)
	neg2 := c.unaryNot(x.Y)
	if neg1 != c.nilUnaryExpr && neg2 != c.nilUnaryExpr {
		x.X = neg1.X
		x.Y = neg2.X
		return true
	}
	return false
}

func (c *boolOptChecker) invertComparison(cur *astutil.Cursor) bool {
	neg := c.unaryNot(cur.Node())
	cmp := c.binaryExpr(astutil.Unparen(neg.X))
	if neg == c.nilUnaryExpr || cmp == c.nilBinaryExpr {
		return false
	}

	// Replace operator to its negated form.
	switch cmp.Op {
	case token.EQL:
		cmp.Op = token.NEQ
	case token.NEQ:
		cmp.Op = token.EQL
	case token.LSS:
		cmp.Op = token.GEQ
	case token.GTR:
		cmp.Op = token.LEQ
	case token.LEQ:
		cmp.Op = token.GTR
	case token.GEQ:
		cmp.Op = token.LSS

	default:
		return false
	}
	cur.Replace(cmp)
	return true
}

// binaryExpr coerces x into binary expr if possible,
// otherwise returns c.nilBinaryExpr.
func (c *boolOptChecker) binaryExpr(x ast.Node) *ast.BinaryExpr {
	binexp, ok := x.(*ast.BinaryExpr)
	if !ok {
		return c.nilBinaryExpr
	}
	return binexp
}

// unaryNot coerces x into unary not if possible,
// otherwise returns c.nilUnaryExpr.
func (c *boolOptChecker) unaryNot(x ast.Node) *ast.UnaryExpr {
	neg, ok := x.(*ast.UnaryExpr)
	if !ok || neg.Op != token.NOT {
		return c.nilUnaryExpr
	}
	return neg
}

func (c *boolOptChecker) warn(cause, suggestion ast.Expr) {
	c.cause = cause
	c.ctx.Warn(cause, "can simplify `%s` to `%s`", cause, suggestion)
}
