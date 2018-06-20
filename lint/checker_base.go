package lint

import (
	"go/ast"
)

// checkerBase is a type to be embedded into every checker type.
type checkerBase struct {
	ctx *context
}

// BindContext saves checker-local context.
// Called once before Init.
//
// Generally, embedding checker needs not to define BindContext
// as default implementation does the right thing.
func (c *checkerBase) BindContext(ctx *context) {
	c.ctx = ctx
}

// Init implements checker initialization.
// It is called for zero value instance only once, inside NewChecker.
//
// Default initialization does nothing.
// If embedding checker has a state that needs to be initialized, it must
// define method with the same signature and perform initialization there.
func (c *checkerBase) Init() {}

// EnterFunc makes checker not enter external functions.
// External functions are functions without body.
func (c *checkerBase) EnterFunc(fn *ast.FuncDecl) bool {
	return fn.Body != nil
}

// EnterChilds makes checker unconditionally traverse every node siblings.
func (c *checkerBase) EnterChilds(x ast.Node) bool {
	return true
}
