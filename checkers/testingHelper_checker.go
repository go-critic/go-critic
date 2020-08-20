package checkers

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/framework/linter"

	"github.com/go-toolsmith/astfmt"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "testingHelper"
	info.Tags = []string{"style", "opinionated", "experimental"}
	info.Summary = "TODO"
	info.Before = `TODO`
	info.After = `TODO`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) linter.FileWalker {
		return astwalk.WalkerForFuncDecl(&testingHelperChecker{ctx: ctx})
	})
}

type testingHelperChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *testingHelperChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	if isUnitTestFunc(c.ctx, decl) {
		return
	}

	params := decl.Type.Params
	if len(params.List) == 0 {
		return
	}

	typ := c.ctx.TypesInfo.TypeOf(decl.Name)
	if sig, ok := typ.(*types.Signature); ok {

		params := sig.Params()
		for i := 0; i < params.Len(); i++ {
			typ := params.At(i).Type().String()
			if typ == "*testing.T" {
				if i != 0 {
					c.warnFirstParam(decl)
				}
				if !c.containsHelperCall(decl.Body) {
					c.warnPossibleHelper(decl)
				}
			}
		}
	}
}

func (c *testingHelperChecker) containsHelperCall(body *ast.BlockStmt) bool {
	for _, s := range body.List {
		stmt, ok := s.(*ast.ExprStmt)
		if !ok {
			continue
		}
		call, ok := stmt.X.(*ast.CallExpr)
		if !ok {
			continue
		}
		if astfmt.Sprint(call) == "t.Helper()" {
			return true
		}
	}
	return false
}

func (c *testingHelperChecker) warnFirstParam(...interface{}) {
	// TODO
}
func (c *testingHelperChecker) warnPossibleHelper(...interface{}) {
	// TODO
}
