package astwalk

import (
	"go/ast"
	"go/types"
)

// FileWalker traverses given file and calls associated visitor
// for every node it's interested in.
type FileWalker interface {
	WalkFile(*ast.File)
}

// WalkerForFile returns file walker implementation for FileVisitor.
func WalkerForFile(v FileVisitor) FileWalker {
	return &fileWalker{visitor: v}
}

// WalkerForFuncDecl returns file walker implementation for FuncDeclVisitor.
func WalkerForFuncDecl(v FuncDeclVisitor) FileWalker {
	return &funcDeclWalker{visitor: v}
}

// WalkerForExpr returns file walker implementation for ExprVisitor.
func WalkerForExpr(v ExprVisitor) FileWalker {
	return &exprWalker{visitor: v}
}

// WalkerForLocalExpr returns file walker implementation for LocalExprVisitor.
func WalkerForLocalExpr(v LocalExprVisitor) FileWalker {
	return &localExprWalker{visitor: v}
}

// WalkerForStmtList returns file walker implementation for StmtListVisitor.
func WalkerForStmtList(v StmtListVisitor) FileWalker {
	return &stmtListWalker{visitor: v}
}

// WalkerForStmt returns file walker implementation for StmtVisitor.
func WalkerForStmt(v StmtVisitor) FileWalker {
	return &stmtWalker{visitor: v}
}

// WalkerForLocalDef returns file walker implementation for LocalDefVisitor.
func WalkerForLocalDef(v LocalDefVisitor, info *types.Info) FileWalker {
	return &localDefWalker{visitor: v, info: info}
}

// WalkerForTypeExpr returns file walker implementation for TypeExprVisitor.
func WalkerForTypeExpr(v TypeExprVisitor, info *types.Info) FileWalker {
	return &typeExprWalker{visitor: v, info: info}
}

// WalkerForLocalComment returns file walker implementation for LocalCommentVisitor.
func WalkerForLocalComment(v LocalCommentVisitor) FileWalker {
	return &localCommentVisitor{visitor: v}
}
