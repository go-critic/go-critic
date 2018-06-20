package astwalk

import (
	"go/ast"
	"go/types"
)

// TODO: try out idea to return false for every visit
// and force checkers to recurse in order to handle siblings.

type FileWalker interface {
	WalkFile(*ast.File)
}

func WalkerForFuncDecl(v FuncDeclVisitor) FileWalker {
	return &funcDeclWalker{visitor: v}
}

func WalkerForExpr(v ExprVisitor) FileWalker {
	return &exprWalker{visitor: v}
}

func WalkerForLocalExpr(v LocalExprVisitor) FileWalker {
	return &localExprWalker{visitor: v}
}

func WalkerForStmtList(v StmtListVisitor) FileWalker {
	return &stmtListWalker{visitor: v}
}

func WalkerForStmt(v StmtVisitor) FileWalker {
	return &stmtWalker{visitor: v}
}

func WalkerForLocalDef(v LocalDefVisitor, info *types.Info) FileWalker {
	return &localDefWalker{visitor: v, info: info}
}

func WalkerForTypeExpr(v TypeExprVisitor, info *types.Info) FileWalker {
	return &typeExprWalker{visitor: v, info: info}
}
