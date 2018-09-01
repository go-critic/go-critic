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
	newMatcherFunc := func(pairs [][2]int) func(*ast.CallExpr) bool {
		return func(call *ast.CallExpr) bool {
			for _, pair := range pairs {
				x := call.Args[pair[0]]
				y := call.Args[pair[1]]
				if !astequal.Expr(x, y) {
					return false
				}
			}
			return true
		}
	}

	xyMatcher := newMatcherFunc([][2]int{
		{0, 1},
	})

	c.matchers = map[string]func(*ast.CallExpr) bool{
		"copy": xyMatcher,

		"strings.Contains":    xyMatcher,
		"strings.Compare":     xyMatcher,
		"strings.EqualFold":   xyMatcher,
		"strings.HasPrefix":   xyMatcher,
		"strings.HasSuffix":   xyMatcher,
		"strings.Index":       xyMatcher,
		"strings.LastIndex":   xyMatcher,
		"strings.Split":       xyMatcher,
		"strings.SplitAfter":  xyMatcher,
		"strings.SplitAfterN": xyMatcher,
		"strings.SplitN":      xyMatcher,

		"bytes.Contains":    xyMatcher,
		"bytes.Compare":     xyMatcher,
		"bytes.Equal":       xyMatcher,
		"bytes.EqualFold":   xyMatcher,
		"bytes.HasPrefix":   xyMatcher,
		"bytes.HasSuffix":   xyMatcher,
		"bytes.Index":       xyMatcher,
		"bytes.LastIndex":   xyMatcher,
		"bytes.Split":       xyMatcher,
		"bytes.SplitAfter":  xyMatcher,
		"bytes.SplitAfterN": xyMatcher,
		"bytes.SplitN":      xyMatcher,

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
