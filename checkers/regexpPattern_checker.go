package checkers

import (
	"go/ast"
	"go/constant"
	"regexp"
	"strings"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "regexpPattern"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects suspicious regexp patterns"
	info.Before = "regexp.MustCompile(`google.com|yandex.ru`)"
	info.After = "regexp.MustCompile(`google\\.com|yandex\\.ru`)"

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		domains := []string{
			"com",
			"org",
			"info",
			"net",
			"ru",
			"de",
		}

		allDomains := strings.Join(domains, "|")
		domainRE := regexp.MustCompile(`[^\\]\.(` + allDomains + `)\b`)
		return astwalk.WalkerForExpr(&regexpPatternChecker{
			ctx:           ctx,
			domainRE:      domainRE,
			beginAnchorRE: regexp.MustCompile(`[^\\^]\^`),
			endAnchorRE:   regexp.MustCompile(`[^\\]\$.+`),
		})
	})
}

type regexpPatternChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext

	domainRE      *regexp.Regexp
	beginAnchorRE *regexp.Regexp
	endAnchorRE   *regexp.Regexp
}

type regexpFlags struct {
	i bool // case-insensitive
	m bool // multi-line mode: ^ and $ match begin/end line in addition to begin/end text
	s bool // let . match \n
	U bool // ungreedy: swap meaning of x* and x*?, x+ and x+?, etc
}

func (c *regexpPatternChecker) VisitExpr(x ast.Expr) {
	call, ok := x.(*ast.CallExpr)
	if !ok {
		return
	}

	switch qualifiedName(call.Fun) {
	case "regexp.Compile", "regexp.MustCompile", "regexp.CompilePOSIX", "regexp.MustCompilePOSIX":
		c.checkPattern(call.Args[0])
	}
}

func (c *regexpPatternChecker) checkPattern(arg ast.Expr) {
	cv := c.ctx.TypesInfo.Types[arg].Value
	if cv == nil || cv.Kind() != constant.String {
		return
	}
	pat := constant.StringVal(cv)

	if m := c.domainRE.FindStringSubmatch(pat); m != nil {
		c.warnDomain(arg, m[1])
	}

	flags, offset := c.regexpFlags(pat)
	if flags.m {
		return
	}
	pat = pat[offset:] // Drop the flags group
	if c.beginAnchorRE.MatchString(pat) {
		c.warnAnchor(arg, "^")
	}
	if c.endAnchorRE.MatchString(pat) {
		c.warnAnchor(arg, "$")
	}
}

func (c *regexpPatternChecker) regexpFlags(pat string) (regexpFlags, int) {
	var flags regexpFlags

	if !strings.HasPrefix(pat, "(?") {
		return flags, 0
	}

	offset := len("(?")
	for offset < len(pat)-1 && pat[offset] != ')' {
		switch pat[offset] {
		case 'i':
			flags.i = true
		case 'm':
			flags.m = true
		case 's':
			flags.s = true
		case 'U':
			flags.U = true
		default:
			return flags, 0
		}
		offset++
	}
	offset += len(")")
	return flags, offset
}

func (c *regexpPatternChecker) warnAnchor(cause ast.Expr, anchor string) {
	c.ctx.Warn(cause, "unescaped %s in the middle of the regexp", anchor)
}

func (c *regexpPatternChecker) warnDomain(cause ast.Expr, domain string) {
	c.ctx.Warn(cause, "'.%s' should probably be '\\.%s'", domain, domain)
}
