package lint

import "go/token"

// Context ...
// TODO: Add description
type Context struct {
	// FileSet is a file set that was used during package parsing.
	FileSet *token.FileSet

	// PkgDir is a path to package being checked.
	PkgDir string
}

// Checker ...
// TODO: Add description
type Checker interface {
	Run(ctx *Context) error
}
