package lint

import (
	"go/ast"
	"go/token"
	"go/types"
)

// WarkingKind describes checker warning category.
// Useful for checkers that can find different kinds of issues.
//
// Should be human-readable, not cryptic.
type WarningKind string

// Warning represents issue found by checker.
type Warning struct {
	Kind WarningKind
	Node ast.Node
	Text string
}

// Context ...
// TODO: Add description
type Context struct {
	// FileSet is a file set that was used during package parsing.
	FileSet *token.FileSet

	// TypesInfo carries parsed packages types information.
	TypesInfo *types.Info
}

// Checker analyzes given file for potential issues.
// Returns a list of linting errors.
//
// If checker encounters unexpected error, it should
// signal it using panic with argument of "error" type,
// but it should never call something like os.Exit or log.Fatal.
type Checker interface {
	Check(f *ast.File) []Warning
}
