package astwalk

import "go/ast"

// Visitor interfaces.
type (
	// FuncDeclVisitor visits every top-level function declaration.
	FuncDeclVisitor interface {
		walkerEvents
		VisitFuncDecl(*ast.FuncDecl)
	}

	// ExprVisitor visits every expression inside AST file.
	ExprVisitor interface {
		walkerEvents
		VisitExpr(ast.Expr)
	}

	// LocalExprVisitor visits every expression inside function body.
	LocalExprVisitor interface {
		walkerEvents
		VisitLocalExpr(ast.Expr)
	}

	// StmtListVisitor visits every statement list inside function body.
	// This includes block statement bodies as well as implicit blocks
	// introduced by case clauses and alike.
	StmtListVisitor interface {
		walkerEvents
		VisitStmtList([]ast.Stmt)
	}

	// StmtVisitor visits every statement inside function body.
	StmtVisitor interface {
		walkerEvents
		VisitStmt(ast.Stmt)
	}

	// LocalDefVisitor visits every name definitions inside function.
	//
	// Next elements are considered as name definitions:
	//	- Function parameters (input, output, receiver)
	//	- Every LHS of ":=" assignment that defines a new name
	//	- Every local var/const declaration.
	//
	// NOTE: this visitor is somewhat experimental.
	LocalDefVisitor interface {
		walkerEvents
		VisitLocalDef(Name, ast.Expr)
	}

	// TypeExprVisitor visits every type describing expression.
	// It also traverses struct types and interface types to run
	// checker over their fields/method signatures.
	TypeExprVisitor interface {
		walkerEvents
		VisitTypeExpr(ast.Expr)
	}
)

// walkerEvents describes common hooks available for every visitor.
type walkerEvents interface {
	// EnterFunc is called for every function declaration that is about
	// to be traversed. If false is returned, function is not visited.
	EnterFunc(*ast.FuncDecl) bool

	// EnterChilds is called for every visited node.
	// Node that was visited is passed as an argument.
	// If visitor returns false, that node siblings will not be traversed.
	//
	// Not applicable to:
	//	- FuncDeclVisitor
	//	- StmtListVisitor
	//	- LocalDefVisitor
	EnterChilds(ast.Node) bool
}

var _ = walkerEvents(nil) // Make sure walkerEvents not marked as unused by megacheck

type (
	// NameKind describes what kind of name Name object holds.
	NameKind int

	// Name holds ver/const/param definition symbol info.
	Name struct {
		ID   *ast.Ident
		Kind NameKind

		// Index is NameVar-specific field that is used to
		// specify nth tuple element being assigned to the name.
		Index int
	}
)

// NOTE: set of name kinds is not stable and may change over time.
//
// TODO(quasilyte): is NameRecv/NameParam/NameResult granularity desired?
// TODO(quasilyte): is NameVar/NameBind (var vs :=) granularity desired?
const (
	// NameParam is function/method receiver/input/output name.
	// Initializing expression is always nil.
	NameParam NameKind = iota
	// NameVar is var or ":=" declared name.
	// Initizlizing expression may be nil for var-declared names
	// without explicit initializing expression.
	NameVar
	// NameConst is const-declared name.
	// Initializing expression is never nil.
	NameConst
)
