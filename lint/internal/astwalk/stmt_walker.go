package astwalk

import "go/ast"

type stmtWalker struct {
	visitor StmtVisitor
}

func (w *stmtWalker) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok || !w.visitor.EnterFunc(decl) {
			continue
		}
		ast.Inspect(decl.Body, func(x ast.Node) bool {
			if x, ok := x.(ast.Stmt); ok {
				w.visitor.VisitStmt(x)
				return w.visitor.EnterChilds(x)
			}
			return true
		})
	}
}
