package lint

import (
	"go/ast"
)

// TypeGuardChecker finds type switches that may benefit from type guard clause.
//
// Rationale: code readability.
type TypeGuardChecker struct {
	ctx *Context
}

// NewTypeGuardChecker returns initialized checker for Go type switch statements.
func NewTypeGuardChecker(ctx *Context) *TypeGuardChecker {
	return &TypeGuardChecker{ctx: ctx}
}

// Check runs type switch checks for f.
func (c *TypeGuardChecker) Check(f *ast.File) []Warning {
	// TODO: implement me. Refs #10.
	return nil
}
