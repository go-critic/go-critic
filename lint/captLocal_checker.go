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

func (c *captLocalChecker) EnterFunc(fn *ast.FuncDecl) bool {
	// Check func params eagerly.
	// We do this to avoid wasting clocks when checkLocals is false.
	if fn.Recv != nil && len(fn.Recv.List[0].Names) != 0 {
		c.checkDef(fn.Recv.List[0].Names[0])
	}
	for _, p := range fn.Type.Params.List {
		for _, id := range p.Names {
			c.checkDef(id)
		}
	}
	if fn.Type.Results != nil {
		for _, p := range fn.Type.Results.List {
			for _, id := range p.Names {
				c.checkDef(id)
			}
		}
	}

	return c.checkLocals && fn.Body != nil
}

func (c *captLocalChecker) VisitLocalDef(def astwalk.Name, _ ast.Expr) {
	if def.Kind == astwalk.NameParam {
		return // Already checked during EnterFunc
	}
	c.checkDef(def.ID)
}

func (c *captLocalChecker) checkDef(id *ast.Ident) {
	switch {
	case c.upcaseNames[id.Name]:
		c.warnUpcase(id)
	case ast.IsExported(id.Name):
		c.warnCapitalized(id)
	}
}

func (c *captLocalChecker) warnUpcase(id *ast.Ident) {
	c.ctx.Warn(id, "consider `%s' name instead of `%s'",
		strings.ToLower(id.Name), id)
}

func (c *captLocalChecker) warnCapitalized(id ast.Node) {
	c.ctx.Warn(id, "`%s' should not be capitalized", id)
}
