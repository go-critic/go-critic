package lint

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"

	"github.com/go-toolsmith/astp"
)

func init() {
	addChecker(&namedConstChecker{}, attrExperimental)
}

type namedConstChecker struct {
	checkerBase
}

func (c *namedConstChecker) InitDocs(d *Documentation) {
	d.Summary = "Detects literals that can be replaced with defined named const"
	d.Before = `
// pos has type of token.Pos.
return pos != 0`
	d.After = `return pos != token.NoPos`
}

func (c *namedConstChecker) EnterChilds(x ast.Node) bool {
	switch x := x.(type) {
	case *ast.BinaryExpr:
		// TODO(quasilyte): figure out how to handle
		// things like 1*time.Second without false
		// positives without adhoc check and without
		// introducing too much of false negatives.
		return !c.isTimeExpr(x)
	default:
		return true
	}
}

func (c *namedConstChecker) isTimeExpr(x *ast.BinaryExpr) bool {
	// The most sensible operations are:
	//	N * time.X
	//	timeExpr / N
	if x.Op != token.MUL && x.Op != token.QUO {
		return false
	}
	typ := c.ctx.typesInfo.TypeOf(x)
	if typ == nil || typ.String() != "time.Duration" {
		return false
	}
	return astp.IsBasicLit(x.X) || astp.IsBasicLit(x.Y)
}

func (c *namedConstChecker) VisitExpr(x ast.Expr) {
	if !astp.IsBasicLit(x) {
		return
	}

	tv, ok := c.ctx.typesInfo.Types[x]
	if !ok {
		return
	}
	usageType, ok := tv.Type.(*types.Named)
	if !ok {
		return
	}

	// This linear search is not efficient, but does not
	// bite hard since only few expressions make it up to
	// this point. May consider building an index though.
	top := usageType.Obj().Pkg().Scope()
	for _, name := range top.Names() {
		cv, ok := top.Lookup(name).(*types.Const)
		if !ok {
			continue
		}
		defType, ok := cv.Type().(*types.Named)
		if !ok || defType.Obj() != usageType.Obj() {
			continue
		}
		if !constant.Compare(tv.Value, token.EQL, cv.Val()) {
			continue
		}
		// Current way to avoid false positives for definitions
		// themself is to compare positions.
		defPos := c.ctx.fileSet.Position(cv.Pos())
		usagePos := c.ctx.fileSet.Position(x.Pos())
		if defPos.Line != usagePos.Line || defPos.Filename != usagePos.Filename {
			c.warn(x, tv.Value, cv)
		}
	}
}

func (c *namedConstChecker) warn(cause ast.Node, v constant.Value, named *types.Const) {
	suggestion := named.Name()
	if named.Pkg() != c.ctx.pkg {
		suggestion = named.Pkg().Name() + "." + suggestion
	}
	c.ctx.Warn(cause, "use %s instead of %s", suggestion, v)
}
