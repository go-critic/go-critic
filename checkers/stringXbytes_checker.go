package checkers

import (
	"go/ast"
	"strings"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
	"github.com/go-toolsmith/astfmt"
	"github.com/go-toolsmith/typep"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "stringXbytes"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects redundant conversions between string and []byte"
	info.Before = `copy(b, []byte(s))`
	info.After = `copy(b, s)`

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return astwalk.WalkerForExpr(&stringXbytes{ctx: ctx})
	})
}

type stringXbytes struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext
}

func (c *stringXbytes) VisitExpr(expr ast.Expr) {
	if x, ok := expr.(*ast.CallExpr); ok {
		switch name := qualifiedName(x.Fun); name {
		case "copy":
			src := x.Args[1]

			if byteCast, ok := src.(*ast.CallExpr); ok &&
				typep.IsTypeExpr(c.ctx.TypesInfo, byteCast.Fun) &&
				typep.HasStringProp(c.ctx.TypesInfo.TypeOf(byteCast.Args[0])) {

				c.warn(byteCast, strings.TrimSuffix(strings.TrimPrefix(astfmt.Sprint(byteCast), "[]byte("), ")"))
			}
		}
	}
}

func (c *stringXbytes) warn(cause *ast.CallExpr, suggestion string) {
	c.ctx.Warn(cause, "can simplify `%s` to `%s`",
		cause, suggestion)
}
