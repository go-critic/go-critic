package lint

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
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

func nodeString(fset *token.FileSet, x ast.Node) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
