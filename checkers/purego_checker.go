package checkers

import (
	"go/ast"
	"go/build/constraint"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "purego"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects bad `//go:build purego` usage"
	info.Before = `//go:build purego
import "unsafe"`
	info.After = `//go:build purego
import "safer_unsafe"
`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return &puregoChecker{ctx: ctx}, nil
	})
}

type puregoChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *puregoChecker) WalkFile(f *ast.File) {
	for _, importDcl := range f.Imports {
		if c.isUnsafeImport(importDcl.Path.Value) {
			c.processComments(f.Comments)
			return
		}
	}
}

func (c *puregoChecker) isUnsafeImport(pkg string) bool {
	switch pkg {
	case "unsafe", "C":
		return true
	default:
		return false
	}
}

func (c *puregoChecker) processComments(comments []*ast.CommentGroup) {
	for _, cg := range comments {
		for _, line := range cg.List {
			s := line.Text
			if !constraint.IsGoBuild(s) {
				continue
			}

			expr, err := constraint.Parse(s)
			if err != nil {
				continue
			}

			ok := expr.Eval(func(tag string) bool {
				return tag == "purego"
			})

			if ok {
				c.warn(cg.List[0])
				return
			}
		}
	}
}

func (c *puregoChecker) warn(comment *ast.Comment) {
	c.ctx.Warn(comment, "Importing `unsafe` or `C` with `go:build purego` is not allowed")
}
