package lint

import (
	"go/ast"
	"strings"

	"github.com/go-critic/go-critic/lint/internal/astwalk"
)

func init() {
	addChecker(&captLocalChecker{}, attrExperimental, attrSyntaxOnly)
}

type captLocalChecker struct {
	checkerBase

	upcaseNames map[string]bool
	checkLocals bool
}

func (c *captLocalChecker) InitDocumentation(d *Documentation) {
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

	c.checkLocals = c.ctx.params.Bool("checkLocals", false)
}

func (c *captLocalChecker) VisitLocalDef(def astwalk.Name, _ ast.Expr) {
	if !c.checkLocals && def.Kind != astwalk.NameParam {
		return
	}

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
