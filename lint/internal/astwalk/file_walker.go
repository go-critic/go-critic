package astwalk

import "go/ast"

type fileWalker struct {
	visitor FileVisitor
}

func (w *fileWalker) WalkFile(f *ast.File) {
	w.visitor.VisitFile(f)
}
