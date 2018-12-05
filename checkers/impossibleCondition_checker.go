package checkers

import (
	"github.com/go-lintpack/lintpack"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "impossibleCondition"
	info.Tags = []string{"style", "experimental"}
	info.Summary = "Detects bool expressions that can be simplified"
	info.Before = `x < y && `
	info.After = `
a := elapsed < expectElapsedMin
b := (x) == (y)`

	// collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
	// 	return astwalk.WalkerForExpr(&boolExprSimplifyChecker{ctx: ctx})
	// })
}
