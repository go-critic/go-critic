package checkers

import (
	"go/ast"
	"go/token"
	"strings"

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

func (t *typeDefFirstChecker) WalkFile(f *ast.File) {
	typeUsageMap := make(map[string]bool)

	if f.Decls == nil {
		return
	}

	for _, declaration := range f.Decls {
		switch decl := declaration.(type) {
		case *ast.FuncDecl:
			if decl.Recv != nil {
				receiver := decl.Recv.List[0]
				typeName := trimAsterisk(NodeToString(t.ctx.FileSet, receiver.Type))
				typeUsageMap[typeName] = true
			}

		case *ast.GenDecl:
			if decl.Tok == token.TYPE {
				for _, spec := range decl.Specs {
					if spec, ok := spec.(*ast.TypeSpec); ok {
						typeName := trimAsterisk(spec.Name.Name)
						if val, ok := typeUsageMap[typeName]; ok && val {
							t.warn(decl, typeName)
						}
					}
				}
			}
		}
	}
}

func (t *typeDefFirstChecker) warn(cause ast.Node, typeName string) {
	t.ctx.Warn(cause, "definition of type '%s' should appear before its methods", typeName)
}

// removing '*' from type
func trimAsterisk(typeName string) string {
	return strings.TrimLeft(typeName, "* ")
}
