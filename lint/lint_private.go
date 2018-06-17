package lint

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/go-toolsmith/astfmt"
)

// checkerPrototypes is a table of existing checkers used
// instantiate new checkers.
//
// Keys are rule names.
var checkerPrototypes = map[string]Checker{}

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
	var rule Rule
	for _, attr := range attrs {
		switch attr {
		case attrExperimental:
			rule.Experimental = true
		case attrSyntaxOnly:
			rule.SyntaxOnly = true
		case attrVeryOpinionated:
			rule.VeryOpinionated = true
		default:
			panic(fmt.Sprintf("unexpected checkerAttribute"))
		}
	}
	typeName := reflect.ValueOf(c).Type().String()
	rule.name = typeName[len("*lint.") : len(typeName)-len("Checker")]
	checkerPrototypes[rule.name] = Checker{
		Rule: &rule,
		init: c.New,
	}
}
