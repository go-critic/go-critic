package lint

import (
	"go/ast"
)

// builtinShadowCheck detects when builtin functions shadowed in assignments
//
// Rationale: avoid bugs.
func builtinShadowCheck(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&builtinShadowChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},

		builtins: map[string]bool{
			"append":  true,
			"cap":     true,
			"close":   true,
			"complex": true,
			"copy":    true,
			"delete":  true,
			"imag":    true,
			"len":     true,
			"make":    true,
			"new":     true,
			"panic":   true,
			"print":   true,
			"println": true,
			"real":    true,
			"recover": true,
		},
	})
}

type builtinShadowChecker struct {
	baseStmtChecker

	builtins map[string]bool
}

func (c *builtinShadowChecker) CheckStmt(stmt ast.Stmt) {
	if assignStmt, ok := stmt.(*ast.AssignStmt); ok {
		for _, v := range assignStmt.Lhs {
			identificator := v.(*ast.Ident)
			if _, isBuiltin := c.builtins[identificator.Name]; isBuiltin {
				c.warn(identificator)
			}
		}
	}
}

func (c *builtinShadowChecker) warn(ident *ast.Ident) {
	c.ctx.Warn(ident, "assigning to builtin function: %s", ident.String())
}
