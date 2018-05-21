package lint

import (
	"fmt"
	"go/ast"
	"go/types"
)

type unexportedCallChecker struct {
	ctx *Context
}

func newUnexportedCallChecker(ctx *Context) Checker {
	return &unexportedCallChecker{
		ctx: ctx,
	}
}

// Check finds calls of unexported method from unexported type
// outside that type.
//
//TODO: update description and warning message
func (c *unexportedCallChecker) Check(f *ast.File) {
	for _, decl := range collectFuncDecls(f) {
		if decl.Body == nil {
			continue
		}
		name := ""
		cond := decl.Recv != nil &&
			len(decl.Recv.List) == 1 &&
			len(decl.Recv.List[0].Names) == 1
		if cond {
			name = decl.Recv.List[0].Names[0].Name
		}
		ast.Inspect(decl.Body, func(n ast.Node) bool {
			if call, ok := n.(*ast.CallExpr); ok {
				c.checkCall(name, call)
			}
			return true
		})

	}
}

func (c *unexportedCallChecker) checkCall(name string, call *ast.CallExpr) {
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
	if recvTyp.Obj().Name() != name {
		c.warn(call)
	}
}

func (c *unexportedCallChecker) warn(n *ast.CallExpr) {
	c.ctx.addWarning(Warning{
		Kind: "unexported-call",
		Node: n,
		Text: fmt.Sprintf("%s should be exported", nodeString(c.ctx.FileSet, n)),
	})
}
