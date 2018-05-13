package lint

import (
	"fmt"
	"go/ast"
)

// BuiltinShadows detects when builtin functions shadowed in assignments
type BuiltinShadows struct {
	ctx *Context

	warnings []Warning
}

func newBuiltinShadowsChecker(ctx *Context) Checker {
	return &BuiltinShadows{
		ctx: ctx,
	}
}

// Check runs builtin functions shadow check on assignments for f
func (c *BuiltinShadows) Check(f *ast.File) []Warning {
	c.warnings = c.warnings[:0]

	builtinFuncs := map[string]struct{}{
		"append":  struct{}{},
		"cap":     struct{}{},
		"close":   struct{}{},
		"complex": struct{}{},
		"copy":    struct{}{},
		"delete":  struct{}{},
		"imag":    struct{}{},
		"len":     struct{}{},
		"make":    struct{}{},
		"new":     struct{}{},
		"panic":   struct{}{},
		"print":   struct{}{},
		"println": struct{}{},
		"real":    struct{}{},
		"recover": struct{}{},
	}

	ast.Inspect(f, func(x ast.Node) bool {
		if stmt, ok := x.(*ast.AssignStmt); ok {
			for _, v := range stmt.Lhs {
				identificator := v.(*ast.Ident)
				if _, isBuiltin := builtinFuncs[identificator.Name]; isBuiltin {
					c.warn(identificator)
					return false
				}
			}
		}
		return true
	})

	return c.warnings
}

func (c *BuiltinShadows) warn(x *ast.Ident) {
	c.warnings = append(c.warnings, Warning{
		Kind: "Shadowing",
		Node: x,
		Text: fmt.Sprintf("%s shadowing", x.Name),
	})
}
