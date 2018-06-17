package lint

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/go-toolsmith/astfmt"
)

// checkerPrototypes is a table of checker prototypes that are
// used to instantiate new checker instances.
//
// Keys are rule names.
var checkerPrototypes = map[string]checkerProto{}

type checkerProto struct {
	rule *Rule

	// clone performs prototype copy and returns it as *Checker.
	clone func(context) *Checker
}

// abstractChecker is a proxy interface to forward checkerBase
// embedding checker into addChecker.
//
// It's an implementation detail that has only these guarantees:
//	- implemented by checkerBase type
//	- it is an argument type for addChecker
//
// Therefore, every checker that embeds checkerBase (normally, every checker)
// also a valid argument to addChecker.
type abstractChecker interface {
	BindContext(*context)
	Init()
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

// addChecker adds checker c to global checkers prototype table.
// Checker must be a pointer to zero value of concrete checker.
//
// Attributes used to fill AttributeSet for the rule inferred from checker.
func addChecker(c abstractChecker, attrs ...checkerAttribute) {
	// Clone abstractChecker underlying object.
	cloneAbstractChecker := func(c abstractChecker) abstractChecker {
		dynType := reflect.ValueOf(c).Elem().Type()
		return reflect.New(dynType).Interface().(abstractChecker)
	}

	bindCheckFunction := func(c abstractChecker) func(*ast.File) {
		// Infer proper AST traversing wrapper (walker).
		switch c := c.(type) {
		case funcDeclChecker:
			return wrapFuncDeclChecker(c)
		case exprChecker:
			return wrapExprChecker(c)
		case localExprChecker:
			return wrapLocalExprChecker(c)
		case stmtListChecker:
			return wrapStmtListChecker(c)
		case stmtChecker:
			return wrapStmtChecker(c)
		case localNameChecker:
			return wrapLocalNameChecker(c)
		case typeExprChecker:
			return wrapTypeExprChecker(c)
		default:
			panic(fmt.Sprintf("can't bind check function for %T", c))
		}
	}

	var rule Rule
	typeName := reflect.ValueOf(c).Type().String()
	rule.name = typeName[len("*lint.") : len(typeName)-len("Checker")]
	// Fill rule attributes using provided attr list.
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

	proto := checkerProto{rule: &rule}
	proto.clone = func(ctx context) *Checker {
		c := cloneAbstractChecker(c)
		clone := Checker{Rule: proto.rule}
		clone.check = bindCheckFunction(c)
		clone.ctx = ctx
		c.BindContext(&clone.ctx)
		c.Init()
		return &clone
	}
	checkerPrototypes[rule.name] = proto
}
