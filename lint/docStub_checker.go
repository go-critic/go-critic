package lint

import (
	"go/ast"
	"regexp"
)

func init() {
	addChecker(&docStubChecker{}, attrSyntaxOnly)
}

type docStubChecker struct {
	checkerBase

	badCommentRE *regexp.Regexp
}

func (c *docStubChecker) Init() {
	re := `//\s?\w+[^a-zA-Z]+$`
	c.badCommentRE = regexp.MustCompile(re)
}

func (c *docStubChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	if decl.Doc != nil && c.badCommentRE.MatchString(decl.Doc.List[0].Text) {
		c.warn(decl)
	}
}

func (c *docStubChecker) warn(decl *ast.FuncDecl) {
	c.ctx.Warn(decl, "silencing go lint doc-comment warnings is unadvised")
}
