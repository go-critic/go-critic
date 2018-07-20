package lint

import (
	"go/ast"

	"github.com/go-critic/go-critic/lint/internal/lintutil"
)

func init() {
	addChecker(&dupCaseChecker{}, attrExperimental)
}

type dupCaseChecker struct {
	checkerBase

	astSet lintutil.AstMap
}

func (c *dupCaseChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects duplicated case clauses inside switch statements"
	d.Before = `
switch x {
case ys[0], ys[1], ys[2], ys[0], ys[4]:
}`
	d.After = `
switch x {
case ys[0], ys[1], ys[2], ys[3], ys[4]:
}`
}

func (c *dupCaseChecker) VisitStmt(stmt ast.Stmt) {
	if stmt, ok := stmt.(*ast.SwitchStmt); ok {
		c.checkSwitch(stmt)
	}
}

func (c *dupCaseChecker) checkSwitch(stmt *ast.SwitchStmt) {
	c.astSet.Clear()
	for i := range stmt.Body.List {
		cc := stmt.Body.List[i].(*ast.CaseClause)
		for _, x := range cc.List {
			if !c.astSet.Insert(x, nil) {
				c.warn(x)
			}
		}
	}
}

func (c *dupCaseChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "'case %s' is duplicated", cause)
}
