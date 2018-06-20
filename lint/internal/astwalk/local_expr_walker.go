package astwalk

import "go/ast"

type localExprWalker struct {
	visitor LocalExprVisitor
}

func (w *localExprWalker) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok || !w.visitor.EnterFunc(decl) {
			continue
		}
		ast.Inspect(decl, func(x ast.Node) bool {
			if x, ok := x.(ast.Expr); ok {
				w.visitor.VisitLocalExpr(x)
				return w.visitor.EnterChilds(x)
			}
			return true
		})
	}
}
