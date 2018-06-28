package lint

//! Detects defer in loop and warns that it will not be executed till the end of function's scope.
//
// @Before:
// for i := range [10]int{} {
// 	defer f(i) // will be executed only at the end of func
// }
//
// @After:
// for i := range [10]int{} {
// 	func(i int) {
// 		defer f(i)
// 	}(i)
// }

import (
	"go/ast"

	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(&deferInLoopChecker{}, attrExperimental)
}

type deferInLoopChecker struct {
	checkerBase
}

func (c *deferInLoopChecker) VisitStmt(stmt ast.Stmt) {
	var body *ast.BlockStmt
	loop, ok := stmt.(*ast.RangeStmt)
	if ok {
		body = loop.Body
	} else {
		loop, ok := stmt.(*ast.ForStmt)
		if !ok {
			return
		}
		body = loop.Body
	}
	for _, s := range body.List {
		if astp.IsDeferStmt(s) {
			c.warn(s)
		}
	}
}

func (c *deferInLoopChecker) warn(cause ast.Stmt) {
	c.ctx.Warn(cause, "defer will be executed only at the end of the func's scope")
}
