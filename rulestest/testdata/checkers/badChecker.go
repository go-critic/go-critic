package checkers

import (
	"go/ast"
)

type BadChecker struct {
	ctx *contextStub
}

func (c *BadChecker) VisitExpr(e ast.Expr) {
	_ = c.ctx.TypesInfo.TypeOf(e) // want `\Quse ctx.TypeOf(e) instead, it's nil-safe`
}
