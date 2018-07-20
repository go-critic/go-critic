package lint

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/lint/internal/lintutil"
	"github.com/go-toolsmith/astequal"
	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	addChecker(&longChainChecker{}, attrExperimental)
}

type longChainChecker struct {
	checkerBase

	chains      lintutil.AstMap
	reportedSet lintutil.AstMap

	minComplexity int
	minChainLen   int
}

func (c *longChainChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects repeated expression chains and suggest to refactor them"
	d.Before = `
v := q.w.e.r[0].x + q.w.e.r[0].y + q.w.e.r[0].z`
	d.After = `
r := q.w.e.r[0]
v := r.x + r.y + r.z`
}

func (c *longChainChecker) Init() {
	c.minComplexity = c.ctx.params.Int("minComplexity", 8)
	c.minChainLen = c.ctx.params.Int("minChainLen", 3)
}

func (c *longChainChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	ast.Inspect(decl.Body, func(x ast.Node) bool {
		switch x := x.(type) {
		case *ast.AssignStmt:
			c.collectExprList(x.Rhs)
			c.check(x.Rhs[0])
		case *ast.SwitchStmt:
			for _, cc := range x.Body.List {
				cc := cc.(*ast.CaseClause)
				c.collectExprList(cc.List)
			}
			c.check(x)
		default:
			return true
		}
		c.reset()
		return false
	})
}

func (c *longChainChecker) reset() {
	c.chains.Clear()
	c.reportedSet.Clear()
}

func (c *longChainChecker) collectExprList(list []ast.Expr) {
	for _, expr := range list {
		ast.Inspect(expr, func(x ast.Node) bool {
			if x, ok := x.(ast.Expr); ok {
				c.checkExpr(x)
			}
			return true
		})
	}
}

func (c *longChainChecker) check(cause ast.Node) {
	for i := 0; i < c.chains.Len(); i++ {
		if *c.chains.ValueAt(i).(*int) < c.minChainLen {
			continue
		}
		node := c.chains.KeyAt(i)
		isSubExpr := false
		for j := 0; j < c.reportedSet.Len(); j++ {
			if c.astContains(c.reportedSet.KeyAt(j), node) {
				isSubExpr = true
				break
			}
		}

		// We don't report sub-expression repetitions
		// of already reported expressions.
		// That means if we already reported "a.b.c.d.e",
		// don't bother reporting repeated "b.c.d".
		if !isSubExpr {
			c.warn(cause, node)
		}
		c.reportedSet.Insert(node, nil)
	}
}

func (c *longChainChecker) astContains(root, sub ast.Node) bool {
	return containsNode(root, func(x ast.Node) bool {
		return astequal.Node(x, sub)
	})
}

func (c *longChainChecker) checkExpr(expr ast.Expr) {
	if c.exprComplexity(expr) < c.minComplexity {
		return
	}

	if count, ok := c.chains.Find(expr).(*int); ok {
		*count++
	} else {
		count := 1
		c.chains.Insert(expr, &count)
	}
}

// exprComplexity returns simplified "cost" value of expression.
// This is not evaluation complexity but rather syntactical weight of the expression.
func (c *longChainChecker) exprComplexity(expr ast.Expr) int {
	cost := 0
	astutil.Apply(expr, nil, func(cur *astutil.Cursor) bool {
		// TODO(quasilyte): make this type switch more exhaustive.
		switch expr := cur.Node().(type) {
		case *ast.ParenExpr:
			// Does not increase cost.
		case *ast.CallExpr:
			// Consider type conversions as safe call expressions.
			// All other call expressions forbid further analysis.
			if !c.isTypeConversion(expr) {
				cost = 0
				return false
			}
			cost++
		case *ast.SelectorExpr, *ast.IndexExpr, *ast.Ident, *ast.BasicLit, *ast.BinaryExpr:
			cost++
		default:
			// For now, skip all nodes that are not explicitly handled.
			cost = 0
			return false
		}
		return true
	})
	return cost
}

// isTypeConversion reports whether call is a type conversion expression.
// That is, T(v) expression that has no side-effects.
func (c *longChainChecker) isTypeConversion(call *ast.CallExpr) bool {
	// Three main conversion cases:
	//	1. T(v)   - call.Fun is *ast.Ident
	//	2. p.T(V) - call.Fun is *ast.SelectorExpr
	//	3. (x)(V) - x if either (1) or (2)
	// T is a type name.

	switch fn := astutil.Unparen(call.Fun).(type) {
	case *ast.Ident:
		_, ok := c.ctx.typesInfo.ObjectOf(fn).(*types.TypeName)
		return ok

	case *ast.SelectorExpr:
		pkg, ok := fn.X.(*ast.Ident)
		if !ok {
			return false
		}
		if _, ok := c.ctx.typesInfo.ObjectOf(pkg).(*types.PkgName); !ok {
			return false
		}
		_, ok = c.ctx.typesInfo.ObjectOf(fn.Sel).(*types.TypeName)
		return ok

	default:
		return false
	}
}

func (c *longChainChecker) warn(cause, expr ast.Node) {
	c.ctx.Warn(cause, "%s repeated multiple times, consider assigning it to local variable", expr)
}
