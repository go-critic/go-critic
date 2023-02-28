package linter

import (
	"go/token"
	"go/types"

	"github.com/go-critic/go-critic/linter"
)

// UnknownType is a special sentinel value that is returned from the CheckerContext.TypeOf
// method instead of the nil type.
// Deprecated: use linter.UnknownType.
var UnknownType = linter.UnknownType

// Deprecated: use linter.GoVersion.
type GoVersion = linter.GoVersion

// CheckerCollection provides additional information for a group of checkers.
// Deprecated: use linter.CheckerCollection.
type CheckerCollection = linter.CheckerCollection

// CheckerParam describes a single checker customizable parameter.
// Deprecated: use linter.CheckerParam.
type CheckerParam = linter.CheckerParam

// CheckerParams holds all checker-specific parameters.
//
// Provides convenient access to the loosely typed underlying map.
// Deprecated: use linter.CheckerParams.
type CheckerParams = linter.CheckerParams

// CheckerInfo holds checker metadata and structured documentation.
// Deprecated: use linter.CheckerInfo.
type CheckerInfo = linter.CheckerInfo

// Checker is an implementation of a check that is described by the associated info.
// Deprecated: use linter.Checker.
type Checker = linter.Checker

// QuickFix is our analysis.TextEdit; we're using it here to avoid
// direct analysis package dependency for now.
// Deprecated: use linter.QuickFix.
type QuickFix = linter.QuickFix

// Warning represents issue that is found by checker.
// Deprecated: use linter.Warning.
type Warning = linter.Warning

// Context is a readonly state shared among every checker.
// Deprecated: use linter.Context.
type Context = linter.Context

// CheckerContext is checker-local context copy.
// Fields that are not from Context itself are writeable.
// Deprecated: use linter.CheckerContext.
type CheckerContext = linter.CheckerContext

// FileWalker is an interface every checker should implement.
//
// The WalkFile method is executed for every Go file inside the
// package that is being checked.
// Deprecated: use linter.FileWalker.
type FileWalker = linter.FileWalker

// Deprecated: use linter.ParseGoVersion.
func ParseGoVersion(version string) (GoVersion, error) {
	return linter.ParseGoVersion(version)
}

// NewChecker returns initialized checker identified by an info.
// info must be non-nil.
// Returns an error if info describes a checker that was not properly registered,
// or if checker fails to initialize.
// Deprecated: use linter.NewChecker.
func NewChecker(ctx *Context, info *CheckerInfo) (*Checker, error) {
	return linter.NewChecker(ctx, info)
}

// NewContext returns new shared context to be used by every checker.
//
// All data carried by the context is readonly for checkers,
// but can be modified by the integrating application.
// Deprecated: use linter.NewContext.
func NewContext(fset *token.FileSet, sizes types.Sizes) *Context {
	return linter.NewContext(fset, sizes)
}

// GetCheckersInfo returns a checkers info list for all registered checkers.
// The slice is sorted by a checker name.
//
// Info objects can be used to instantiate checkers with NewChecker function.
// Deprecated: use linter.GetCheckersInfo.
func GetCheckersInfo() []*CheckerInfo {
	return linter.GetCheckersInfo()
}
