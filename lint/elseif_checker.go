package lint

import (
	"go/ast"
)

func init() {
	addChecker(elseifChecker{}, attrSyntaxOnly)
}

type elseifChecker struct {
	baseStmtChecker

	cause   *ast.IfStmt
	visited map[*ast.IfStmt]bool
}

func (c elseifChecker) New(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&elseifChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},
	})
}

func (c *elseifChecker) PerFuncInit(fn *ast.FuncDecl) bool {
	if fn.Body == nil {
		return false
	}
	c.visited = make(map[*ast.IfStmt]bool)
	return true
}

func (c *elseifChecker) CheckStmt(stmt ast.Stmt) {
	if stmt, ok := stmt.(*ast.IfStmt); ok {
		if c.visited[stmt] {
			return
		}
		c.cause = stmt
		c.checkIfStmt(stmt)
	}
}

func (c *elseifChecker) checkIfStmt(stmt *ast.IfStmt) {
	const minThreshold = 2
	if c.countIfelseLen(stmt) >= minThreshold {
		c.warn()
	}
}

func (c *elseifChecker) countIfelseLen(stmt *ast.IfStmt) int {
	count := 0
	for {
		switch e := stmt.Else.(type) {
		case *ast.IfStmt:
			// Else if.
			stmt = e
			count++
			c.visited[e] = true
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
	c.ctx.Warn(c.cause, "should rewrite if-else to switch statement")
}
