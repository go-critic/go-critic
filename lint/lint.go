package lint

import "go/token"

// Flags ...
// TODO: Add description
type Flags struct {
}

// Context ...
// TODO: Add description
type Context struct {
	FileSet *token.FileSet
	Flags   *Flags
}

// Checker ...
// TODO: Add description
type Checker interface {
	Run(c *Context) error
}
