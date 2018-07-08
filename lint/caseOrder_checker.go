package lint

import (
	"go/ast"
	"go/types"
)

func init() {
	addChecker(&caseOrderChecker{}, attrExperimental)
}

type caseOrderChecker struct {
	checkerBase
}

func (c *caseOrderChecker) InitDocs(d *Documentation) {
	d.Summary = "Detects erroneous case order inside switch statements"
	d.Before = `
switch x.(type) {
case ast.Expr:
	fmt.Println("expr")
case *ast.BasicLit:
	fmt.Println("basic lit") // Never executed
}`
	d.After = `
switch x.(type) {
case *ast.BasicLit:
	fmt.Println("basic lit") // Now reachable
case ast.Expr:
	fmt.Println("expr")
}`
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
		name string
		typ  *types.Interface
	}
	var ifaces []ifaceType // Interfaces seen so far
	for _, cc := range s.Body.List {
		cc := cc.(*ast.CaseClause)
		for _, x := range cc.List {
			typ := c.ctx.typesInfo.TypeOf(x)
			for _, iface := range ifaces {
				if types.Implements(typ, iface.typ) {
					c.warnTypeSwitch(cc, typ, iface.name)
					break
				}
			}
			if iface, ok := typ.Underlying().(*types.Interface); ok {
				// Need to retrive object to get proper interface name.
				name := "interface{}"
				if obj := c.ctx.typesInfo.ObjectOf(identOf(x)); obj != nil {
					name = obj.Name()
				}
				ifaces = append(ifaces, ifaceType{name: name, typ: iface})
			}
		}
	}
}

func (c *caseOrderChecker) warnTypeSwitch(cause ast.Node, typ types.Type, iface string) {
	concrete := types.TypeString(typ, func(p *types.Package) string {
		if p == nil || p == c.ctx.pkg {
			return ""
		}
		return p.Name()
	})
	c.ctx.Warn(cause, "case %s must go before the %s case", concrete, iface)
}

func (c *caseOrderChecker) checkSwitch(s *ast.SwitchStmt) {
	// TODO(quasilyte): can handle expression cases that overlap.
	// Cases that have narrower value range should go before wider ones.
}
