package lint

import (
	"go/ast"
	"strings"

	"github.com/go-critic/go-critic/lint/internal/astwalk"
)

func init() {
	addChecker(&captLocalChecker{}, attrSyntaxOnly)
}

type captLocalChecker struct {
	checkerBase

	upcaseNames map[string]bool
}

func (c *captLocalChecker) InitDocs(d *Documentation) {
	d.Summary = "Detects capitalized names for local variables"
	d.Before = `func f(IN int, OUT *int) (ERR error) {}`
	d.After = `func f(in int, out *int) (err error) {}`
}

func (c *captLocalChecker) Init() {
	c.upcaseNames = map[string]bool{
		"IN":    true,
		"OUT":   true,
		"INOUT": true,

		// TODO: add common acronyms like HTTP and URL?
	}
}

func (c *captLocalChecker) VisitLocalDef(def astwalk.Name, _ ast.Expr) {
	switch {
	case c.upcaseNames[def.ID.Name]:
		c.warnUpcase(def.ID)
	case ast.IsExported(def.ID.Name):
		c.warnCapitalized(def.ID)
	}
}

func (c *captLocalChecker) warnUpcase(id *ast.Ident) {
	c.ctx.Warn(id, "consider `%s' name instead of `%s'",
		strings.ToLower(id.Name), id)
}

func (c *captLocalChecker) warnCapitalized(id ast.Node) {
	c.ctx.Warn(id, "`%s' should not be capitalized", id)
}
