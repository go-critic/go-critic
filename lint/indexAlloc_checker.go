package lint

import (
	"go/ast"

	"github.com/go-critic/go-critic/lint/internal/lintutil"
)

func init() {
	addChecker(&indexAllocChecker{}, attrExperimental, attrPerformance)
}

type indexAllocChecker struct {
	checkerBase
}

func (c *indexAllocChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects strings.Index calls that may cause unwanted allocs"
	d.Before = `strings.Index(string(x), y)`
	d.After = `bytes.Index(x, []byte(y))`
	d.Note = `See Go issue for details: https://github.com/golang/go/issues/25864`
}

func (c *indexAllocChecker) VisitExpr(e ast.Expr) {
	call := lintutil.AsCallExpr(e)
	if qualifiedName(call.Fun) != "strings.Index" {
		return
	}
	stringConv := lintutil.AsCallExpr(call.Args[0])
	if qualifiedName(stringConv.Fun) != "string" {
		return
	}
	x := stringConv.Args[0]
	y := call.Args[1]
	if isSafeExpr(c.ctx.typesInfo, x) && isSafeExpr(c.ctx.typesInfo, y) {
		c.warn(e, x, y)
	}
}

func (c *indexAllocChecker) warn(cause ast.Node, x, y ast.Expr) {
	c.ctx.Warn(cause, "consider replacing %s with bytes.Index(%s, []byte(%s))",
		cause, x, y)
}
