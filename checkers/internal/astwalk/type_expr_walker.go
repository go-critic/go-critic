package astwalk

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/go-toolsmith/astp"
	"github.com/go-toolsmith/typep"
)

type typeExprWalker struct {
	visitor TypeExprVisitor
	info    *types.Info
}

func (w *typeExprWalker) WalkFile(f *ast.File) {
	if !w.visitor.EnterFile(f) {
		return
	}

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
	return !w.visitor.skipChilds()
}

func (w *typeExprWalker) walk(x ast.Node) bool {
	switch x := x.(type) {
	case *ast.ChanType:
		return w.visit(x)
	case *ast.ParenExpr:
		if typep.IsTypeExpr(w.info, x.X) {
			return w.visit(x)
		}
		return true
	case *ast.CallExpr:
		// Pointer conversions require parenthesis around pointer type.
		// These casts are represented as call expressions.
		// Because it's impossible for the visitor to distinguish such
		// "required" parenthesis, walker skips outmost parenthesis in such cases.
		return w.inspectInner(x.Fun)
	case *ast.SelectorExpr:
		// Like with conversions, method expressions are another special.
		return w.inspectInner(x.X)
	case *ast.StarExpr:
		if typep.IsTypeExpr(w.info, x.X) {
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

// isRecvOnlyChanType reports whether x is a receive-only channel type, `<-chan T`.
//
// The Go spec requires a parenthesis around a conversion type that starts with
// `*`, `<-`, or the `func` keyword: without it `(<-chan int)(nil)` parses as
// `<-(chan int(nil))` and does not compile. A send-only `chan<- int(nil)` is
// unambiguous, so only the receive direction is special here.
func isRecvOnlyChanType(x ast.Expr) bool {
	ch, ok := x.(*ast.ChanType)
	return ok && ch.Dir == ast.RECV
}

func (w *typeExprWalker) inspectInner(x ast.Expr) bool {
	parens, ok := x.(*ast.ParenExpr)
	shouldInspect := ok &&
		typep.IsTypeExpr(w.info, parens.X) &&
		(astp.IsStarExpr(parens.X) || astp.IsFuncType(parens.X) || isRecvOnlyChanType(parens.X))
	if shouldInspect {
		ast.Inspect(parens.X, w.walk)
		return false
	}
	return true
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
