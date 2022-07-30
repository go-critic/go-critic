package checkers

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"strings"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/framework/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "protoGetter"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects reading fields of proto message structs without the getter"
	info.Before = `
t.Data.Value
`
	info.After = `
t.GetData().GetValue()
`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &protoGetterChecker{ctx: ctx, ignores: make(map[token.Pos]struct{})}
		return c, nil
	})
}

type protoGetterChecker struct {
	astwalk.WalkHandler
	ctx     *linter.CheckerContext
	ignores map[token.Pos]struct{}
}

func (p *protoGetterChecker) WalkFile(f *ast.File) {
	if !p.EnterFile(f) {
		return
	}

	for _, decl := range f.Decls {
		ast.Inspect(decl, func(node ast.Node) bool {
			if node != nil {
				p.VisitNode(node)
			}
			return true
		})
	}
}

func (p protoGetterChecker) VisitNode(n ast.Node) {
	var oldExpr, newExpr string

	switch x := n.(type) {
	case *ast.AssignStmt:
		for _, lhs := range x.Lhs {
			p.ignores[lhs.Pos()] = struct{}{}
		}

	case *ast.CallExpr:
		f, ok := x.Fun.(*ast.SelectorExpr)
		if !ok || !isProtoMessage(p.ctx, f.X) {
			for _, arg := range x.Args {
				var a *ast.UnaryExpr
				a, ok = arg.(*ast.UnaryExpr)
				if !ok || a.Op != token.AND {
					continue
				}

				p.ignores[a.X.Pos()] = struct{}{}
			}

			p.ignores[x.Pos()] = struct{}{}
			return
		}

		oldExpr, newExpr = makeFromCallAndSelectorExpr(x)
		if oldExpr == "" || newExpr == "" {
			p.ignores[x.Pos()] = struct{}{}
			return
		}

	case *ast.SelectorExpr:
		if !isProtoMessage(p.ctx, x.X) {
			return
		}

		oldExpr, newExpr = handleExpr(x, x.X)
	}

	if _, ok := p.ignores[n.Pos()]; ok {
		return
	}

	if oldExpr == "" || newExpr == "" {
		return
	}
	p.ignores[n.Pos()] = struct{}{}

	if oldExpr == newExpr {
		return
	}

	p.ctx.WarnFixable(n, linter.QuickFix{
		From:        n.Pos(),
		To:          n.End(),
		Replacement: []byte(newExpr),
	}, `proto message field read without getter: %q should be %q`, oldExpr, newExpr)
}

func makeFromCallAndSelectorExpr(expr *ast.CallExpr) (oldExpr, newExpr string) {
	switch f := expr.Fun.(type) {
	case *ast.SelectorExpr:
		oldExpr, newExpr = handleExpr(nil, f.X)
		if oldExpr == "" || newExpr == "" {
			return "", ""
		}

		oldExpr = fmt.Sprintf("%s.%s()", oldExpr, f.Sel.Name)
		newExpr = fmt.Sprintf("%s.%s()", newExpr, f.Sel.Name)

	default:
		fmt.Printf("makeFromCallAndSelectorExpr: not implemented for type: %s\n", reflect.TypeOf(f))
	}

	return oldExpr, newExpr
}

func handleExpr(base, child ast.Expr) (newExpr, oldExpr string) {
	switch c := child.(type) {
	case *ast.Ident:
		oldExpr, newExpr = handleIdent(base, c)

	case *ast.SelectorExpr:
		oldExpr, newExpr = handleSelectorExpr(base, c)

	case *ast.IndexExpr:
		oldExpr, newExpr = handleIndexExpr(base, c)

	case *ast.CallExpr:
		oldExpr, newExpr = handleCallExpr(base, c)

	default:
		fmt.Printf("handleExpr: not implemented for type: %s\n", reflect.TypeOf(c))
	}

	return oldExpr, newExpr

}

func handleIdent(base ast.Expr, c *ast.Ident) (oldExpr, newExpr string) {
	if base == nil {
		return "", ""
	}

	switch b := base.(type) {
	case *ast.SelectorExpr:
		oldExpr = fmt.Sprintf("%s.%s", c.Name, b.Sel.Name)
		newExpr = fmt.Sprintf("%s.Get%s()", c.Name, b.Sel.Name)

	case *ast.IndexExpr:
		var index string
		switch i := b.Index.(type) {
		case *ast.BasicLit:
			index = i.Value

		case *ast.Ident:
			index = i.Name

		default:
			fmt.Printf("handleIdent: base is IndexExpr: not implemented for type: %s\n", reflect.TypeOf(i))
		}

		oldExpr = fmt.Sprintf("%s[%s]", c.Name, index)
		newExpr = fmt.Sprintf("%s[%s]", c.Name, index)

	default:
		fmt.Printf("handleIdent: not implemented for type: %s\n", reflect.TypeOf(b))
	}

	return oldExpr, newExpr
}

func handleSelectorExpr(base ast.Expr, c *ast.SelectorExpr) (oldExpr, newExpr string) {
	oldExpr, newExpr = handleExpr(c, c.X)

	if base == nil {
		return oldExpr, newExpr
	}

	switch b := base.(type) {
	case *ast.SelectorExpr:
		oldExpr = fmt.Sprintf("%s.%s", oldExpr, b.Sel.Name)
		newExpr = fmt.Sprintf("%s.Get%s()", newExpr, b.Sel.Name)

	case *ast.CallExpr:
		oldExpr += "()"
		newExpr = strings.ReplaceAll(newExpr, "GetGet", "Get")

	default:
		fmt.Printf("handleSelectorExpr: not implemented for type: %s\n", reflect.TypeOf(b))
	}

	return oldExpr, newExpr
}

func handleCallExpr(base ast.Expr, c *ast.CallExpr) (newExpr, oldExpr string) {
	oldExpr, newExpr = handleExpr(c, c.Fun)

	if base == nil {
		return oldExpr, newExpr
	}

	switch b := base.(type) {
	case *ast.SelectorExpr:
		oldExpr = fmt.Sprintf("%s.%s", oldExpr, b.Sel.Name)
		newExpr = fmt.Sprintf("%s.Get%s()", newExpr, b.Sel.Name)

	default:
		fmt.Printf("handleCallExpr: not implemented for type: %s\n", reflect.TypeOf(b))
	}

	return oldExpr, newExpr
}

func handleIndexExpr(base ast.Expr, c *ast.IndexExpr) (oldExpr, newExpr string) {
	oldExpr, newExpr = handleExpr(c, c.X)

	if base == nil {
		return oldExpr, newExpr
	}

	switch b := base.(type) {
	case *ast.SelectorExpr:
		oldExpr = fmt.Sprintf("%s.%s", oldExpr, b.Sel.Name)
		newExpr = fmt.Sprintf("%s.Get%s()", newExpr, b.Sel.Name)

	default:
		fmt.Printf("handleIndexExpr: not implemented for type: %s\n", reflect.TypeOf(b))
	}

	return oldExpr, newExpr
}

func isProtoMessage(ctx *linter.CheckerContext, expr ast.Expr) bool {
	t := ctx.TypesInfo.TypeOf(expr)
	if t == nil {
		return false
	}
	ptr, ok := t.Underlying().(*types.Pointer)
	if !ok {
		return false
	}
	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}
	sct, ok := named.Underlying().(*types.Struct)
	if !ok {
		return false
	}
	if sct.NumFields() == 0 {
		return false
	}

	field := sct.Field(0)

	hasProtoImport := func() bool {
		for _, i := range field.Pkg().Imports() {
			if i.Path() == "google.golang.org/protobuf/reflect/protoreflect" {
				return true
			}

		}

		return false
	}()

	return field.Name() == "state" && hasProtoImport
}
