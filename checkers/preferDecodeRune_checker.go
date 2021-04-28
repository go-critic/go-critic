package checkers

import (
	"go/ast"
	"go/types"

	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/typep"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/framework/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "preferDecodeRune"
	info.Tags = []string{"performance"}
	info.Summary = "Detects expressions like []rune(s)[0] that may cause unwanted rune slice allocation"
	info.Before = `r := []rune(s)[0]`
	info.After = `r, _ := utf8.DecodeRuneInString(s)`
	info.Note = `See Go issue for details: https://github.com/golang/go/issues/45260`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&preferDecodeRuneChecker{ctx: ctx}), nil
	})
}

type preferDecodeRuneChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *preferDecodeRuneChecker) VisitExpr(e ast.Expr) {
	indexExpr := astcast.ToIndexExpr(e)
	indexed := indexExpr.X

	// Check that indexed is a []rune.
	slice, ok := c.ctx.TypeOf(indexed).(*types.Slice)
	if !ok || !typep.HasInt32Kind(slice.Elem()) {
		return
	}

	cast := astcast.ToCallExpr(indexed)
	if len(cast.Args) != 1 || !typep.HasStringProp(c.ctx.TypeOf(cast.Args[0])) {
		return
	}
	arg := cast.Args[0]

	if astcast.ToBasicLit(indexExpr.Index).Value == "0" {
		c.warnUseDecodeRune(indexExpr, arg, "DecodeRuneInString")
	}
}

func (c *preferDecodeRuneChecker) warnUseDecodeRune(cause *ast.IndexExpr, arg ast.Expr, method string) {
	c.ctx.Warn(cause, "consider replacing %s with utf8.%s(%s)", cause, method, arg)
}
