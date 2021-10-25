package checkers

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/framework/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "deferInLoop"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects loops inside functions that use defer"
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
/* ... */
for _, filename := range []string{"foo", "bar"} {
	process(filename)
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForFuncDecl(&deferInLoopChecker{ctx: ctx}), nil
	})
}

type deferInLoopChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *deferInLoopChecker) VisitFuncDecl(fn *ast.FuncDecl) {
	// TODO: add func args check
	// example: t.Run(123, func() { for { defer println(); break } })
	var blockParser func(*ast.BlockStmt, bool)
	blockParser = func(block *ast.BlockStmt, inFor bool) {
		for _, cur := range block.List {
			switch n := cur.(type) {
			case *ast.DeferStmt:
				if inFor {
					c.warn(n)
				}
			case *ast.RangeStmt:
				blockParser(n.Body, true)
			case *ast.ForStmt:
				blockParser(n.Body, true)
			case *ast.GoStmt:
				if f, ok := n.Call.Fun.(*ast.FuncLit); ok {
					blockParser(f.Body, false)
				}
			case *ast.ExprStmt:
				if f, ok := n.X.(*ast.CallExpr); ok {
					if anon, ok := f.Fun.(*ast.FuncLit); ok {
						blockParser(anon.Body, false)
					}
				}
			case *ast.BlockStmt:
				blockParser(n, inFor)
			}
		}
	}
	pre := func(cur *astutil.Cursor) bool {
		if n, ok := cur.Node().(*ast.BlockStmt); ok {
			blockParser(n, false)
		}
		return false
	}

	astutil.Apply(fn.Body, pre, nil)
}

func (c *deferInLoopChecker) warn(cause *ast.DeferStmt) {
	c.ctx.Warn(cause, "Possible resource leak, 'defer' is called in the 'for' loop")
}
