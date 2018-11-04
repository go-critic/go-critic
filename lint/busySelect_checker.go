package lint

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(&busySelectChecker{}, attrExperimental)
}

type busySelectChecker struct {
	checkerBase
}

func (c *busySelectChecker) InitDocumentation(d *Documentation) {
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

func (c *busySelectChecker) VisitStmt(stmt ast.Stmt) {
	forStmt, ok := stmt.(*ast.ForStmt)
	if !ok || forStmt.Cond != nil {
		return
	}

	selectStmt, ok := findNode(forStmt, astp.IsSelectStmt).(*ast.SelectStmt)
	if !ok {
		return
	}

	for _, s := range selectStmt.Body.List {
		s := s.(*ast.CommClause)
		if s.Comm == nil {
			if !c.hasBlockingStmt(s.Body) && !c.hasBlockingStmt(forStmt.Body.List) {
				c.warn(s)
			}
		}
	}
}

func (c *busySelectChecker) hasBlockingStmt(stmts []ast.Stmt) bool {
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
				return qualifiedName(expr.Fun) == "time.Sleep"
			}
		}
	}
	return false
}

func (c *busySelectChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "default case without a blocking operation or sleep might waste a CPU time")
}
