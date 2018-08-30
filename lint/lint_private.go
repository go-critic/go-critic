package lint

import (
	"fmt"
	"go/ast"
	"log"
	"reflect"
	"strings"

	"github.com/go-critic/go-critic/lint/internal/astwalk"
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
// abstractChecker is implemented by checkerBase directly and completely,
// making any checker that embeds it a valid argument to addChecker.
//
// See checkerBase and its implementation of this interface for more info.
type abstractChecker interface {
	BindContext(*context) // See checkerBase.BindContext
	Init()                // See checkerBase.Init

	// InitDocumentation fills Documentation object associated with checker.
	// Mandatory fields are Summary, Before and After.
	// See other checkers implementation for examples.
	InitDocumentation(*Documentation)
}

type checkerAttribute int

const (
	attrExperimental checkerAttribute = iota
	attrSyntaxOnly
	attrVeryOpinionated
)

type parameters map[string]interface{}

func (p parameters) Int(key string, defaultValue int) int {
	if value, ok := p[key]; ok {
		if value, ok := value.(int); ok {
			return value
		}
		log.Printf("incorrect value for `%s`, want int", key)
	}
	return defaultValue
}

func (p parameters) String(key, defaultValue string) string {
	if value, ok := p[key]; ok {
		if value, ok := value.(string); ok {
			return value
		}
		log.Printf("incorrect value for `%s`, want int", key)
	}
	return defaultValue
}

func (p parameters) Bool(key string, defaultValue bool) bool {
	if value, ok := p[key]; ok {
		if value, ok := value.(bool); ok {
			return value
		}
		log.Printf("incorrect value for `%s`, want bool", key)
	}
	return defaultValue
}

// context is checker-local context copy.
// Fields that are not from Context itself are writeable.
type context struct {
	*Context

	// printer used to format warning text.
	printer *astfmt.Printer

	params parameters

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

	trimDocs := func(d *Documentation) {
		fields := []*string{
			&d.Summary,
			&d.Details,
			&d.Before,
			&d.After,
			&d.Note,
		}
		for _, f := range fields {
			*f = strings.TrimSpace(*f)
		}
	}

	validateDocs := func(r *Rule) {
		// TODO(quasilyte): validate documentation.
	}

	newFileWalker := func(ctx *context, c abstractChecker) astwalk.FileWalker {
		// Infer proper AST traversing wrapper (walker).
		switch v := c.(type) {
		case astwalk.FileVisitor:
			return astwalk.WalkerForFile(v)
		case astwalk.FuncDeclVisitor:
			return astwalk.WalkerForFuncDecl(v)
		case astwalk.ExprVisitor:
			return astwalk.WalkerForExpr(v)
		case astwalk.LocalExprVisitor:
			return astwalk.WalkerForLocalExpr(v)
		case astwalk.StmtListVisitor:
			return astwalk.WalkerForStmtList(v)
		case astwalk.StmtVisitor:
			return astwalk.WalkerForStmt(v)
		case astwalk.LocalDefVisitor:
			return astwalk.WalkerForLocalDef(v, ctx.typesInfo)
		case astwalk.TypeExprVisitor:
			return astwalk.WalkerForTypeExpr(v, ctx.typesInfo)
		case astwalk.LocalCommentVisitor:
			return astwalk.WalkerForLocalComment(v)
		case astwalk.DocCommentVisitor:
			return astwalk.WalkerForDocComment(v)
		default:
			panic(fmt.Sprintf("%T does not implement known visitor interface", c))
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
	// Prepare associated documentation.
	c.InitDocumentation(&rule.Doc)
	trimDocs(&rule.Doc)
	validateDocs(&rule)

	proto := checkerProto{rule: &rule}
	proto.clone = func(ctx context) *Checker {
		c := cloneAbstractChecker(c)
		clone := &Checker{
			Rule: proto.rule,
			ctx:  ctx,
		}
		clone.walker = newFileWalker(&clone.ctx, c)
		c.BindContext(&clone.ctx)
		c.Init()
		return clone
	}
	checkerPrototypes[rule.name] = proto
}
