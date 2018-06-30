package lint

import "go/ast"

//! Finds where nesting level could be reduced.
//
// @Before:
// for _, v := range a {
//    if v.Bool {
//        ...
//    }
// }
//
// @After:
// for _, v := range a {
//    if ! v.Bool {
//        continue
//    }
//    ...
// }

func init() {
	addChecker(&nestingReduceChecker{}, attrExperimental)
}

type nestingReduceChecker struct {
	checkerBase
}

func (c *nestingReduceChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	c.checkBody(decl.Body.List)
	for _, stmt := range decl.Body.List {
		switch stmt := stmt.(type) {
		case *ast.RangeStmt:
			c.checkBody(stmt.Body.List)
		case *ast.ForStmt:
			c.checkBody(stmt.Body.List)
		case *ast.IfStmt:
			c.checkBody(stmt.Body.List)
		default:
		}
	}
}

func (c *nestingReduceChecker) checkBody(body []ast.Stmt) {
	if len(body) == 1 {
		if stmt, ok := body[0].(*ast.IfStmt); ok && c.checkIf(stmt) {
			c.warn(stmt)
		}
		return
	}
}

func (c *nestingReduceChecker) checkIf(stmt *ast.IfStmt) bool {
	const warnLen = 4
	if stmt.Init != nil {
		return false
	}
	if len(stmt.Body.List) > warnLen {
		return true
	}
	return false
}

func (c *nestingReduceChecker) warn(node ast.Node) {
	c.ctx.Warn(node, "nesting level could be reduced")
}
