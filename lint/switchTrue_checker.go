package lint

//! Detects switch-over-bool statements that use explicit `true` tag value.
//
// @Before:
// switch true {
// case x > y:
// 	// ...
// }
//
// @After:
// switch {
// case x > y:
// 	// ...
// }

import "go/ast"

func init() {
	addChecker(&switchTrueChecker{}, attrSyntaxOnly)
}

type switchTrueChecker struct {
	checkerBase
}

func (c *switchTrueChecker) VisitStmt(stmt ast.Stmt) {
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
