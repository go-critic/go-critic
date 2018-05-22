package lint

import (
	"errors"
	"go/ast"
	"go/token"
	"go/types"
	"sort"
	"sync"
)

// WarningKind describes checker warning category.
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

	// SizesInfo carries alignment and type size information.
	// Arch-dependent.
	SizesInfo types.Sizes

	mutex sync.Mutex // Protects the rest of the fields

	// Warnings contains warnings from all checkers
	Warnings []Warning
}

func (c *Context) addWarning(w Warning) {
	c.mutex.Lock()
	c.Warnings = append(c.Warnings, w)
	c.mutex.Unlock()
}

// Checker analyzes given file for potential issues.
// Returns a list of linting errors.
//
// If checker encounters unexpected error, it should
// signal it using panic with argument of "error" type,
// but it should never call something like os.Exit or log.Fatal.
type Checker interface {
	Check(f *ast.File)
}

var checkers = map[string]func(c *Context) Checker{
	"param-name":        newParamNameChecker,
	"type-guard":        newTypeGuardChecker,
	"parenthesis":       newParenthesisChecker,
	"underef":           newUnderefChecker,
	"param-duplication": newParamDuplicationChecker,
	"elseif":            newElseifChecker,
	"big-copy":          newBigCopyChecker,
	"long-chain":        newLongChainChecker,
	"switchif":          newSwitchifChecker,
}

// NewChecker returns checker that implements check of specified name.
func NewChecker(name string, ctx *Context) (Checker, error) {
	c, ok := checkers[name]
	if !ok {
		return nil, errors.New("checker not found")
	}
	return c(ctx), nil
}

// AvailableCheckers returns all check names that are supported.
// Strings returned from this function can be safely used in NewChecker call.
func AvailableCheckers() []string {
	var names []string
	for name := range checkers {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
