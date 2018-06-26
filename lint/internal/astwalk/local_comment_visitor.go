package astwalk

import "go/ast"

type localCommentVisitor struct {
	visitor LocalCommentVisitor
}

func (w *localCommentVisitor) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok || !w.visitor.EnterFunc(decl) {
			continue
		}
		begin := decl.Pos()
		end := decl.End()
		for _, c := range f.Comments {
			// Not sure that decls/comments are sorted
			// by positions, so do a naive full scan for now.
			if c.Pos() < begin || c.Pos() > end {
				continue
			}
			w.visitor.VisitLocalComment(c)
		}
	}
}
