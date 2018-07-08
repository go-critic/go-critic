package lint

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"sort"

	"github.com/go-critic/go-critic/lint/internal/astwalk"
	"github.com/go-toolsmith/astfmt"
)

// RuleList returns a slice of all rules that can be used to create checkers.
// Slice is sorted by rule names.
func RuleList() []*Rule {
	rules := make([]*Rule, 0, len(checkerPrototypes))
	for _, c := range checkerPrototypes {
		rules = append(rules, c.rule)
	}
	sort.SliceStable(rules, func(i, j int) bool {
		return rules[i].Name() < rules[j].Name()
	})
	return rules
}

// AttributeSet describes rule implementation properties that may be
// related to implementation (experimental) or be more fundamental (opinionated).
type AttributeSet struct {
	// SyntaxOnly marks rules that can be checker by using AST only.
	// Such rule implementations should not use types info.
	//
	// It's OK not to mark rule as SyntaxOnly even if current
	// implementation does not use types info.
	SyntaxOnly bool

	// Experimental marks rule implementation as experimental.
	// Reasons to mark anything as experimental:
	//	- rule name can change in near time
	//	- checker gives more false positives than anticipated
	//	- rule requires more testing or alternative design
	Experimental bool

	// VeryOpinionated marks rule as controversial for some audience and
	// that it might be not suitable for everyone.
	VeryOpinionated bool
}

// Rule describes a named check that can be performed by the linter.
type Rule struct {
	AttributeSet
	Doc Documentation

	name string
}

// String returns r short printed representation (name only).
func (r *Rule) String() string { return r.name }

// Name returns rule name.
func (r *Rule) Name() string { return r.name }

// NewChecker returns checker for the given rule.
//
// Rule must be non-nil and known to the lint package.
// Valid rule list can be obtained by RuleList call.
func NewChecker(rule *Rule, ctx *Context) *Checker {
	if rule == nil {
		panic("nil rule given")
	}
	c, ok := checkerPrototypes[rule.Name()]
	if !ok {
		panic(fmt.Sprintf("rule %q is undefined", rule.Name()))
	}
	return c.clone(context{
		Context: ctx,
		printer: astfmt.NewPrinter(ctx.fileSet),
	})
}

// Checker analyzes given file for potential issues.
// The checks that are being run depends on the associated rule.
type Checker struct {
	Rule *Rule

	ctx context

	walker astwalk.FileWalker
}

// Check runs rule checker over file f.
func (c *Checker) Check(f *ast.File) []Warning {
	c.ctx.warnings = c.ctx.warnings[:0]
	c.walker.WalkFile(f)
	return c.ctx.warnings
}

// Warning represents issue that is found by rule checker.
type Warning struct {
	// Node is an AST node that caused warning to trigger.
	// Can be used to obtain proper error location.
	Node ast.Node

	// Text is warning message without source location info.
	Text string
}

// Context is a readonly state shared among every checker.
type Context struct {
	// filename is a currently checked file name.
	filename string

	// fileSet is a file set that was used during package parsing.
	fileSet *token.FileSet

	// pkg describes package that is being checked.
	pkg *types.Package

	// typesInfo carries parsed packages types information.
	typesInfo *types.Info

	// sizesInfo carries alignment and type size information.
	// Arch-dependent.
	sizesInfo types.Sizes
}

// NewContext returns new shared context to be used by every checker.
//
// All data carried by the context is readonly for checkers,
// but can be modified by integrating application.
func NewContext(fset *token.FileSet, sizes types.Sizes) *Context {
	return &Context{
		fileSet:   fset,
		sizesInfo: sizes,
		typesInfo: &types.Info{},
	}
}

// FileSet returns file set used upon context creation.
func (c *Context) FileSet() *token.FileSet { return c.fileSet }

// SetPackageInfo sets package-related metadata.
//
// Must be called for every package being checked.
//
// Type information may be nil if all requested checkers
// have SyntaxOnly attribute.
func (c *Context) SetPackageInfo(info *types.Info, pkg *types.Package) {
	if info != nil {
		// We do this kind of assignment to avoid
		// changing c.typesInfo field address after
		// every re-assignment.
		*c.typesInfo = *info
	}
	c.pkg = pkg
}

// SetFileInfo sets file-related metadata.
//
// Must be called for every source code file being checked.
func (c *Context) SetFileInfo(filename string) {
	c.filename = filename
}

// Documentation holds rule structured documentation.
type Documentation struct {
	// Summary is a short one sentence description.
	// Should not end with a period.
	Summary string
	// Details extends summary with additional info. Optional.
	Details string
	// Before is a code snippet of code that will violate rule.
	Before string
	// After is a code snippet of fixed code that complies to the rule.
	After string
	// Note is an optional caution message or advice.
	Note string
}
