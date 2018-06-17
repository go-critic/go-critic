package lint

//! Detects comments that silence go lint complaints about doc-comment.
//
// Before:
// // Foo ...
// func Foo() {
//      // ...
// }
//
// After:
// func Foo() {
//      // ...
// }
//
// Note:
// > You can either remove a comment to let go lint find it or change stub to useful comment.
// > This checker makes it easier to detect stubs, the action is up to you.

import (
	"go/ast"
	"regexp"
)

func init() {
	addChecker(&docStubChecker{}, attrSyntaxOnly)
}

type docStubChecker struct {
	ctx          *context
	badCommentRE *regexp.Regexp
}

func (c *docStubChecker) New(ctx *context) func(*ast.File) {
	re := `//\s?\w+[^a-zA-Z]+$`
	c.ctx = ctx
	c.badCommentRE = regexp.MustCompile(re)
	return c.CheckFile
}

func (c *docStubChecker) CheckFile(f *ast.File) {
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.FuncDecl); ok {
			if decl.Doc != nil && c.badCommentRE.MatchString(decl.Doc.List[0].Text) {
				c.warn(decl)
			}
		}
	}
}

func (c *docStubChecker) warn(decl *ast.FuncDecl) {
	c.ctx.Warn(decl, "silencing go lint doc-comment warnings is unadvised")
}
