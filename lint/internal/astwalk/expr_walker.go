package astwalk

import "go/ast"

type exprWalker struct {
	visitor ExprVisitor
}

func (w *exprWalker) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.FuncDecl); ok {
			if !w.visitor.EnterFunc(decl) {
				continue
			}
		}
		ast.Inspect(decl, func(x ast.Node) bool {
			if x, ok := x.(ast.Expr); ok {
				w.visitor.VisitExpr(x)
				return w.visitor.EnterChilds(x)
			}
			return true
		})
	}
}
