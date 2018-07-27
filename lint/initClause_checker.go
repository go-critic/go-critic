package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(&initClauseChecker{}, attrSyntaxOnly)
}

type initClauseChecker struct {
	checkerBase
}

func (c *initClauseChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects non-assignment statements inside if/switch init clause."
	d.Before = `123`
	d.After = `123s`
}

func (c *initClauseChecker) VisitStmt(stmt ast.Stmt) {
	if expr, ok := c.getInitClause(stmt).(*ast.ExprStmt); ok {
		if astp.IsCallExpr(expr.X) {
			c.warn(expr, stmt)
		}
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

func (c *initClauseChecker) warn(expr *ast.ExprStmt, stmt ast.Stmt) {
	name := "if"
	if astp.IsSwitchStmt(stmt) {
		name = "switch"
	}
	c.ctx.Warn(expr, "consider to move `%s` before %s", expr, name)
}
