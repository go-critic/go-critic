package lint

//! Detects potentially unsafe dependencies on evaluation order.
//
// @Before:
// return mayModifySlice(&xs), xs[0]
//
// @After:
// // A)
// v := mayModifySlice(&xs)
// return v, xs[0]
// // B)
// v := xs[0]
// return mayModifySlice(&xs), v

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/lint/internal/lintutil"
)

func init() {
	addChecker(&evalOrderChecker{}, attrExperimental)
}

type evalOrderChecker struct {
	checkerBase

	// passive represents dependencis that are not harmful on their
	// own, but can result in undefined result if there is
	// equal active dependency for them.
	passive []ast.Expr
	// active represents depencies that can cause undefined evaluation
	// result if there are any equal active or passive dependency for any of them.
	active []ast.Expr
	depSet lintutil.AstSet
}

func (c *evalOrderChecker) VisitStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ReturnStmt:
		c.checkReturn(stmt)
	case *ast.AssignStmt:
		// TODO(quasilyte):
		//	https://github.com/golang/go/issues/23188
		//	https://github.com/golang/go/issues/24448
		//	https://github.com/golang/go/issues/23017
	case *ast.DeclStmt:
		//	TODO(quasilyte): do same handling as in AssignStmt.
	}
}

func (c *evalOrderChecker) reset() {
	c.passive = c.passive[:0]
	c.active = c.active[:0]
}

func (c *evalOrderChecker) checkReturn(ret *ast.ReturnStmt) {
	if len(ret.Results) < 2 {
		return
	}
	c.reset()
	for _, x := range ret.Results {
		c.collectDeps(x)
	}
	// TODO(quasilyte): we can parametrize the threshold later.
	if deps := c.depsCount(); deps != 0 {
		c.warn(ret, deps)
		return
	}
}

func (c *evalOrderChecker) collectCallDeps(call *ast.CallExpr) {
	if c.isSafeCall(call) {
		return
	}

	for _, arg := range call.Args {
		typ := c.ctx.typesInfo.TypeOf(arg)
		if typ == nil {
			continue
		}
		if _, ok := typ.Underlying().(*types.Pointer); ok {
			switch arg := arg.(type) {
			case *ast.UnaryExpr:
				c.active = append(c.active, arg.X)
			default:
				c.active = append(c.active, arg)
			}
		}
	}
}

func (c *evalOrderChecker) collectDeps(x ast.Expr) {
	ast.Inspect(x, func(x ast.Node) bool {
		switch x := x.(type) {
		case *ast.FuncLit:
			return false
		case *ast.CallExpr:
			c.collectCallDeps(x)
		case *ast.IndexExpr:
			c.passive = append(c.passive, x.X)
		}
		return true
	})
}

func (c *evalOrderChecker) isSafeCall(x *ast.CallExpr) bool {
	// Can check many things, but for now,
	// only check for "unsafe" package functions and conversion expressions
	// as they are not real function calls.
	switch qualifiedName(x) {
	case "unsafe.Sizeof", "unsafe.Alignof", "unsafe.Offsetof":
		return true // Unsafe function call
	}
	// May be possible to visit called function body and tell whether
	// it really modifies it's arguments or it's not.
	// Can't mark as stable until these false positives go away.
	return lintutil.IsTypeExpr(c.ctx.typesInfo, x.Fun) // Type conversion
}

func (c *evalOrderChecker) depsCount() int {
	// Every equal active dependency is 1 point.
	//
	// Passive dependencies are only give a point
	// if there is an equal active dependency for it.
	c.depSet.Clear()
	deps := 0
	for _, x := range c.active {
		if !c.depSet.Insert(x) {
			deps++
		}
	}
	for _, x := range c.passive {
		if c.depSet.Contains(x) {
			deps++
		}
	}
	return deps
}

func (c *evalOrderChecker) warn(cause ast.Node, deps int) {
	// Until we have close-to-zero false positives,
	// at least provide confidence/severity level to the user.
	var tag string
	switch deps {
	case 1:
		tag = "low"
	case 2:
		tag = "average"
	default:
		tag = "high"
	}
	c.ctx.Warn(cause, "potential dependency on evaluation order (%s)", tag)
}
