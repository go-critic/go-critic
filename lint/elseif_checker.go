package lint

import (
	"fmt"
	"go/ast"
)

type elseifChecker struct {
	ctx *Context

	rootStmt ast.Stmt
}

func newElseifChecker(ctx *Context) Checker {
	return &elseifChecker{ctx: ctx}
}

// Check finds repeated if-else statements and suggests to replace
// them with switch statement.
//
// Features
//
// Permits single else or else-if; repeated else-if or else + else-if
// will trigger suggestion to use switch statement.
func (c *elseifChecker) Check(f *ast.File) {
	ast.Inspect(f, func(x ast.Node) bool {
		if stmt, ok := x.(*ast.IfStmt); ok {
			c.rootStmt = stmt
			return !c.checkIfStmt(stmt)
		}
		return true
	})
}

func (c *elseifChecker) checkIfStmt(stmt *ast.IfStmt) bool {
	const minThreshold = 2
	if c.countIfelseLen(stmt) >= minThreshold {
		c.warn()
		return true
	}
	return false
}

func (c *elseifChecker) countIfelseLen(stmt *ast.IfStmt) int {
	count := 0
	for {
		switch e := stmt.Else.(type) {
		case *ast.IfStmt:
			// Else if.
			stmt = e
			count++
		case *ast.BlockStmt:
			// Else branch.
			return count + 1
		default:
			// No else or else if.
			return count
		}
	}
}

func (c *elseifChecker) warn() {
	c.ctx.addWarning(Warning{
		Kind: "elseif",
		Node: c.rootStmt,
		Text: fmt.Sprintf("should rewrite if-else to switch statement"),
	})
}
