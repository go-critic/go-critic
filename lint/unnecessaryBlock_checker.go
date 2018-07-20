package lint

import (
	"go/ast"
)

func init() {
	addChecker(&unnecessaryBlockChecker{}, attrSyntaxOnly)
}

type unnecessaryBlockChecker struct {
	checkerBase
}

func (c *unnecessaryBlockChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects unnecessary braced statement blocks"
	d.Before = `
x := 1
{
print(x)
}`
	d.After = `
x := 1
print(x)`
}

func (c *unnecessaryBlockChecker) VisitStmtList(statements []ast.Stmt) {
	for _, stmt := range statements {
		if blockStmt, ok := stmt.(*ast.BlockStmt); ok {
			if c.hasAssignmentBlock(blockStmt) {
				continue
			}

			c.warn(blockStmt)
		}
	}
}

func (c *unnecessaryBlockChecker) hasAssignmentBlock(stmt *ast.BlockStmt) bool {
	for _, bs := range stmt.List {
		if _, ok := bs.(*ast.AssignStmt); ok {
			return true
		}
	}

	return false
}

func (c *unnecessaryBlockChecker) warn(expr ast.Stmt) {
	c.ctx.Warn(expr, "block doesn't have assignment statements, can be simply deleted")
}
