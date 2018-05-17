package lint

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/Quasilyte/astcmp"
)

type longChainChecker struct {
	ctx *Context
}

func newLongChainChecker(ctx *Context) Checker {
	return &longChainChecker{
		ctx: ctx,
	}
}

func (c *longChainChecker) Check(f *ast.File) {
	for _, decl := range f.Decls {
		if decl, ok := decl.(*ast.FuncDecl); ok {
			if decl.Body == nil {
				continue
			}
			for _, stmt := range decl.Body.List {
				if stmt, ok := stmt.(*ast.SwitchStmt); ok {
					c.checkSwitch(stmt)
				}
			}
		}
	}
}

func (c *longChainChecker) checkSwitch(stmt *ast.SwitchStmt) {
	exprs := []ast.Expr{}
	for _, st := range stmt.Body.List {
		s := st.(*ast.CaseClause)
		if s.List == nil {
			continue
		}
		exprs = append(exprs, s.List...)
	}
	if len(exprs) < 2 {
		return
	}

	cp := c.exprToList(exprs[0])
	for _, e := range exprs[1:] {
		cp = c.commonPrefix(cp, c.exprToList(e))
	}

	const n = 3

	if len(cp) > n {
		c.warn(cp, stmt)
	}
}

func (c *longChainChecker) exprToList(expr ast.Expr) []ast.Expr {
	res := []ast.Expr{}
	tmp := expr
	for {
		switch t := tmp.(type) {
		case *ast.SelectorExpr:
			res = append(res, t.Sel)
			tmp = t.X
		default:
			res = append(res, tmp)
			reverseExprSlice(res)
			return res

		}
	}
}

func reverseExprSlice(e []ast.Expr) {
	for i := len(e)/2 - 1; i >= 0; i-- {
		j := len(e) - 1 - i
		e[i], e[j] = e[j], e[i]
	}
}

func (c *longChainChecker) commonPrefix(xs, ys []ast.Expr) []ast.Expr {
	res := []ast.Expr{}

	l := c.min(len(xs), len(ys))

	for i := 0; i < l; i++ {
		if astcmp.EqualExpr(xs[i], ys[i]) {
			res = append(res, xs[i])
		} else {
			break
		}
	}

	return res
}

func (c *longChainChecker) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *longChainChecker) warn(exprs []ast.Expr, stmt ast.Stmt) {
	s := make([]string, len(exprs))
	for i, e := range exprs {
		s[i] = nodeString(c.ctx.FileSet, e)
	}

	c.ctx.addWarning(Warning{
		Kind: "long-chain",
		Node: stmt,
		Text: fmt.Sprintf("Expression chain %s repeated multiple times consider assigning it to local variable",
			strings.Join(s, ".")),
	})
}
