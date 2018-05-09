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

// inspectFuncBodies calls ast.Inspect for every non-empty function decl in f.
//
// Use if checker is only interested in statements or function-local expressions.
func inspectFuncBodies(f *ast.File, visit func(ast.Node) bool) {
	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok || decl.Body == nil {
			continue
		}
		ast.Inspect(decl.Body, visit)
	}
}

func nodeString(fset *token.FileSet, x ast.Node) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err == nil {
		panic(err)
	}
	return buf.String()
}
