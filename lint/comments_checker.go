// new checker:
// 1. create file lint/foo_checker.go
// 2. choose base implementation from wrappers.go
// 3. define checker type that embeds base implementation
// 4. declare constructor for check function
// 5. add checker entry to checkFunctions from lint/lint.go
// 6. add tests in cmd/kfulint/testdata/foo/checker_tests.go
// 7. implement checker

package lint

import (
	"go/ast"
	"regexp"
)

type commentsChecker struct {
	ctx    *context
	regexp *regexp.Regexp
}

func commentsCheck(ctx *context) func(*ast.File) {
	re := `//\s?\w+[^a-zA-Z]+$`
	c := commentsChecker{ctx: ctx, regexp: regexp.MustCompile(re)}
	return c.CheckFile
}

func (c *commentsChecker) CheckFile(f *ast.File) {
	// 1. loop over f.Decls
	// 2. find func decls
	// 3. for each funcdecl check decl.Doc
	// 4. if decl.Doc is not nil, check its text

	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.FuncDecl); ok {
			if decl.Doc != nil {
				// fmt.Println(decl.Doc.Text())
				//fmt.Printf("%q\n", decl.Doc.List[0].Text)
				if c.regexp.MatchString(decl.Doc.List[0].Text) {
					c.warn(decl)
				}
			}
		}
	}
}

func (c *commentsChecker) warn(decl *ast.FuncDecl) {
	c.ctx.Warn(decl, "comment should contain function name")
}

/*

// decl.Doc
	var decl *ast.FuncDecl
	decl.Doc.Text() // => comment text

// Foo ...
func Foo() {
}

"Foo ..."

decl.Doc.Text() == decl.Name.Name + " ..."
"Foo ..." == "Foo ..."

*/
