package lint

import (
	"fmt"
	"go/ast"
	"go/types"
)

type unexportedCallChecker struct {
	baseLocalExprChecker

	funcName string
}

func unexportedCallCheck(ctx *context) func(*ast.File) {
	return wrapLocalExprChecker(&unexportedCallChecker{
		baseLocalExprChecker: baseLocalExprChecker{ctx: ctx},
	})
}

func (c *unexportedCallChecker) PerFuncInit(decl *ast.FuncDecl) bool {
	if decl.Body == nil {
		return false
	}
	c.funcName = ""
	cond := decl.Recv != nil &&
		len(decl.Recv.List) == 1 &&
		len(decl.Recv.List[0].Names) == 1
	if cond {
		c.funcName = decl.Recv.List[0].Names[0].Name
	}
	return true
}

// Check finds calls of unexported method from unexported type
// outside that type.
//
//TODO: update description and warning message
func (c *unexportedCallChecker) CheckLocalExpr(expr ast.Expr) {
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
	typ := c.ctx.TypesInfo.TypeOf(sel.Sel)
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
	if recvTyp.Obj().Name() != c.funcName {
		c.warn(call)
	}
}

func (c *unexportedCallChecker) warn(n *ast.CallExpr) {
	c.ctx.Warn(n, fmt.Sprintf("%s should be exported",
		nodeString(c.ctx.FileSet, n)))
}
