package checkers

import (
	"go/ast"
	"go/token"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/framework/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "typeDefFirst"
	info.Tags = []string{"style", "experimental"}
	info.Summary = "File-scoped checker, that requires type definition before its method definitions"
	info.Before = `
func (r rec) Method() {}
type rec struct{}
`
	info.After = `
type rec struct{}
func (r rec) Method() {}
`
	collection.AddChecker(&info, func(ctx *linter.CheckerContext) linter.FileWalker {
		return &typeDefFirstChecker{
			ctx: ctx,
		}
	})
}

type typeDefFirstChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *typeDefFirstChecker) WalkFile(f *ast.File) {
	if f.Decls == nil {
		return
	}

	typeUsageMap := make(map[string]bool)
	for _, declaration := range f.Decls {
		switch decl := declaration.(type) {
		case *ast.FuncDecl:
			if decl.Recv != nil {
				receiver := decl.Recv.List[0]
				typeName := c.receiverType(receiver.Type)
				typeUsageMap[typeName] = true
			}

		case *ast.GenDecl:
			if decl.Tok != token.TYPE {
				continue
			}
			for _, spec := range decl.Specs {
				if spec, ok := spec.(*ast.TypeSpec); ok {
					typeName := spec.Name.Name
					if val, ok := typeUsageMap[typeName]; ok && val {
						c.warn(decl, typeName)
					}
				}
			}
		}
	}
}

func (c *typeDefFirstChecker) warn(cause ast.Node, typeName string) {
	c.ctx.Warn(cause, "definition of type '%s' should appear before its methods", typeName)
}

func (c *typeDefFirstChecker) receiverType(e ast.Expr) string {
	switch e := e.(type) {
	case *ast.StarExpr:
		return c.receiverType(e.X)
	case *ast.Ident:
		return e.Name
	default:
		panic("unreachable")
	}
}
