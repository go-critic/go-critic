package lint

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

func init() {
	addChecker(&longChainChecker{}, attrExperimental)
}

type longChainChecker struct {
	checkerBase

	// chains is a {expr string => count} mapping.
	chains map[string]int

	// reported is a set of reported expression strings.
	// Used to avoid duplicated warnings.
	reported map[string]bool
}

func (c *longChainChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects repeated expression chains and suggest to refactor them"
	d.Before = `
a := q.w.e.r.t + 1
b := q.w.e.r.t + 2
c := q.w.e.r.t + 3
v := (a + xs[i+1]) + (b + xs[i+1]) + (c + xs[i+1])`
	d.After = `
x := xs[i+1]
qwert := q.w.e.r.t
a := qwert + 1
b := qwert + 2
c := qwert + 3
v := (a + x) + (b + x) + (c + x)`
}

func (c *longChainChecker) EnterFunc(fn *ast.FuncDecl) bool {
	// Avoid checking functions of 1 statement.
	// Both performance and false-positives reasons.
	if fn.Body == nil || len(fn.Body.List) == 1 {
		return false
	}

	c.chains = make(map[string]int)
	c.reported = make(map[string]bool)
	return true
}

func (c *longChainChecker) VisitLocalExpr(expr ast.Expr) {
	// These constants are purely heuristical.
	//
	// TODO: for very big functions should increase minChainLen
	// threshould or consider matching expressions only inside limited "window".
	// This may not be worthwhile. Needs discussion.
	const (
		minComplexity = 6
		minChainLen   = 3
	)

	if c.exprComplexity(expr) < minComplexity {
		return
	}

	s := c.ctx.printer.Sprint(expr)
	if c.reported[s] {
		return
	}
	c.chains[s]++
	if c.chains[s] >= minChainLen {
		isSubExpr := false
		for s2 := range c.reported {
			if strings.Contains(s2, s) {
				isSubExpr = true
				break
			}
		}

		// We don't report sub-expression repetitions
		// of already reported expressions.
		// That means if we already reported "a.b.c.d.e",
		// don't bother reporting repeated "b.c.d".
		if !isSubExpr {
			c.warn(expr, s)
		}
		c.reported[s] = true
	}
}

// exprComplexity returns simplified "cost" value of expression.
// This is not evaluation complexity but rather syntactical weight of the expression.
func (c *longChainChecker) exprComplexity(expr ast.Expr) int {
	cost := 0
	astutil.Apply(expr, nil, func(cur *astutil.Cursor) bool {
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
		default:
			cost++
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

func (c *longChainChecker) warn(node ast.Node, s string) {
	c.ctx.Warn(node, "%s repeated multiple times, consider assigning it to local variable", s)
}
