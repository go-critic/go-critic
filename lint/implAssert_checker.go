package lint

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"
)

func init() {
	addChecker(implAssertChecker{}, &ruleInfo{})
}

type implAssertChecker struct {
	ctx *context
}

func (c implAssertChecker) New(ctx *context) func(*ast.File) {
	c.ctx = ctx
	return c.CheckFile
}

func (c implAssertChecker) CheckFile(f *ast.File) {
	if strings.HasSuffix(c.ctx.Filename, "_test.go") ||
		strings.HasSuffix(f.Name.Name, "_test") {
		return
	}
	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if decl.Tok != token.VAR {
			continue
		}
		// TODO(cristaloleg): we might want to support multiple assertion in the future.
		if len(decl.Specs) != 1 {
			return
		}
		valueSpec, ok := decl.Specs[0].(*ast.ValueSpec)
		if !ok {
			return
		}
		if len(valueSpec.Names) != 1 || valueSpec.Names[0].Name != "_" {
			return
		}
		if _, ok := c.ctx.TypesInfo.TypeOf(valueSpec.Type).(*types.Interface); !ok {
			return
		}
		if len(valueSpec.Values) != 1 {
			return
		}
		if _, ok := c.ctx.TypesInfo.TypeOf(valueSpec.Values[0]).(*types.Struct); !ok {
			return
		}
		c.warn(decl)
	}
}

func (c *implAssertChecker) warn(decl *ast.GenDecl) {
	c.ctx.Warn(decl, "consider to move implementation assertion to test package or test file")
}
