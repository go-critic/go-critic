package checkers

import (
	"go/ast"
	"strings"

	"github.com/go-lintpack/lintpack"
	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-toolsmith/astcast"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "filepathJoin"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects problems in filepath.Join() function calls"
	info.Before = `filepath.Join("dir/", filename)`
	info.After = `filepath.Join("dir", filename)`

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return astwalk.WalkerForExpr(&filepathJoinChecker{ctx: ctx})
	})
}

type filepathJoinChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext
}

func (c *filepathJoinChecker) VisitExpr(expr ast.Expr) {
	call := astcast.ToCallExpr(expr)
	if qualifiedName(call.Fun) != "filepath.Join" {
		return
	}

	for _, arg := range call.Args {
		arg, ok := arg.(*ast.BasicLit)
		if ok && c.hasSeparator(arg) {
			c.warnSeparator(arg)
		}
	}
}

func (c *filepathJoinChecker) hasSeparator(v *ast.BasicLit) bool {
	return strings.ContainsAny(v.Value, `/\`)
}

func (c *filepathJoinChecker) warnSeparator(sep ast.Expr) {
	c.ctx.Warn(sep, "%s contains a path separator", sep)
}
