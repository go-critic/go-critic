package checkers

import (
	"go/ast"
	"strings"
	"unicode"

	"github.com/go-toolsmith/astcast"

	"github.com/go-toolsmith/astp"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "hexLiterals"
	info.Tags = []string{"style", "experimental"}
	info.Summary = ""
	info.Before = `
x := 0X12
y := 0xfF`
	info.After = `
x := 0x12
// (A)
y := 0xff
// (B)
y := 0xFF`

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return astwalk.WalkerForExpr(&hexLiteralChecker{ctx: ctx})
	})
}

type hexLiteralChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext
}

func (c *hexLiteralChecker) VisitExpr(x ast.Expr) {
	if !astp.IsBasicLit(x) {
		return
	}
	v := astcast.ToBasicLit(x)

	if !strings.Contains(v.Value, "0X") && !strings.Contains(v.Value, "0x") {
		return
	}

	value := v.Value[2:]

	if isAnyLetter(value) {
		c.ctx.Warn(x,
			"Should be 0x%s or 0x%s",
			strings.ToLower(v.Value[2:]),
			strings.ToUpper(v.Value[2:]))
		return
	}

	c.ctx.Warn(x,
		"Should be 0x%s",
		strings.ToLower(v.Value[2:]))
}

func isAnyLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
