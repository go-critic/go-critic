package checkers

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/framework/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "deferInCycle"
	info.Tags = []string{"diagnostic"}
	info.Summary = "Detects cycles inside functions that use defer"
	info.Before = `
for _, filename := range []string{"kek", "shrek"} {
	 f, err := os.Open(filename)
	
	defer f.Close()
}
`
	info.After = `
func process(filename string) {
	 f, err := os.Open(filename)
	
	defer f.Close()
}
...
...
for _, filename := range []string{"kek", "shrek"} {
	process(filename)
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&deferInCycleChecker{ctx: ctx}), nil
	})
}

type deferInCycleChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *deferInCycleChecker) VisitFuncDecl(fn *ast.FuncDecl) {
	var forBlock *ast.BlockStmt
	pre := func(cur *astutil.Cursor) bool {
		switch n := cur.Node().(type) {
		case *ast.DeferStmt:
			// "defer" inside "for" body
			if forBlock != nil && forBlock.Pos() < n.Pos() && forBlock.End() > n.End() {
				c.warn(n)
			}
		case *ast.RangeStmt:
			// check if its inner loop skip
			if forBlock != nil && n.Pos() < forBlock.End() {
				return true
			}
			forBlock = n.Body
		case *ast.ForStmt:
			if forBlock != nil && n.Pos() < forBlock.End() {
				return true
			}
			forBlock = n.Body
		}
		return true
	}
	astutil.Apply(fn.Body, pre, nil)
}

func (c *deferInCycleChecker) warn(cause *ast.DeferStmt) {
	c.ctx.Warn(cause, "Possible resource leak, 'defer' is called in the 'for' loop")
}
