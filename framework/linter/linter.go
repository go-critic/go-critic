package linter

import (
	"go/token"
	"go/types"

	"github.com/go-critic/go-critic/linter"
)

// Deprecated: use linter.UnknownType.
var UnknownType = linter.UnknownType

// Deprecated: use linter.GoVersion.
type GoVersion = linter.GoVersion

// Deprecated: use linter.CheckerCollection.
type CheckerCollection = linter.CheckerCollection

// Deprecated: use linter.CheckerParam.
type CheckerParam = linter.CheckerParam

// Deprecated: use linter.CheckerParams.
type CheckerParams = linter.CheckerParams

// Deprecated: use linter.CheckerInfo.
type CheckerInfo = linter.CheckerInfo

// Deprecated: use linter.Checker.
type Checker = linter.Checker

// Deprecated: use linter.QuickFix.
type QuickFix = linter.QuickFix

// Deprecated: use linter.Warning.
type Warning = linter.Warning

// Deprecated: use linter.Context.
type Context = linter.Context

// Deprecated: use linter.CheckerContext.
type CheckerContext = linter.CheckerContext

// Deprecated: use linter.FileWalker.
type FileWalker = linter.FileWalker

// Deprecated: use linter.ParseGoVersion.
func ParseGoVersion(version string) (GoVersion, error) {
	return linter.ParseGoVersion(version)
}

// Deprecated: use linter.NewChecker.
func NewChecker(ctx *Context, info *CheckerInfo) (*Checker, error) {
	return linter.NewChecker(ctx, info)
}

// Deprecated: use linter.NewContext.
func NewContext(fset *token.FileSet, sizes types.Sizes) *Context {
	return linter.NewContext(fset, sizes)
}

// Deprecated: use linter.GetCheckersInfo.
func GetCheckersInfo() []*CheckerInfo {
	return linter.GetCheckersInfo()
}
