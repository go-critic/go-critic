package checkers

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"

	"github.com/go-critic/go-critic/checkers/internal/astwalk"
	"github.com/go-critic/go-critic/framework/linter"
)

func init() {
	var info linter.CheckerInfo
	info.Name = "caseOrder"
	info.Tags = []string{"diagnostic"}
	info.Summary = "Detects erroneous case order inside switch statements"
	info.Before = `
switch x.(type) {
case ast.Expr:
	fmt.Println("expr")
case *ast.BasicLit:
	fmt.Println("basic lit") // Never executed
}`
	info.After = `
switch x.(type) {
case *ast.BasicLit:
	fmt.Println("basic lit") // Now reachable
case ast.Expr:
	fmt.Println("expr")
}`

	collection.AddChecker(&info, func(ctx *linter.CheckerContext) (linter.FileWalker, error) {
		return astwalk.WalkerForStmt(&caseOrderChecker{ctx: ctx}), nil
	})
}

type caseOrderChecker struct {
	astwalk.WalkHandler
	ctx *linter.CheckerContext
}

func (c *caseOrderChecker) VisitStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.TypeSwitchStmt:
		c.checkTypeSwitch(stmt)
	case *ast.SwitchStmt:
		c.checkSwitch(stmt)
	}
}

func (c *caseOrderChecker) checkTypeSwitch(s *ast.TypeSwitchStmt) {
	type ifaceType struct {
		node ast.Node
		typ  *types.Interface
	}
	var ifaces []ifaceType // Interfaces seen so far
	for _, cc := range s.Body.List {
		cc := cc.(*ast.CaseClause)
		for _, x := range cc.List {
			typ := c.ctx.TypeOf(x)
			if typ == linter.UnknownType {
				c.warnUnknownType(cc, x)
				return
			}
			for _, iface := range ifaces {
				if types.Implements(typ, iface.typ) {
					c.warnTypeSwitch(cc, x, iface.node)
					break
				}
			}
			if iface, ok := typ.Underlying().(*types.Interface); ok {
				ifaces = append(ifaces, ifaceType{node: x, typ: iface})
			}
		}
	}
}

func (c *caseOrderChecker) warnTypeSwitch(cause, concrete, iface ast.Node) {
	c.ctx.Warn(cause, "case %s must go before the %s case", concrete, iface)
}

func (c *caseOrderChecker) warnUnknownType(cause, concrete ast.Node) {
	c.ctx.Warn(cause, "type is not defined %s", concrete)
}

func (c *caseOrderChecker) checkSwitch(stmt *ast.SwitchStmt) {
	var cases []*ast.CaseClause

	for i := range stmt.Body.List {
		curCase := stmt.Body.List[i].(*ast.CaseClause)
		for _, cc := range cases {
			if isOverlappedCases(cc, curCase) {
				c.warnSwitch(curCase, curCase, cc)
				break
			}
		}

		cases = append(cases, curCase)
	}
}

// isOverlappedCases - check that case1 wider value range than case2
func isOverlappedCases(case1, case2 *ast.CaseClause) bool {
	var (
		y1, y2               *ast.BasicLit
		operator1, operator2 token.Token
	)

	compare := func(lss func() bool, eql func() bool) bool {
		less := lss()
		equal := eql()

		if operator2 == token.EQL {
			if (operator1 == token.LSS && less) || (operator1 == token.LEQ && (equal || less)) {
				return true
			}
			if (operator1 == token.GTR && !less) || (operator1 == token.GEQ && (equal || !less)) {
				return true
			}
		}

		return false
	}

	exprs1, exprs2 := collectExpressions(case1), collectExpressions(case2)

	for i := range exprs1 {
		y1 = exprs1[i].value
		operator1 = exprs1[i].operator
		for ii := range exprs2 {
			y2 = exprs2[ii].value
			operator2 = exprs2[ii].operator

			if y1.Kind != y2.Kind {
				continue
			}

			switch y1.Kind {
			case token.INT:
				v1, _ := strconv.Atoi(y1.Value)
				v2, _ := strconv.Atoi(y2.Value)

				if compare(func() bool { return v1 > v2 }, func() bool { return v1 == v2 }) {
					return true
				}
			case token.FLOAT:
				v1, _ := strconv.ParseFloat(y1.Value, 64)
				v2, _ := strconv.ParseFloat(y2.Value, 64)

				if compare(func() bool { return v1 > v2 }, func() bool { return v1 == v2 }) {
					return true
				}
			}
		}
	}

	return false
}

type expression struct {
	value    *ast.BasicLit
	operator token.Token
}

func collectExpressions(cc *ast.CaseClause) []expression {
	var (
		exprs    = make([]expression, 0, 1)
		operator token.Token
	)

	invertOperator := func(op token.Token) token.Token {
		switch op {
		case token.LEQ:
			return token.GEQ
		case token.LSS:
			return token.GTR
		case token.GEQ:
			return token.LEQ
		case token.GTR:
			return token.LSS
		default:
			return op
		}
	}

	for i := range cc.List {
		expr, ok := cc.List[i].(*ast.BinaryExpr)
		if !ok {
			continue
		}

		y, ok := expr.Y.(*ast.BasicLit)
		if !ok {
			y, ok = expr.X.(*ast.BasicLit)
			if !ok {
				continue
			}

			operator = invertOperator(expr.Op)
		} else {
			operator = expr.Op
		}

		exprs = append(exprs, expression{
			value:    y,
			operator: operator,
		})
	}

	return exprs
}

func (c *caseOrderChecker) warnSwitch(cause ast.Node, concrete, node *ast.CaseClause) {
	var args []interface{}
	prettyPrint := func(cc *ast.CaseClause) string {
		s := "case %s"
		args = append(args, interface{}(cc.List[0]))
		if len(cc.List) == 1 {
			return s
		}

		for _, l := range cc.List[1:] {
			s += ", %s"
			args = append(args, interface{}(l))
		}

		return s
	}

	c.ctx.Warn(cause, prettyPrint(concrete)+" should go before the "+prettyPrint(node), args...)
}
