//go:generate go run ./goosChecker_gen/gen.go -goroot $GOROOT -o goosChecker_syslist_gen.go

package checkers

import (
	"go/ast"
	"go/constant"
	"go/token"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "goosChecker"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects comparisons of runtime.GOOS against unknown values"
	info.Before = `if runtime.GOOS == "foobar" {}`
	info.After = `if runtime.GOOS == "linux" {}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForExpr(&goosChecker{ctx: ctx}), nil
	})
}

type goosChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *goosChecker) VisitExpr(expr ast.Expr) {
	binExpr, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return
	}
	if binExpr.Op != token.EQL && binExpr.Op != token.NEQ {
		return
	}

	// Check both orientations: runtime.GOOS == "x" and "x" == runtime.GOOS.
	c.check(binExpr, binExpr.X, binExpr.Y)
	c.check(binExpr, binExpr.Y, binExpr.X)
}

func (c *goosChecker) check(cause *ast.BinaryExpr, runtimeSide, valueSide ast.Expr) {
	if qualifiedName(runtimeSide) != "runtime.GOOS" {
		return
	}

	tv, ok := c.ctx.TypesInfo.Types[valueSide]
	if !ok || tv.Value == nil || tv.Value.Kind() != constant.String {
		return
	}
	val := constant.StringVal(tv.Value)

	// Don't warn on empty string — could be a zero-value guard.
	if val == "" {
		return
	}

	if !knownGOOS[val] {
		c.ctx.Warn(cause, "unknown GOOS value %q", val)
	}
}
