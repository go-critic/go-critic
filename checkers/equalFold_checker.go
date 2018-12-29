package checkers

import (
	"go/ast"
	"go/token"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
	"github.com/go-toolsmith/astcast"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "equalFold"
	info.Tags = []string{"performance", "experimental"}
	info.Summary = "Detects unoptimal string equal"
	info.Before = `strings.ToLower(x) == y`
	info.After = `strings.EqualFold(x, y)`

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return astwalk.WalkerForExpr(&equalFoldChecker{ctx: ctx})
	})
}

type equalFoldChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext
}

func (c *equalFoldChecker) VisitExpr(e ast.Expr) {
	expr := astcast.ToBinaryExpr(e)
	if expr.Op != token.EQL && expr.Op != token.NEQ {
		return
	}

	callX := astcast.ToCallExpr(expr.X)
	callY := astcast.ToCallExpr(expr.Y)

	if qualifiedName(callX.Fun) != "strings.ToLower" &&
		qualifiedName(callX.Fun) != "strings.ToUpper" &&
		qualifiedName(callY.Fun) != "strings.ToLower" &&
		qualifiedName(callY.Fun) != "strings.ToUpper" {
		return
	}

	x := expr.X
	y := expr.Y

	if len(callX.Args) != 0 {
		x = callX.Args[0]
	}
	if len(callY.Args) != 0 {
		y = callY.Args[0]
	}
	c.warn(e, x, y)
}

func (c *equalFoldChecker) warn(cause ast.Node, x, y ast.Expr) {
	c.ctx.Warn(cause, "consider replacing with strings.EqualFold(%s, %s)", x, y)
}
