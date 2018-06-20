package lint

import (
	"go/ast"
)

func init() {
	addChecker(&singleCaseSwitchChecker{}, attrSyntaxOnly)
}

type singleCaseSwitchChecker struct {
	checkerBase
}

func (c *singleCaseSwitchChecker) VisitStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.SwitchStmt:
		c.checkSwitchStmt(stmt, stmt.Body)
	case *ast.TypeSwitchStmt:
		c.checkSwitchStmt(stmt, stmt.Body)
	}
}

func (c *singleCaseSwitchChecker) checkSwitchStmt(stmt ast.Stmt, body *ast.BlockStmt) {
	if len(body.List) == 1 {
		if body.List[0].(*ast.CaseClause).List == nil {
			// default case.
			c.warnDefault(stmt)
		} else if len(body.List[0].(*ast.CaseClause).List) == 1 {
			c.warn(stmt)
		}
	}
}

func (c *singleCaseSwitchChecker) warn(stmt ast.Stmt) {
	c.ctx.Warn(stmt, "should rewrite switch statement to if statement")
}

func (c *singleCaseSwitchChecker) warnDefault(stmt ast.Stmt) {
	c.ctx.Warn(stmt, "found switch with default case only")
}
