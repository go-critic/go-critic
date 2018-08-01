package lint

import (
	"go/ast"
	"go/token"
)

func init() {
	addChecker(&spinningDefaultChecker{}, attrExperimental)
}

type spinningDefaultChecker struct {
	checkerBase
}

func (c *spinningDefaultChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects default statement inside a select without a sleep that might waste a CPU time"
	d.Before = `for {
	select {
		case <-ch:
			// ...
		default:
			// will waste CPU time
	}
}`
	d.After = `for {
	select {
		case <-ch:
			// ...
		default:
			time.Sleep(100 * time.Millisecond)
	}
}`
}

func (c *spinningDefaultChecker) VisitStmt(stmt ast.Stmt) {
	forStmt, ok := stmt.(*ast.ForStmt)
	if !ok || forStmt.Cond != nil {
		return
	}

	selectStmt := c.getSelectStmt(forStmt)
	if selectStmt == nil {
		return
	}

	for _, s := range selectStmt.Body.List {
		s := s.(*ast.CommClause)
		if s.Comm == nil {
			if !c.hasBlockingStmt(s.Body) {
				c.warn(s)
			}
		}
	}
}

func (c *spinningDefaultChecker) getSelectStmt(x ast.Stmt) *ast.SelectStmt {
	switch x := x.(type) {
	case *ast.RangeStmt:
		for _, s := range x.Body.List {
			if s, ok := s.(*ast.SelectStmt); ok {
				return s
			}
		}
	case *ast.ForStmt:
		for _, s := range x.Body.List {
			if s, ok := s.(*ast.SelectStmt); ok {
				return s
			}
		}
	}
	return nil
}

func (c *spinningDefaultChecker) hasBlockingStmt(stmts []ast.Stmt) bool {
	for _, s := range stmts {
		switch s := s.(type) {
		case *ast.SendStmt:
			// ch <- ...
			return true

		case *ast.ExprStmt:
			switch expr := s.X.(type) {
			case *ast.UnaryExpr:
				// <- ch
				// <- time.After(...)
				// <- time.Tick(...)
				return expr.Op == token.ARROW

			case *ast.CallExpr:
				// time.Sleep(...)
				call, ok := expr.Fun.(*ast.SelectorExpr)
				return ok && call.X.(*ast.Ident).Name == "time" && call.Sel.Name == "Sleep"
			}
		}
	}
	return false
}

func (c *spinningDefaultChecker) warn(node ast.Node) {
	c.ctx.Warn(node, "default case without a blocking operation or sleep might waste a CPU time")
}
