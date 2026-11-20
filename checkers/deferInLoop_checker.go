package checkers

import (
	"go/ast"
	"go/token"
	"slices"

	"github.com/go-toolsmith/astp"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "deferInLoop"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects loops inside functions that use defer"
	info.Before = `
for _, filename := range []string{"foo", "bar"} {
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
	ctx   *linter.CheckerContext
	inFor bool
}

func (c *deferInLoopChecker) VisitFuncDecl(fn *ast.FuncDecl) {
	ast.Inspect(fn.Body, c.traversalFunc)
}

func (c deferInLoopChecker) traversalFunc(cur ast.Node) bool {
	switch n := cur.(type) {
	case *ast.BlockStmt:
		if c.inFor && len(n.List) > 1 && astp.IsReturnStmt(n.List[len(n.List)-1]) {
			c.inspectReturningBlock(n.List)
			return false
		}
	case *ast.DeferStmt:
		if c.inFor {
			c.warn(n)
		}
	case *ast.RangeStmt, *ast.ForStmt:
		if !c.inFor {
			ast.Inspect(cur, deferInLoopChecker{ctx: c.ctx, inFor: true}.traversalFunc)
			return false
		}
	case *ast.FuncLit:
		ast.Inspect(n.Body, deferInLoopChecker{ctx: c.ctx, inFor: false}.traversalFunc)
		return false
	case nil:
		return false
	}
	return true
}

func (c deferInLoopChecker) inspectReturningBlock(stmts []ast.Stmt) {
	for len(stmts) != 0 && !astp.IsDeferStmt(stmts[0]) {
		ast.Inspect(stmts[0], c.traversalFunc)
		stmts = stmts[1:]
	}

	// A direct defer after the last escaping jump cannot repeat before the return.
	firstSafe := 0
	for i := len(stmts) - 2; i > 0; i-- {
		if c.branchLeaves(stmts[i]) {
			firstSafe = i + 1
			break
		}
	}
	for i, stmt := range stmts {
		if d, ok := stmt.(*ast.DeferStmt); ok && i >= firstSafe {
			// Its call can still contain a function literal with a loop of its own.
			ast.Inspect(d.Call, c.traversalFunc)
		} else {
			ast.Inspect(stmt, c.traversalFunc)
		}
	}
}

// branchLeaves reports whether a jump can leave stmt and bypass the block's return.
func (c deferInLoopChecker) branchLeaves(stmt ast.Stmt) bool {
	leaves := false
	ast.PreorderStack(stmt, nil, func(n ast.Node, parents []ast.Node) bool {
		if leaves || astp.IsFuncLit(n) {
			return false
		}
		branch, ok := n.(*ast.BranchStmt)
		if !ok {
			return true
		}
		if branch.Label != nil {
			target := c.ctx.TypesInfo.ObjectOf(branch.Label)
			leaves = target == nil || target.Pos() < stmt.Pos() || target.Pos() >= stmt.End()
		} else if branch.Tok != token.FALLTHROUGH {
			leaves = !slices.ContainsFunc(parents, func(parent ast.Node) bool {
				switch parent.(type) {
				case *ast.ForStmt, *ast.RangeStmt:
					return true
				case *ast.SwitchStmt, *ast.TypeSwitchStmt, *ast.SelectStmt:
					return branch.Tok == token.BREAK
				default:
					return false
				}
			})
		}
		return false
	})
	return leaves
}

func (c *deferInLoopChecker) warn(cause *ast.DeferStmt) {
	c.ctx.Warn(cause, "Possible resource leak, 'defer' is called in the 'for' loop")
}
