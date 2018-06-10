package lint

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"sort"
)

type checkFunction interface {
	New(*context) func(*ast.File)
}

type ruleInfo struct {
	Experimental bool
	New          func(*context) func(*ast.File)
}

// checkFunctions is a table of all available check functions
// as well as their metadata (like "experimental" attributes).
//
// Keys are rule names.
var checkFunctions = map[string]*ruleInfo{}

// RuleList returns a slice of all rules that can be used to create checkers.
// Slice is sorted by rule names.
func RuleList() []*Rule {
	rules := make([]*Rule, 0, len(checkFunctions))
	for ruleName, info := range checkFunctions {
		rules = append(rules, &Rule{
			name:         ruleName,
			experimental: info.Experimental,
		})
	}
	sort.SliceStable(rules, func(i, j int) bool {
		return rules[i].Name() < rules[j].Name()
	})
	return rules
}

// Rule describes a named check that can be performed by the linter.
type Rule struct {
	name         string
	experimental bool
}

// String returns r short printed representation (name only).
func (r *Rule) String() string { return r.name }

// Name returns rule name.
func (r *Rule) Name() string { return r.name }

// Experimental reports whether rule is experimental.
//
// Rules are considered experimental when they have many
// false positives or some unresoleved bugs in them.
func (r *Rule) Experimental() bool { return r.experimental }

// NewChecker returns checker for the given rule.
//
// Rule must be non-nil and known to the lint package.
// Valid rules list can be obtained by RuleList call.
func NewChecker(rule *Rule, ctx *Context) *Checker {
	if rule == nil {
		panic("nil rule given")
	}
	c := &Checker{
		Rule: rule,
		ctx:  context{Context: ctx},
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
	Node ast.Node

	// Text is warning message without source location info.
	Text string
}

// Context is a readonly state shared among every checker.
type Context struct {
	// FileSet is a file set that was used during package parsing.
	FileSet *token.FileSet

	// TypesInfo carries parsed packages types information.
	TypesInfo *types.Info

	// SizesInfo carries alignment and type size information.
	// Arch-dependent.
	SizesInfo types.Sizes
}

// context is checker-local context copy.
// Fields that are not from Context itself are writeable.
type context struct {
	*Context
	warnings []Warning
}

// Warn adds a Warning to checker output.
func (ctx *context) Warn(node ast.Node, format string, args ...interface{}) {
	ctx.warnings = append(ctx.warnings, Warning{
		Text: fmt.Sprintf(format, args...),
		Node: node,
	})
}

func addChecker(c checkFunction, inf *ruleInfo) {
	typeName := reflect.ValueOf(c).Type().Name()
	ruleName := typeName[:len(typeName)-len("Checker")]
	inf.New = c.New
	checkFunctions[ruleName] = inf
}
