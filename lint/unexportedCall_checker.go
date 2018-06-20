package lint

//! Detects calls of unexported method from unexported type outside that type.
//
// Before:
// type foo struct{}
// func (f foo) bar() int { return 1 }
// func baz() {
// 	var fo foo
// 	fo.bar()
// }
//
// After:
// type foo struct{}
// func (f foo) Bar() int { return 1 } // now Bar is exported
// func baz() {
// 	var fo foo
// 	fo.Bar()
// }

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(&unexportedCallChecker{}, attrVeryOpinionated)
}

type unexportedCallChecker struct {
	checkerBase

	recvName string
}

func (c *unexportedCallChecker) EnterFunc(decl *ast.FuncDecl) bool {
	if decl.Body == nil {
		return false
	}
	c.recvName = ""
	cond := decl.Recv != nil &&
		len(decl.Recv.List) == 1 &&
		len(decl.Recv.List[0].Names) == 1
	if cond {
		c.recvName = decl.Recv.List[0].Names[0].Name
	}
	return true
}

// TODO: update description and warning message
func (c *unexportedCallChecker) VisitLocalExpr(expr ast.Expr) {
	if call, ok := expr.(*ast.CallExpr); ok {
		c.checkCall(call)
	}
}

func (c *unexportedCallChecker) checkCall(call *ast.CallExpr) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}
	if sel.Sel.IsExported() {
		return
	}
	typ := c.ctx.typesInfo.TypeOf(sel.Sel)
	sig, ok := typ.(*types.Signature)
	if !ok {
		return
	}

	if sig.Recv() == nil || sig.Recv().Type() == nil {
		return
	}
	recvTyp, ok := sig.Recv().Type().(*types.Named)
	if !ok {
		return
	}
	if recvTyp.Obj().Name() != c.recvName {
		c.warn(call)
	}
}

func (c *unexportedCallChecker) warn(n *ast.CallExpr) {
	c.ctx.Warn(n, "%s should be exported", n)
}
