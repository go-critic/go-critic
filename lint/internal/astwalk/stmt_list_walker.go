package astwalk

import (
	"go/ast"
)

type stmtListWalker struct {
	visitor StmtListVisitor
}

func (w *stmtListWalker) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok || !w.visitor.EnterFunc(decl) {
			continue
		}
		ast.Inspect(decl.Body, func(x ast.Node) bool {
			switch x := x.(type) {
			case *ast.BlockStmt:
				w.visitor.VisitStmtList(x.List)
			case *ast.CaseClause:
				w.visitor.VisitStmtList(x.Body)
			case *ast.CommClause:
				w.visitor.VisitStmtList(x.Body)
			}
			return true
		})
	}
}
