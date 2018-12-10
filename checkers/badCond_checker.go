package checkers

import (
	"go/ast"
	"go/constant"
	"go/token"

	"github.com/go-critic/go-critic/checkers/internal/lintutil"
	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/astequal"
	"github.com/go-toolsmith/typep"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "badCond"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects suspicious condition expressions"
	info.Before = `
for i := 0; i > n; i++ {
	xs[i] = 0
}`
	info.After = `
for i := 0; i < n; i++ {
	xs[i] = 0
}`

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return astwalk.WalkerForFuncDecl(&badCondChecker{ctx: ctx})
	})
}

type badCondChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext
}

func (c *badCondChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	ast.Inspect(decl.Body, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.ForStmt:
			c.checkForStmt(n)
		case ast.Expr:
			c.checkExpr(n)
		}
		return true
	})
}

func (c *badCondChecker) checkExpr(expr ast.Expr) {
	// x < a && x > b; Where `a` less than `b`.
	// TODO(Quasilyte): recognize more patterns.

	cond := astcast.ToBinaryExpr(expr)
	lt := astcast.ToBinaryExpr(cond.X)
	gt := astcast.ToBinaryExpr(cond.Y)
	if cond.Op != token.LAND || lt.Op != token.LSS || gt.Op != token.GTR {
		return
	}
	if !astequal.Expr(lt.X, gt.X) {
		return
	}
	a := c.ctx.TypesInfo.Types[lt.Y].Value
	b := c.ctx.TypesInfo.Types[gt.Y].Value
	if a == nil || b == nil || !constant.Compare(a, token.LSS, b) {
		return
	}

	c.warnComparison(expr, cond)
}

func (c *badCondChecker) checkForStmt(stmt *ast.ForStmt) {
	// TODO(Quasilyte): handle other kinds of bad conditionals.

	init := astcast.ToAssignStmt(stmt.Init)
	if init.Tok != token.DEFINE || len(init.Lhs) != 1 || len(init.Rhs) != 1 {
		return
	}
	if astcast.ToBasicLit(init.Rhs[0]).Value != "0" {
		return
	}

	iter := astcast.ToIdent(init.Lhs[0])
	cond := astcast.ToBinaryExpr(stmt.Cond)
	if cond.Op != token.GTR || !astequal.Expr(iter, cond.X) {
		return
	}
	if !typep.SideEffectFree(c.ctx.TypesInfo, cond.Y) {
		return
	}

	post := astcast.ToIncDecStmt(stmt.Post)
	if post.Tok != token.INC || !astequal.Expr(iter, post.X) {
		return
	}

	mutated := lintutil.CouldBeMutated(c.ctx.TypesInfo, stmt.Body, cond.Y) ||
		lintutil.CouldBeMutated(c.ctx.TypesInfo, stmt.Body, iter)
	if mutated {
		return
	}

	c.warnForStmt(stmt, cond)
}

func (c *badCondChecker) warnForStmt(cause ast.Node, cond *ast.BinaryExpr) {
	suggest := astcopy.BinaryExpr(cond)
	suggest.Op = token.LSS
	c.ctx.Warn(cause, "`%s` in loop; probably meant `%s`?",
		cond, suggest)
}

func (c *badCondChecker) warnComparison(cause ast.Node, cond *ast.BinaryExpr) {
	suggest := astcopy.BinaryExpr(cond)
	suggest.Op = token.LOR
	c.ctx.Warn(cause, "`%s` is always false; probably meant `%s`?",
		cond, suggest)
}
