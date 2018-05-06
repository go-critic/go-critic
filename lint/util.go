package lint

import (
	"go/ast"
)

func collectFuncDecls(f *ast.File) []*ast.FuncDecl {
	var decls []*ast.FuncDecl
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.FuncDecl); ok {
			decls = append(decls, decl)
		}
	}
	return decls
}
