package lint

//! Detects dead code that follow panic/fatal logging.
//
// @Before:
// log.Fatal("exits function")
// return
//
// @After:
// log.Fatal("exits function")

import (
	"go/ast"
)

func init() {
	addChecker(&deadCodeAfterLogFatalChecker{}, attrExperimental)
}

type deadCodeAfterLogFatalChecker struct {
	checkerBase
}

func (c *deadCodeAfterLogFatalChecker) VisitStmtList(stmts []ast.Stmt) {
	for i, stmt := range stmts {
		if stmt, ok := stmt.(*ast.ExprStmt); ok {
			if exprStmt, ok := stmt.X.(*ast.CallExpr); ok {
				switch name := qualifiedName(exprStmt.Fun); name {
				case "log.Fatal", "log.Fatalf", "log.Fatalln", "log.Panic", "log.Panicf", "log.Panicln":
					// if the panic/fatal is not the last statement of the block, we have some dead code.
					if i+1 != len(stmts) {
						c.warn(stmt, name)
					}
				}
			}
		}
	}
}

func (c *deadCodeAfterLogFatalChecker) warn(cause ast.Node, name string) {
	c.ctx.Warn(cause, "remove dead code after '"+name+"'")
}
