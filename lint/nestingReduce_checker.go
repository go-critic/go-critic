package lint

import "go/ast"

func init() {
	addChecker(&nestingReduceChecker{}, attrExperimental)
}

type nestingReduceChecker struct {
	checkerBase

	bodyWidth int
}

func (c *nestingReduceChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Finds where nesting level could be reduced"
	d.Before = `
for _, v := range a {
	if v.Bool {
		body()
	}
}`
	d.After = `
for _, v := range a {
	if !v.Bool {
		continue
	}
	body()
}`
}

func (c *nestingReduceChecker) Init() {
	c.bodyWidth = c.ctx.params.Int("bodyWidth", 5)
}

func (c *nestingReduceChecker) VisitStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ForStmt:
		c.checkLoopBody(stmt.Body.List)
	case *ast.RangeStmt:
		c.checkLoopBody(stmt.Body.List)
	}
}

func (c *nestingReduceChecker) checkLoopBody(body []ast.Stmt) {
	if len(body) != 1 {
		return
	}
	stmt, ok := body[0].(*ast.IfStmt)
	if !ok {
		return
	}
	if len(stmt.Body.List) >= c.bodyWidth && stmt.Else == nil {
		c.warnLoop(stmt)
	}
}

func (c *nestingReduceChecker) warnLoop(cause ast.Node) {
	c.ctx.Warn(cause, "invert if cond, replace body with `continue`, move old body after the statement")
}
