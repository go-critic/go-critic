package astwalk

import (
	"go/ast"
	"go/token"
	"go/types"
)

type typeExprWalker struct {
	visitor TypeExprVisitor
	info    *types.Info
}

func (w *typeExprWalker) WalkFile(f *ast.File) {
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.FuncDecl); ok {
			if !w.visitor.EnterFunc(decl) {
				continue
			}
		}
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			if !w.visitor.EnterFunc(decl) {
				continue
			}
			w.walkSignature(decl.Type)
			ast.Inspect(decl.Body, w.walk)
		case *ast.GenDecl:
			if decl.Tok == token.IMPORT {
				continue
			}
			ast.Inspect(decl, w.walk)
		}
	}
}

func (w *typeExprWalker) visit(x ast.Expr) bool {
	w.visitor.VisitTypeExpr(x)
	return w.visitor.EnterChilds(x)
}

func (w *typeExprWalker) walk(x ast.Node) bool {
	switch x := x.(type) {
	case *ast.ParenExpr:
		if w.isTypeExpr(x.X) {
			return w.visit(x)
		}
		return true
	case *ast.MapType:
		return w.visit(x)
	case *ast.FuncType:
		return w.visit(x)
	case *ast.StructType:
		return w.visit(x)
	case *ast.InterfaceType:
		if !w.visit(x) {
			return false
		}
		for _, method := range x.Methods.List {
			switch x := method.Type.(type) {
			case *ast.FuncType:
				w.walkSignature(x)
			default:
				// Embedded interface.
				w.walk(x)
			}
		}
		return false
	case *ast.ArrayType:
		return w.visit(x)
	}
	return true
}

func (w *typeExprWalker) isTypeExpr(x ast.Expr) bool {
	switch x := x.(type) {
	case *ast.Ident:
		// Identifier may be a type expression if object
		// it reffers to is a type name.
		_, ok := w.info.ObjectOf(x).(*types.TypeName)
		return ok

	case *ast.FuncType, *ast.StructType, *ast.InterfaceType, *ast.ArrayType, *ast.MapType:
		return true

	default:
		return false
	}
}

func (w *typeExprWalker) walkSignature(typ *ast.FuncType) {
	for _, p := range typ.Params.List {
		ast.Inspect(p.Type, w.walk)
	}
	if typ.Results != nil {
		for _, p := range typ.Results.List {
			ast.Inspect(p.Type, w.walk)
		}
	}
}
