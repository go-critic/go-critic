package lint

import (
	"go/ast"
	"go/token"

	"github.com/go-toolsmith/astequal"
)

func init() {
	addChecker(&nilValReturnChecker{}, attrExperimental)
}

type nilValReturnChecker struct {
	checkerBase
}

func (c *nilValReturnChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects return statements those results evaluate to nil"
	d.Before = `
if err == nil {
	return err
}`
	d.After = `
// (A) - return nil explicitly
if err == nil {
	return nil
}
// (B) - typo in "==", change to "!="
if err != nil {
	return nil
}`
}

func (c *nilValReturnChecker) VisitStmt(stmt ast.Stmt) {
	ifStmt, ok := stmt.(*ast.IfStmt)
	if !ok || len(ifStmt.Body.List) != 1 {
		return
	}
	ret, ok := ifStmt.Body.List[0].(*ast.ReturnStmt)
	if !ok || len(ret.Results) != 1 {
		return
	}
	expr, ok := ifStmt.Cond.(*ast.BinaryExpr)
	cond := ok &&
		expr.Op == token.EQL &&
		isSafeExpr(expr.X) &&
		qualifiedName(expr.Y) == "nil" &&
		astequal.Expr(expr.X, ret.Results[0])
	if cond {
		c.warn(ret, expr.X)
	}
}

func (c *nilValReturnChecker) warn(cause, val ast.Node) {
	c.ctx.Warn(cause, "returned expr is always nil; replace %s with nil", val)
}
