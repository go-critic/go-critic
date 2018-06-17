package lint

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/go-toolsmith/astfmt"
)

// checkFunctions is a table of all available check functions
// as well as their metadata (like "experimental" attributes).
//
// Keys are rule names.
var checkFunctions = map[string]*ruleInfo{}

type ruleInfo struct {
	AttributeSet

	New func(*context) func(*ast.File)
}

type checkFunction interface {
	New(*context) func(*ast.File)
}

type checkerAttribute int

const (
	attrExperimental checkerAttribute = iota
	attrSyntaxOnly
	attrVeryOpinionated
)

// context is checker-local context copy.
// Fields that are not from Context itself are writeable.
type context struct {
	*Context

	// printer used to format warning text.
	printer *astfmt.Printer

	warnings []Warning
}

// Warn adds a Warning to checker output.
func (ctx *context) Warn(node ast.Node, format string, args ...interface{}) {
	ctx.warnings = append(ctx.warnings, Warning{
		Text: ctx.printer.Sprintf(format, args...),
		Node: node,
	})
}

func addChecker(c checkFunction, attrs ...checkerAttribute) {
	var info ruleInfo
	for _, attr := range attrs {
		switch attr {
		case attrExperimental:
			info.Experimental = true
		case attrSyntaxOnly:
			info.SyntaxOnly = true
		case attrVeryOpinionated:
			info.VeryOpinionated = true
		default:
			panic(fmt.Sprintf("unexpected checkerAttribute"))
		}
	}
	typeName := reflect.ValueOf(c).Type().String()
	ruleName := typeName[len("*lint.") : len(typeName)-len("Checker")]
	info.New = c.New
	checkFunctions[ruleName] = &info
}
