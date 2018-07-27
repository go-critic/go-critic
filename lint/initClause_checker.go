package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(&initClauseChecker{}, attrSyntaxOnly, attrExperimental)
}

type initClauseChecker struct {
	checkerBase
}

func (c *initClauseChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects non-assignment statements inside if/switch init clause"
	d.Before = `if sideEffect(); cond {
}`
	d.After = `sideEffect()
if cond {
}`
}

func (c *initClauseChecker) VisitStmt(stmt ast.Stmt) {
	initClause := c.getInitClause(stmt)
	if initClause != nil && !astp.IsAssignStmt(initClause) {
		c.warn(stmt, initClause)
	}
}

func (c *initClauseChecker) getInitClause(x ast.Stmt) ast.Stmt {
	switch x := x.(type) {
	case *ast.IfStmt:
		return x.Init
	case *ast.SwitchStmt:
		return x.Init
	default:
		return nil
	}
}

func (c *initClauseChecker) warn(stmt, clause ast.Stmt) {
	name := "if"
	if astp.IsSwitchStmt(stmt) {
		name = "switch"
	}
	c.ctx.Warn(stmt, "consider to move `%s` before %s", clause, name)
}
