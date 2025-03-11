package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
	"golang.org/x/tools/go/ast/astutil"

	"github.com/go-toolsmith/astp"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "returnAfterHttpError"
	info.Tags = []string{linter.DiagnosticTag, linter.ExperimentalTag}
	info.Summary = "Detects suspicious http.Error call without following return"
	info.Before = `if err != nil { http.Error(...); }`
	info.After = `if err != nil { http.Error(...); return; }`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		c := &returnAfterHttpErrorChecker{ctx: ctx}
		return astwalk.WalkerForFuncDecl(c), nil
	})
}

type returnAfterHttpErrorChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *returnAfterHttpErrorChecker) VisitFuncDecl(fn *ast.FuncDecl) {
	var httpErrorCall *ast.CallExpr

	astutil.Apply(
		fn.Body,
		func(cur *astutil.Cursor) bool {
			if exprStmt, ok := cur.Node().(*ast.ExprStmt); ok {
				if callStmt, ok := exprStmt.X.(*ast.CallExpr); ok {
					if qualifiedName(callStmt.Fun) == "http.Error" {
						// Found http.Error call.
						httpErrorCall = callStmt
						return false
					}
				}
			}

			if httpErrorCall != nil {
				if astp.IsReturnStmt(cur.Node()) {
					// Good - Found return statement after http.Error call.
					httpErrorCall = nil
					return false
				}

				if astp.IsStmt(cur.Node()) {
					c.ctx.Warn(httpErrorCall, "Possibly return is missed after the http.Error call")
					// Clear the call after reporting it to avoid duplicate warnings.
					httpErrorCall = nil
				}
			}

			return true
		},
		nil,
	)
}
