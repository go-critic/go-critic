package lint

import (
	"go/ast"
	"go/constant"
	"strings"

	"github.com/go-critic/go-critic/lint/internal/lintutil"
)

func init() {
	addChecker(&flagNameChecker{}, attrExperimental)
}

type flagNameChecker struct {
	checkerBase
}

func (c *flagNameChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects flag names with whitespace"
	d.Before = `b := flag.Bool(" foo ", false, "description")`
	d.After = `b := flag.Bool("foo", false, "description")`
}

func (c *flagNameChecker) VisitExpr(expr ast.Expr) {
	call := lintutil.AsCallExpr(expr)
	switch qualifiedName(call.Fun) {
	case "flag.Bool", "flag.Duration", "flag.Float64", "flag.String",
		"flag.Int", "flag.Int64", "flag.Uint", "flag.Uint64":
		c.checkFlagName(call, call.Args[0])
	case "flag.BoolVar", "flag.DurationVar", "flag.Float64Var", "flag.StringVar",
		"flag.IntVar", "flag.Int64Var", "flag.UintVar", "flag.Uint64Var":
		c.checkFlagName(call, call.Args[1])
	}
}

func (c *flagNameChecker) checkFlagName(call *ast.CallExpr, arg ast.Expr) {
	cv := c.ctx.typesInfo.Types[arg].Value
	if cv == nil {
		return // Non-constant name
	}
	name := constant.StringVal(cv)
	if strings.Contains(name, " ") {
		c.warnWhitespace(call, name)
	}
}

func (c *flagNameChecker) warnWhitespace(cause ast.Node, name string) {
	c.ctx.Warn(cause, "flag name %q contains whitespace", name)
}
