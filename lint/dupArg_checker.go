package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astequal"
)

func init() {
	addChecker(&dupArgChecker{}, attrExperimental)
}

type dupArgChecker struct {
	checkerBase

	matchers map[string]func(*ast.CallExpr) bool
}

func (c *dupArgChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects suspicious duplicated arguments"
	d.Before = `copy(dst, dst)`
	d.After = `copy(dst, src)`
}

func (c *dupArgChecker) Init() {
	// newMatcherFunc returns a function that matches a call if
	// args[xIndex] and args[yIndex] are equal.
	newMatcherFunc := func(xIndex, yIndex int) func(*ast.CallExpr) bool {
		return func(call *ast.CallExpr) bool {
			x := call.Args[xIndex]
			y := call.Args[yIndex]
			return astequal.Expr(x, y)
		}
	}

	// m maps pattern string to a matching function.
	// String patterns are used for documentation purposes (readability).
	m := map[string]func(*ast.CallExpr) bool{
		"(x, x, ...)":    newMatcherFunc(0, 1),
		"(x, _, x, ...)": newMatcherFunc(0, 2),
	}

	// TODO(quasilyte): handle x.Equal(x) cases.
	// Example: *math/Big.Int.Cmp method.

	// TODO(quasilyte): more perky mode that will also
	// report things like io.Copy(x, x).
	// Probably safe thing to do even without that option
	// if `x` is not interface (requires type checks
	// that are not incorporated into this checker yet).

	c.matchers = map[string]func(*ast.CallExpr) bool{
		"copy": m["(x, x, ...)"],

		"reflect.Copy":      m["(x, x, ...)"],
		"reflect.DeepEqual": m["(x, x, ...)"],

		"strings.Contains":    m["(x, x, ...)"],
		"strings.Compare":     m["(x, x, ...)"],
		"strings.EqualFold":   m["(x, x, ...)"],
		"strings.HasPrefix":   m["(x, x, ...)"],
		"strings.HasSuffix":   m["(x, x, ...)"],
		"strings.Index":       m["(x, x, ...)"],
		"strings.LastIndex":   m["(x, x, ...)"],
		"strings.Split":       m["(x, x, ...)"],
		"strings.SplitAfter":  m["(x, x, ...)"],
		"strings.SplitAfterN": m["(x, x, ...)"],
		"strings.SplitN":      m["(x, x, ...)"],

		"bytes.Contains":    m["(x, x, ...)"],
		"bytes.Compare":     m["(x, x, ...)"],
		"bytes.Equal":       m["(x, x, ...)"],
		"bytes.EqualFold":   m["(x, x, ...)"],
		"bytes.HasPrefix":   m["(x, x, ...)"],
		"bytes.HasSuffix":   m["(x, x, ...)"],
		"bytes.Index":       m["(x, x, ...)"],
		"bytes.LastIndex":   m["(x, x, ...)"],
		"bytes.Split":       m["(x, x, ...)"],
		"bytes.SplitAfter":  m["(x, x, ...)"],
		"bytes.SplitAfterN": m["(x, x, ...)"],
		"bytes.SplitN":      m["(x, x, ...)"],

		"types.Identical":           m["(x, x, ...)"],
		"types.IdenticalIgnoreTags": m["(x, x, ...)"],

		"draw.Draw": m["(x, _, x, ...)"],

		// TODO(quasilyte): more of these.
	}
}

func (c *dupArgChecker) VisitExpr(expr ast.Expr) {
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return
	}

	m := c.matchers[qualifiedName(call.Fun)]
	if m != nil && m(call) {
		c.warn(call)
	}
}

func (c *dupArgChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "suspicious duplicated args in `%s`", cause)
}
