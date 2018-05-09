package lint

import "go/ast"

type UnderefChecker struct {
	ctx *Context

	warnings []Warning
}

func NewUnderedChecker(ctx *Context) {
	return &UnderefChecker{
		ctx: ctx,
	}
}

func (c *UnderefChecker) Check(f *ast.File) []Warning {

}
