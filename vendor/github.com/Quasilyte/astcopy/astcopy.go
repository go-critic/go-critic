package astcopy

import (
	"go/ast"
)

// This package is not yet pushed to separate repo.
// It will be vendored until it's API is fully stable.
//
// TODO:
// - copying with objects?
// - copying with comments?
// - handle other nil cases properly
// - verify correctness
// - functions per every node type to avoid type assertions on call site?

// Node returns x deep copy.
func Node(x ast.Node) ast.Node {
	return nodeCopy(x)
}

// Expr returns x deep copy.
func Expr(x ast.Expr) ast.Expr {
	return copyExpr(x)
}

// Stmt returns x deep copy.
func Stmt(x ast.Stmt) ast.Stmt {
	return copyStmt(x)
}

// Decl returns x deep copy.
func Decl(x ast.Decl) ast.Decl {
	return copyDecl(x)
}

func nodeCopy(x ast.Node) ast.Node {
	switch x := x.(type) {
	case ast.Expr:
		return copyExpr(x)
	case ast.Stmt:
		return copyStmt(x)
	case ast.Decl:
		return copyDecl(x)
	case *ast.FieldList:
		return copyFieldList(x)
	default:
		panic("unhandled node")
	}
}

func copyFuncType(x *ast.FuncType) *ast.FuncType {
	cp := *x
	cp.Params = copyFieldList(x.Params)
	cp.Results = copyFieldList(x.Results)
	return &cp
}

func copyBlockStmt(x *ast.BlockStmt) *ast.BlockStmt {
	cp := *x
	cp.List = copyStmtList(x.List)
	return &cp
}

func copyStmtList(xs []ast.Stmt) []ast.Stmt {
	if xs == nil {
		return nil
	}
	cp := make([]ast.Stmt, len(xs))
	for i := range xs {
		cp[i] = copyStmt(xs[i])
	}
	return cp
}

func copyExprList(xs []ast.Expr) []ast.Expr {
	if xs == nil {
		return nil
	}
	cp := make([]ast.Expr, len(xs))
	for i := range xs {
		cp[i] = copyExpr(xs[i])
	}
	return cp
}

func copyIdent(x *ast.Ident) *ast.Ident {
	cp := *x
	return &cp
}

func copyBasicLit(x *ast.BasicLit) *ast.BasicLit {
	if x == nil {
		return nil
	}
	cp := *x
	return &cp
}

func copyIdentList(xs []*ast.Ident) []*ast.Ident {
	if xs == nil {
		return nil
	}
	cp := make([]*ast.Ident, len(xs))
	for i := range xs {
		cp[i] = copyIdent(xs[i])
	}
	return cp
}

func copyField(x *ast.Field) *ast.Field {
	cp := *x
	cp.Names = copyIdentList(x.Names)
	cp.Type = copyExpr(x.Type)
	cp.Tag = copyBasicLit(x.Tag)
	return &cp
}

func copyFieldList(x *ast.FieldList) *ast.FieldList {
	if x == nil {
		return nil
	}
	cp := *x
	if x.List != nil {
		cp.List = make([]*ast.Field, len(x.List))
		for i := range x.List {
			cp.List[i] = copyField(x.List[i])
		}
	}
	return &cp
}

func copyCallExpr(x *ast.CallExpr) *ast.CallExpr {
	cp := *x
	cp.Fun = copyExpr(x.Fun)
	cp.Args = copyExprList(x.Args)
	return &cp
}

func copySpec(x ast.Spec) ast.Spec {
	switch x := x.(type) {
	case *ast.ImportSpec:
		cp := *x
		cp.Name = copyIdent(x.Name)
		cp.Path = copyBasicLit(x.Path)
		return &cp
	case *ast.ValueSpec:
		cp := *x
		cp.Names = copyIdentList(x.Names)
		cp.Values = copyExprList(x.Values)
		cp.Type = copyExpr(x.Type)
		return &cp
	case *ast.TypeSpec:
		cp := *x
		cp.Name = copyIdent(x.Name)
		cp.Type = copyExpr(x.Type)
		return &cp
	default:
		panic("unhandled spec")
	}
}

func copySpecList(xs []ast.Spec) []ast.Spec {
	cp := make([]ast.Spec, len(xs))
	for i := range xs {
		cp[i] = copySpec(xs[i])
	}
	return cp
}

func copyExpr(x ast.Expr) ast.Expr {
	if x == nil {
		return nil
	}

	switch x := x.(type) {
	case *ast.BadExpr:
		cp := *x
		return &cp
	case *ast.Ident:
		return copyIdent(x)
	case *ast.Ellipsis:
		cp := *x
		cp.Elt = copyExpr(x.Elt)
		return &cp
	case *ast.BasicLit:
		return copyBasicLit(x)
	case *ast.FuncLit:
		cp := *x
		cp.Type = copyFuncType(x.Type)
		cp.Body = copyBlockStmt(x.Body)
		return &cp
	case *ast.CompositeLit:
		cp := *x
		cp.Type = copyExpr(x.Type)
		cp.Elts = copyExprList(x.Elts)
		return &cp
	case *ast.ParenExpr:
		cp := *x
		cp.X = copyExpr(x.X)
		return &cp
	case *ast.SelectorExpr:
		return &ast.SelectorExpr{
			X:   copyExpr(x.X),
			Sel: copyIdent(x.Sel),
		}
	case *ast.IndexExpr:
		cp := *x
		cp.X = copyExpr(x.X)
		cp.Index = copyExpr(x.Index)
		return &cp
	case *ast.SliceExpr:
		cp := *x
		cp.X = copyExpr(x.X)
		cp.Low = copyExpr(x.Low)
		cp.High = copyExpr(x.High)
		cp.Max = copyExpr(x.Max)
		return &cp
	case *ast.TypeAssertExpr:
		cp := *x
		cp.X = copyExpr(x.X)
		cp.Type = copyExpr(x.Type)
		return &cp
	case *ast.CallExpr:
		return copyCallExpr(x)
	case *ast.StarExpr:
		cp := *x
		cp.X = copyExpr(x.X)
		return &cp
	case *ast.UnaryExpr:
		cp := *x
		cp.X = copyExpr(x.X)
		return &cp
	case *ast.BinaryExpr:
		cp := *x
		cp.X = copyExpr(x.X)
		cp.Y = copyExpr(x.Y)
		return &cp
	case *ast.KeyValueExpr:
		cp := *x
		cp.Key = copyExpr(x.Key)
		cp.Value = copyExpr(x.Value)
		return &cp
	case *ast.ArrayType:
		cp := *x
		cp.Len = copyExpr(x.Len)
		cp.Elt = copyExpr(x.Elt)
		return &cp
	case *ast.StructType:
		cp := *x
		cp.Fields = copyFieldList(x.Fields)
		return &cp
	case *ast.FuncType:
		return copyFuncType(x)
	case *ast.InterfaceType:
		cp := *x
		cp.Methods = copyFieldList(x.Methods)
		return &cp
	case *ast.MapType:
		cp := *x
		cp.Key = copyExpr(x.Key)
		cp.Value = copyExpr(x.Value)
		return &cp
	case *ast.ChanType:
		cp := *x
		cp.Value = copyExpr(x.Value)
		return &cp
	default:
		// panic(fmt.Sprintf("unhandled expr: %T", x))
		panic("unhandled expr")
	}
}

func copyStmt(x ast.Stmt) ast.Stmt {
	switch x := x.(type) {
	case *ast.BadStmt:
		cp := *x
		return &cp
	case *ast.DeclStmt:
		return &ast.DeclStmt{Decl: copyDecl(x.Decl)}
	case *ast.EmptyStmt:
		cp := *x
		return &cp
	case *ast.LabeledStmt:
		cp := *x
		cp.Label = copyIdent(x.Label)
		cp.Stmt = copyStmt(x.Stmt)
		return &cp
	case *ast.ExprStmt:
		return &ast.ExprStmt{X: copyExpr(x.X)}
	case *ast.SendStmt:
		cp := *x
		cp.Chan = copyExpr(x.Chan)
		cp.Value = copyExpr(x.Value)
		return &cp
	case *ast.IncDecStmt:
		cp := *x
		cp.X = copyExpr(x.X)
		return &cp
	case *ast.AssignStmt:
		cp := *x
		cp.Lhs = copyExprList(x.Lhs)
		cp.Rhs = copyExprList(x.Rhs)
		return &cp
	case *ast.GoStmt:
		cp := *x
		cp.Call = copyCallExpr(x.Call)
		return &cp
	case *ast.DeferStmt:
		cp := *x
		cp.Call = copyCallExpr(x.Call)
		return &cp
	case *ast.ReturnStmt:
		cp := *x
		cp.Results = copyExprList(x.Results)
		return &cp
	case *ast.BranchStmt:
		cp := *x
		cp.Label = copyIdent(x.Label)
		return &cp
	case *ast.BlockStmt:
		return copyBlockStmt(x)
	case *ast.IfStmt:
		cp := *x
		cp.Init = copyStmt(x.Init)
		cp.Cond = copyExpr(x.Cond)
		cp.Body = copyBlockStmt(x.Body)
		cp.Else = copyStmt(x.Else)
		return &cp
	case *ast.CaseClause:
		cp := *x
		cp.List = copyExprList(x.List)
		cp.Body = copyStmtList(x.Body)
		return &cp
	case *ast.SwitchStmt:
		cp := *x
		cp.Init = copyStmt(x.Init)
		cp.Tag = copyExpr(x.Tag)
		cp.Body = copyBlockStmt(x.Body)
		return &cp
	case *ast.TypeSwitchStmt:
		cp := *x
		cp.Init = copyStmt(x.Init)
		cp.Assign = copyStmt(x.Assign)
		cp.Body = copyBlockStmt(x.Body)
		return &cp
	case *ast.CommClause:
		cp := *x
		cp.Comm = copyStmt(x.Comm)
		cp.Body = copyStmtList(x.Body)
		return &cp
	case *ast.SelectStmt:
		cp := *x
		cp.Body = copyBlockStmt(x.Body)
		return &cp
	case *ast.ForStmt:
		cp := *x
		cp.Init = copyStmt(x.Init)
		cp.Cond = copyExpr(x.Cond)
		cp.Post = copyStmt(x.Post)
		cp.Body = copyBlockStmt(x.Body)
		return &cp
	case *ast.RangeStmt:
		cp := *x
		cp.Key = copyExpr(x.Key)
		cp.Value = copyExpr(x.Value)
		cp.X = copyExpr(x.X)
		cp.Body = copyBlockStmt(x.Body)
		return &cp
	default:
		panic("unhandled stmt")
	}
}

func copyDecl(x ast.Decl) ast.Decl {
	switch x := x.(type) {
	case *ast.BadDecl:
		cp := *x
		return &cp
	case *ast.GenDecl:
		cp := *x
		cp.Specs = copySpecList(x.Specs)
		return &cp
	case *ast.FuncDecl:
		cp := *x
		cp.Recv = copyFieldList(x.Recv)
		cp.Name = copyIdent(x.Name)
		cp.Type = copyFuncType(x.Type)
		cp.Body = copyBlockStmt(x.Body)
		return &cp
	default:
		panic("unhandled decl")
	}
}
