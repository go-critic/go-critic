package lint

import (
	"go/ast"
)

func switchifCheck(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&switchifChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

type switchifChecker struct {
	baseStmtChecker
}

func (c *switchifChecker) CheckStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.SwitchStmt:
		c.checkSwitchStmt(stmt, stmt.Body)
	case *ast.TypeSwitchStmt:
		c.checkSwitchStmt(stmt, stmt.Body)
	}
}

func (c *switchifChecker) checkSwitchStmt(stmt ast.Stmt, body *ast.BlockStmt) {
	if len(body.List) == 1 {
		if body.List[0].(*ast.CaseClause).List == nil {
			// default case.
			c.warnDefault(stmt)
		} else {
			c.warn(stmt)
		}
	}
}

func (c *switchifChecker) warn(stmt ast.Stmt) {
	c.ctx.Warn(stmt, "should rewrite switch statement to if statement")
}

func (c *switchifChecker) warnDefault(stmt ast.Stmt) {
	c.ctx.Warn(stmt, "found switch with default case only")
}
