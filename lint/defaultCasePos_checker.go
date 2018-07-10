package lint

//! Detects when default case in switch isn't on 1st or last position.
//
// @Before:
// switch {
// case x > y:
// 	// ...
// default: // <- not the best position
//	// ...
// case x == 10:
//	// ...
// }
//
// @After:
// switch {
// case x > y:
// 	// ...
// case x == 10:
//	// ...
// default: // <- everything is good
//	// ...
// }

import "go/ast"

func init() {
	addChecker(&defaultCasePosChecker{}, attrSyntaxOnly)
}

type defaultCasePosChecker struct {
	checkerBase
}

func (c *defaultCasePosChecker) VisitStmt(stmt ast.Stmt) {
	swtch, ok := stmt.(*ast.SwitchStmt)
	if !ok || swtch.Body == nil || swtch.Body.List == nil {
		return
	}
	for i, stmt := range swtch.Body.List {
		caseStmt, ok := stmt.(*ast.CaseClause)
		if !ok {
			continue
		}
		if caseStmt.List == nil && (i != 0 && i != len(swtch.Body.List)-1) {
			c.warn(caseStmt)
		}
	}
}

func (c *defaultCasePosChecker) warn(cause *ast.CaseClause) {
	c.ctx.Warn(cause, "consider to make `default` case as first or as last case")
}
