package lint

import (
	"go/ast"
)

func init() {
	addChecker(builtinShadowChecker{}, &ruleInfo{
		SyntaxOnly: true,
	})
}

type builtinShadowChecker struct {
	baseStmtChecker

	builtins map[string]bool
}

func (c builtinShadowChecker) New(ctx *context) func(*ast.File) {
	return wrapStmtChecker(&builtinShadowChecker{
		baseStmtChecker: baseStmtChecker{ctx: ctx},

		builtins: map[string]bool{
			// Types
			"bool":       true,
			"byte":       true,
			"complex64":  true,
			"complex128": true,
			"error":      true,
			"float32":    true,
			"float64":    true,
			"int":        true,
			"int8":       true,
			"int16":      true,
			"int32":      true,
			"int64":      true,
			"rune":       true,
			"string":     true,
			"uint":       true,
			"uint8":      true,
			"uint16":     true,
			"uint32":     true,
			"uint64":     true,
			"uintptr":    true,

			// Constants
			"true":  true,
			"false": true,
			"iota":  true,

			// Zero value
			"nil": true,

			// Functions
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

func (c *builtinShadowChecker) CheckStmt(stmt ast.Stmt) {
	if assignStmt, ok := stmt.(*ast.AssignStmt); ok {
		for _, v := range assignStmt.Lhs {
			ident, ok := v.(*ast.Ident)
			if !ok {
				continue
			}
			if _, isBuiltin := c.builtins[ident.Name]; isBuiltin {
				c.warn(ident)
			}
		}
	}
}

func (c *builtinShadowChecker) warn(ident *ast.Ident) {
	c.ctx.Warn(ident, "assigning to predeclared identifier: %s", ident)
}
