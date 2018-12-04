package checkers

import (
	"go/ast"
	"go/token"

	"github.com/go-lintpack/lintpack"
	"github.com/go-lintpack/lintpack/astwalk"
	"github.com/go-toolsmith/astcast"
	"github.com/go-toolsmith/astcopy"
	"github.com/go-toolsmith/astequal"
	"github.com/go-toolsmith/typep"
)

func init() {
	var info lintpack.CheckerInfo
	info.Name = "nilPtrTest"
	info.Tags = []string{"diagnostic", "experimental"}
	info.Summary = "Detects potentially erroneous nil pointer checks"
	info.Before = `if x == nil { return *x }`
	info.After = `if x != nil { return *x }`

	collection.AddChecker(&info, func(ctx *lintpack.CheckerContext) lintpack.FileWalker {
		return astwalk.WalkerForStmt(&nilPtrTestChecker{ctx: ctx})
	})
}

type nilPtrTestChecker struct {
	astwalk.WalkHandler
	ctx *lintpack.CheckerContext

	foundDeref bool
}

func (c *nilPtrTestChecker) VisitStmt(stmt ast.Stmt) {
	ifStmt := astcast.ToIfStmt(stmt)
	cmp := astcast.ToBinaryExpr(ifStmt.Cond)
	if cmp.Op != token.EQL {
		return
	}
	if astcast.ToIdent(cmp.Y).Name != "nil" {
		return
	}
	ptr := cmp.X
	if !typep.IsPointer(c.ctx.TypesInfo.TypeOf(ptr)) {
		return
	}

	// ptr is a pointer and it's checked to be nil.
	// If it's dereferenced somewhere inside that if statement
	// witout initialization, emit a warning.

	// For convenience, do it in a 2 passes.
	// First, ensure that there is no assignments to ptr.
	ptrInitialized := containsNode(ifStmt.Body, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.AssignStmt:
			if n.Tok != token.ASSIGN {
				return false
			}
			for _, assignee := range n.Lhs {
				if astequal.Expr(ptr, assignee) {
					return true
				}
			}
		case *ast.UnaryExpr:
			// Someone takes ptr address.
			// It might lead to indirect initialization
			// and we're not performing pointer analysis here.
			// Just give up.
			if n.Op == token.AND {
				return true
			}
		}
		return false
	})
	if ptrInitialized {
		// TODO(Quasilyte): less fragile analysis?
		// We can try to understand the program flow
		// instead of giving up so easily.
		// Maybe SSA form could work here better.
		return
	}

	// TODO(Quasilyte): see positive tests for more insights.
	// Implicit dereferences are also a thing.
	dereferenced := containsNode(ifStmt.Body, func(n ast.Node) bool {
		return astequal.Expr(astcast.ToStarExpr(n).X, ptr)
	})

	if dereferenced {
		c.warn(ifStmt, cmp)
	}
}

func (c *nilPtrTestChecker) warn(cause ast.Node, cmp *ast.BinaryExpr) {
	suggest := astcopy.BinaryExpr(cmp)
	suggest.Op = token.NEQ
	c.ctx.Warn(cause, "nil ptr deref possible; probably `%s` was intended", suggest)
}
