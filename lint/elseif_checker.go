package lint

import (
	"fmt"
	"go/ast"
)

// ElseifChecker finds repeated if-else statements and suggests to replace
// them with switch statement.
//
// Rationale: readability.
type ElseifChecker struct {
	ctx *Context

	rootStmt ast.Stmt

	warnings []Warning
}

// NewElseifChecker returns initilized checker for if statements.
func NewElseifChecker(ctx *Context) *ElseifChecker {
	return &ElseifChecker{ctx: ctx}
}

// Check runs if-else inspections for f.
//
// Features
//
// Permits single else or else-if; repeated else-if or else + else-if
// will trigger suggestion to use switch statement.
func (c *ElseifChecker) Check(f *ast.File) []Warning {
	c.warnings = c.warnings[:0]
	ast.Inspect(f, func(x ast.Node) bool {
		if stmt, ok := x.(*ast.IfStmt); ok {
			c.rootStmt = stmt
			return !c.checkIfStmt(stmt)
		}
		return true
	})
	return c.warnings
}

func (c *ElseifChecker) checkIfStmt(stmt *ast.IfStmt) bool {
	const minThreshold = 2
	if c.countIfelseLen(stmt) >= minThreshold {
		c.warn()
		return true
	}
	return false
}

func (c *ElseifChecker) countIfelseLen(stmt *ast.IfStmt) int {
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

func (c *ElseifChecker) warn() {
	c.warnings = append(c.warnings, Warning{
		Node: c.rootStmt,
		Text: fmt.Sprintf("should rewrite if-else to switch statement"),
	})
}
