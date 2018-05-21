package lint

import (
	"fmt"
	"go/ast"
)

type switchifChecker struct {
	ctx *Context
}

func newSwitchifChecker(ctx *Context) Checker {
	return &switchifChecker{ctx: ctx}
}

func (c *switchifChecker) Check(f *ast.File) {
	ast.Inspect(f, func(x ast.Node) bool {
		switch stmt := x.(type) {
		case *ast.SwitchStmt:
			c.checkSwitchStmt(stmt, stmt.Body)
		case *ast.TypeSwitchStmt:
			c.checkSwitchStmt(stmt, stmt.Body)
		}
		return true
	})
}

func (c *switchifChecker) checkSwitchStmt(stmt ast.Stmt, body *ast.BlockStmt) {
	if len(body.List) == 1 {
		if body.List[0].(*ast.CaseClause).List == nil {
			// default case
			c.warnDefault(stmt)
		} else {
			c.warn(stmt)
		}
	}
}

func (c *switchifChecker) warn(stmt ast.Stmt) {
	c.ctx.addWarning(Warning{
		Kind: "switchif",
		Node: stmt,
		Text: fmt.Sprintf("should rewrite switch statement to if statement"),
	})
}

func (c *switchifChecker) warnDefault(stmt ast.Stmt) {
	c.ctx.addWarning(Warning{
		Kind: "switchif",
		Node: stmt,
		Text: fmt.Sprintf("found switch with default case only"),
	})
}
