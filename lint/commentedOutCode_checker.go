package lint

//! Detects commented-out code inside function bodies.
//
// @Before:
// // fmt.Println("Debugging hard")
// foo(1, 2)
//
// @After:
// foo(1, 2)

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/go-toolsmith/strparse"
)

func init() {
	addChecker(&commentedOutCodeChecker{}, attrExperimental)
}

type commentedOutCodeChecker struct {
	checkerBase
}

func (c *commentedOutCodeChecker) VisitLocalComment(cg *ast.CommentGroup) {
	s := cg.Text() // Collect text once

	// We do multiple heuristics to avoid false positives.
	// Many things can be improved here.

	markers := []string{
		"TODO", // TODO comments with code are permitted.

		// "http://" is interpreted as a label with comment.
		// There are other protocols we might want to include.
		"http://",
		"https://",

		"e.g. ", // Clearly not a "selector expr" (mostly due to extra space)
	}
	for _, m := range markers {
		if strings.Contains(s, m) {
			return
		}
	}

	// Some very short comment that can be skipped.
	// Usually triggering on these results in false positive.
	// Unless there is a very popular call like print/println.
	cond := len(s) < len("quite too short") &&
		!strings.Contains(s, "print") &&
		!strings.Contains(s, "fmt.") &&
		!strings.Contains(s, "log.")
	if cond {
		return
	}

	stmt := strparse.Stmt(s)
	if stmt == strparse.BadStmt {
		return // Most likely not a code
	}

	if !c.isPermittedStmt(stmt) {
		c.warn(cg)
	}
}

func (c *commentedOutCodeChecker) isPermittedStmt(stmt ast.Stmt) bool {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		return c.isPermittedExpr(stmt.X)
	case *ast.LabeledStmt:
		return c.isPermittedStmt(stmt.Stmt)
	case *ast.DeclStmt:
		decl := stmt.Decl.(*ast.GenDecl)
		return decl.Tok == token.TYPE
	default:
		return false
	}
}

func (c *commentedOutCodeChecker) isPermittedExpr(x ast.Expr) bool {
	// Permit anything except expressions that can be used
	// with complete result discarding.
	switch x := x.(type) {
	case *ast.CallExpr:
		return false
	case *ast.UnaryExpr:
		// "<-" channel receive is not permitted.
		return x.Op != token.ARROW
	default:
		return true
	}
}

func (c *commentedOutCodeChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "may want to remove commented-out code")
}
