package lint

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
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

func paramNamesStr(idents []*ast.Ident) string {
	if idents == nil {
		return "_"
	}
	names := []string{}
	for _, id := range idents {
		names = append(names, id.Name)
	}
	return strings.Join(names, " ,")
}
