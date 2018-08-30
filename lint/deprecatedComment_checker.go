package lint

import (
	"go/ast"
	"regexp"
	"strings"
)

func init() {
	addChecker(&deprecatedCommentChecker{}, attrExperimental, attrSyntaxOnly)
}

type deprecatedCommentChecker struct {
	checkerBase
}

func (c *deprecatedCommentChecker) InitDocumentation(d *Documentation) {
	d.Summary = `Detects malformed "deprecated" doc-comments`
	d.Before = `
// deprecated, use FuncNew instead
func FuncOld() int
`
	d.After = `
// Deprecated: use FuncNew instead
func FuncOld() int
`
}

func (c *deprecatedCommentChecker) VisitDocComment(doc *ast.CommentGroup) {
	// There are 3 accepted forms of deprecation comments:
	//
	// 1. inline, that can't be handled with a DocCommentVisitor.
	//    Note that "Deprecated: " may not even be the comment prefix there.
	//    Example: "The line number in the input. Deprecated: Kept for compatibility."
	//    TODO(quasilyte): fix it.
	//
	// 2. Longer form-1. It's a doc-comment that only contains "deprecation" notice.
	//
	// 3. Like form-2, but may also include doc-comment text.
	//    Distinguished by an empty line.
	//
	// See https://github.com/golang/go/issues/10909#issuecomment-136492606.
	//
	// It's desirable to see how people make mistakes with the format,
	// this is why there is currently no special treatment for these cases.
	// TODO(quasilyte): do more audits and grow the negative tests suite.
	//
	// TODO(quasilyte): there are also multi-line deprecation comments.

	commonPatterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)this (?:function|type) is deprecated`),
		regexp.MustCompile(`(?i)deprecated[.!]? use \S* instead`),
		// TODO(quasilyte): more of these?
	}

	for _, l := range strings.Split(doc.Text(), "\n") {
		if len(l) < len("Deprecated: ") {
			continue
		}

		// Check whether someone messed up with a prefix casing.
		upcase := strings.ToUpper(l)
		if strings.HasPrefix(upcase, "DEPRECATED: ") && !strings.HasPrefix(l, "Deprecated: ") {
			c.warnCasing(doc, l)
			return
		}

		// Check is someone used comma instead of a colon.
		if strings.HasPrefix(l, "Deprecated, ") {
			c.warnComma(doc)
			return
		}

		// Check for other commonly used patterns.
		for _, pat := range commonPatterns {
			if pat.MatchString(l) {
				c.warnPattern(doc)
				return
			}
		}
	}
}

func (c *deprecatedCommentChecker) warnCasing(cause ast.Node, line string) {
	prefix := line[:len("DEPRECATED: ")]
	c.ctx.Warn(cause, "use `Deprecated: ` (note the casing) instead of `%s`", prefix)
}

func (c *deprecatedCommentChecker) warnPattern(cause ast.Node) {
	c.ctx.Warn(cause, "the proper format is `Deprecated: <text>`")
}

func (c *deprecatedCommentChecker) warnComma(cause ast.Node) {
	c.ctx.Warn(cause, "use `:` instead of `,` in `Deprecated, `")
}
