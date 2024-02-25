package checkers

import (
	"go/ast"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "deferUnlambda"
	info.Tags = []string{linter.StyleTag, linter.ExperimentalTag}
	info.Summary = "Detects deferred function literals that can be simplified"
	info.Before = "defer func() { f() }()"
	info.After = "defer f()"

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmtList(&deferUnlambdaChecker{ctx: ctx}), nil
	})
}

type deferUnlambdaChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *deferUnlambdaChecker) VisitStmtList(x ast.Node, list []ast.Stmt) {
	for i, stmt := range list {
		ds, ok := stmt.(*ast.DeferStmt)
		if !ok {
			continue
		}

		innerFuncIdent, callExpr, needsUnlambda := c.deferStmtHasOnlyOneFuncCallWithNoReturns(ds)
		if needsUnlambda && !c.isSilencedBuiltinFunc(innerFuncIdent) && !c.isIdentifierAssigned(innerFuncIdent, list[i:]) {
			c.ctx.Warn(stmt, "can rewrite as `defer %s`", callExpr)
		}
	}
}

func (c *deferUnlambdaChecker) deferStmtHasOnlyOneFuncCallWithNoReturns(ds *ast.DeferStmt) (*ast.Ident, ast.Expr, bool) {
	if len(ds.Call.Args) != 0 {
		return nil, nil, false
	}

	fl, ok := ds.Call.Fun.(*ast.FuncLit)
	if !ok {
		return nil, nil, false
	}

	if len(fl.Body.List) != 1 {
		return nil, nil, false
	}

	stmt := fl.Body.List[0]
	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return nil, nil, false
	}

	callExpr, ok := exprStmt.X.(*ast.CallExpr)
	if !ok {
		return nil, nil, false
	}
	callFuncIdent := identOf(callExpr.Fun)

	se, ok := callExpr.Fun.(*ast.SelectorExpr)
	if ok {
		// check if the func is called on an object
		objIdent, ok := se.X.(*ast.Ident)
		if !ok || objIdent.Obj != nil {
			return nil, nil, false
		}
	}

	if callFuncIdent != nil {
		return callFuncIdent, callExpr, c.allCallArgsAreConst(callExpr.Args)
	}
	return nil, nil, false
}

func (c *deferUnlambdaChecker) isIdentifierAssigned(target *ast.Ident, list []ast.Stmt) bool {
	for _, stmt := range list {
		assignStmt, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}

		for _, lhs := range assignStmt.Lhs {
			ident, ok := lhs.(*ast.Ident)
			if !ok {
				continue
			}

			if ident.Name == target.Name {
				return true
			}
		}
	}
	return false
}

func (c *deferUnlambdaChecker) isSilencedBuiltinFunc(i *ast.Ident) bool {
	if !isBuiltin(i.Name) {
		return false
	}
	return i.Name == "panic" || i.Name == "recover"
}

func (c *deferUnlambdaChecker) allCallArgsAreConst(args []ast.Expr) bool {
	for _, a := range args {
		if !c.isConstExpr(a) {
			return false
		}
	}

	return true
}

func (c *deferUnlambdaChecker) isConstExpr(e ast.Expr) bool {
	return c.ctx.TypesInfo.Types[e].Value != nil
}
