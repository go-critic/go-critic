package lint

import (
	"go/ast"

	"github.com/go-toolsmith/astp"
)

func init() {
	// Opinionated because it does give questionable advices for cases
	// where else with nested if is used for readability with preceding if body.
	addChecker(&elseifChecker{}, attrExperimental, attrVeryOpinionated)
}

type elseifChecker struct {
	checkerBase

	skipBalanced bool
}

func (c *elseifChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects else with nested if statement that can be replaced with else-if"
	d.Before = `
if cond1 {
} else {
	if x := cond2; x {
	}
}`
	d.After = `
if cond1 {
} else if x := cond2; x {
}`
}

func (c *elseifChecker) Init() {
	c.skipBalanced = c.ctx.params.Bool("skipBalanced", true)
}

func (c *elseifChecker) VisitStmt(stmt ast.Stmt) {
	if stmt, ok := stmt.(*ast.IfStmt); ok {
		elseBody, ok := stmt.Else.(*ast.BlockStmt)
		if !ok || len(elseBody.List) != 1 {
			return
		}
		innerIfStmt, ok := elseBody.List[0].(*ast.IfStmt)
		if !ok {
			return
		}
		balanced := len(stmt.Body.List) == 1 &&
			astp.IsIfStmt(stmt.Body.List[0])
		if balanced && c.skipBalanced {
			return // Configured to skip balanced statements
		}
		if innerIfStmt.Else != nil {
			return
		}
		c.warn(stmt.Else)
	}
}

func (c *elseifChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "can replace 'else {if cond {}}' with 'else if cond {}'")
}
