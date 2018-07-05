package lint

//! Detects when predeclared identifiers shadowed in assignments.
//
// @Before:
// func main() {
// 	// shadowing len function
// 	len := 10
// 	println(len)
// }
//
// @After:
// func main() {
// 	// change identificator name
// 	length := 10
// 	println(length)
// }

import (
	"go/ast"

	"github.com/go-critic/go-critic/lint/internal/astwalk"
)

func init() {
	addChecker(&builtinShadowChecker{}, attrSyntaxOnly)
}

type builtinShadowChecker struct {
	checkerBase

	builtins map[string]bool
}

func (c *builtinShadowChecker) Init() {
	c.builtins = map[string]bool{
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
	}
}

func (c *builtinShadowChecker) VisitLocalDef(name astwalk.Name, _ ast.Expr) {
	if _, isBuiltin := c.builtins[name.ID.String()]; isBuiltin {
		c.warn(name.ID)
	}
}

func (c *builtinShadowChecker) warn(ident *ast.Ident) {
	c.ctx.Warn(ident, "shadowing of predeclared identifier: %s", ident)
}
