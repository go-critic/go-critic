package lint

import (
	"go/ast"
	"regexp"
	"strings"
)

func init() {
	addChecker(&docStubChecker{}, attrSyntaxOnly)
}

type docStubChecker struct {
	checkerBase

	goodCommentRE *regexp.Regexp
}

func (c *docStubChecker) Init() {
	re := `// \w+ \w+`
	c.goodCommentRE = regexp.MustCompile(re)
}

func (c *docStubChecker) CheckFuncDecl(decl *ast.FuncDecl) {
	if ast.IsExported(decl.Name.Name) && decl.Doc != nil {
		doc := decl.Doc.List[0].Text
		prefix := "// " + decl.Name.Name + " "
		if !strings.HasPrefix(doc, prefix) || !c.goodCommentRE.MatchString(doc) {
			c.warn(decl)
		}
	}
}

func (c *docStubChecker) warn(decl *ast.FuncDecl) {
	c.ctx.Warn(decl, "silencing go lint doc-comment warnings is unadvised")
}
