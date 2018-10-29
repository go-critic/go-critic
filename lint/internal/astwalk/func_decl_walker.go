package astwalk

import (
	"go/ast"
)

type funcDeclWalker struct {
	visitor FuncDeclVisitor
}

func (w *funcDeclWalker) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok || !w.visitor.EnterFunc(decl) {
			continue
		}
		w.visitor.VisitFuncDecl(decl)
	}
}
