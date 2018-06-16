package lint

import "go/ast"

func init() {
	addChecker(switchTrueChecker{}, &ruleInfo{
		SyntaxOnly: true,
	})
}

type switchTrueChecker struct {
	baseStmtChecker
}

func (c switchTrueChecker) New(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&switchTrueChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

func (c *switchTrueChecker) CheckStmt(stmt ast.Stmt) {
	if stmt, ok := stmt.(*ast.SwitchStmt); ok {
		if qualifiedName(stmt.Tag) == "true" {
			c.warn(stmt)
		}
	}
}

func (c *switchTrueChecker) warn(cause *ast.SwitchStmt) {
	if cause.Init == nil {
		c.ctx.Warn(cause, "replace 'switch true {}' with 'switch {}'")
	} else {
		c.ctx.Warn(cause, "replace 'switch %s; true {}' with 'switch %s; {}'",
			cause.Init, cause.Init)
	}
}
