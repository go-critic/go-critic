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
	for _, s := range stmt.Body.List {
		cas := s.(*ast.CaseClause)
		if cas.List == nil {
			continue
		}
		for _, expr := range cas.List {
			exprs = append(exprs, expr)
		}
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

			for i := len(res)/2 - 1; i >= 0; i-- {
				opp := len(res) - 1 - i
				res[i], res[opp] = res[opp], res[i]
			}
			return res

		}
	}
}

func (c *longChainChecker) commonPrefix(a, b []ast.Expr) []ast.Expr {
	res := []ast.Expr{}

	len := c.min(len(a), len(b))

	for i := 0; i < len; i++ {
		if astcmp.EqualExpr(a[i], b[i]) {
			res = append(res, a[i])
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
	var s []string
	for _, e := range exprs {
		s = append(s, nodeString(c.ctx.FileSet, e))
	}

	c.ctx.addWarning(Warning{
		Kind: "long-chain",
		Node: stmt,
		Text: fmt.Sprintf("Expression chain %s repeated multiple times consider assigning it to local variable",
			strings.Join(s, ".")),
	})
}
