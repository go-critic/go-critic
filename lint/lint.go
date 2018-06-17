package lint

import (
	"go/ast"
	"go/token"
	"go/types"
	"sort"

	"github.com/go-toolsmith/astfmt"
)

// RuleList returns a slice of all rules that can be used to create checkers.
// Slice is sorted by rule names.
func RuleList() []*Rule {
	rules := make([]*Rule, 0, len(checkFunctions))
	for ruleName, info := range checkFunctions {
		rules = append(rules, &Rule{
			name:         ruleName,
			AttributeSet: info.AttributeSet,
		})
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

	name string
}

// String returns r short printed representation (name only).
func (r *Rule) String() string { return r.name }

// Name returns rule name.
func (r *Rule) Name() string { return r.name }

// NewChecker returns checker for the given rule.
//
// Rule must be non-nil and known to the lint package.
// Valid rules list can be obtained by RuleList call.
func NewChecker(rule *Rule, ctx *Context) *Checker {
	if rule == nil {
		panic("nil rule given")
	}
	// TODO(quasilyte): it would be great to have SyntaxOnly
	// checkers TypeInfo set to nil, so if they ever
	// start using it that will be detected by the tests.
	// This can require some lint package refactoring though.
	// Postponing this idea for now.
	c := &Checker{
		Rule: rule,
		ctx: context{
			Context: ctx,
			printer: astfmt.NewPrinter(ctx.FileSet),
		},
	}
	c.check = checkFunctions[rule.Name()].New(&c.ctx)
	return c
}

// Checker analyzes given file for potential issues.
// The checks that are being run depends on the associated rule.
type Checker struct {
	Rule *Rule

	ctx context

	check func(*ast.File)
}

// Check runs rule checker over file f.
func (c *Checker) Check(f *ast.File) []Warning {
	c.ctx.warnings = c.ctx.warnings[:0]
	c.check(f)
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
	Filename string

	// FileSet is a file set that was used during package parsing.
	FileSet *token.FileSet

	// TypesInfo carries parsed packages types information.
	TypesInfo *types.Info

	// Package describes package that is being checked.
	Package *types.Package

	// SizesInfo carries alignment and type size information.
	// Arch-dependent.
	SizesInfo types.Sizes
}
