package checkers

import (
	"go/ast"
	"strings"
	"unicode"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/astp"
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

func (c *hexLiteralChecker) VisitExpr(expr ast.Expr) {
	if !astp.IsBasicLit(expr) {
		return
	}
	v := astcast.ToBasicLit(expr)

	if !strings.HasPrefix(v.Value, "0X") && !strings.HasPrefix(v.Value, "0x") {
		return
	}

	prefix := v.Value[:2]
	value := v.Value[2:]

	switch prefix {
	case "0X":
		if isAnyLetter(value) {
			if isGoodHex(value) {
				c.ctx.Warn(expr, "Should be 0x%s", value)
				return
			}
			c.ctx.Warn(expr,
				"Should be 0x%s or 0x%s",
				strings.ToLower(value),
				strings.ToUpper(value))
			return
		}

		c.ctx.Warn(expr, "Should be 0x%s", value)
	case "0x":
		if isAnyLetter(value) {
			if isGoodHex(value) {
				return
			}
			c.ctx.Warn(expr,
				"Should be 0x%s or 0x%s",
				strings.ToLower(value),
				strings.ToUpper(value))
			return
		}

		c.ctx.Warn(expr, "Should be 0x%s", value)
	}
}

func isAnyLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func isGoodHex(value string) bool {
	return value == strings.ToLower(value) || value == strings.ToUpper(value)
}
