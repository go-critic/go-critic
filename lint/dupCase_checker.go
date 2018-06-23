package lint

//! Detects duplicated case clauses inside switch statements.
//
// @Before:
// switch x {
// case ys[0], ys[1], ys[2], ys[0], ys[4]:
// }
//
// @After:
// switch x {
// case ys[0], ys[1], ys[2], ys[3], ys[4]:
// }

import (
	"go/ast"

	"github.com/go-toolsmith/astequal"
)

func init() {
	addChecker(&dupCaseChecker{}, attrExperimental)
}

type dupCaseChecker struct {
	checkerBase

	astSet []ast.Node
}

func (c *dupCaseChecker) VisitStmt(stmt ast.Stmt) {
	if stmt, ok := stmt.(*ast.SwitchStmt); ok {
		c.checkSwitch(stmt)
	}
}

func (c *dupCaseChecker) checkSwitch(stmt *ast.SwitchStmt) {
	c.astSet = c.astSet[:0]
	for i := range stmt.Body.List {
		cc := stmt.Body.List[i].(*ast.CaseClause)
		for _, x := range cc.List {
			if !c.setInsert(x) {
				c.warn(x)
			}
		}
	}
}

func (c *dupCaseChecker) setInsert(x ast.Node) bool {
	for i := range c.astSet {
		if astequal.Node(c.astSet[i], x) {
			return false
		}
	}
	c.astSet = append(c.astSet, x)
	return true
}

func (c *dupCaseChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "'case %s' is duplicated", cause)
}
