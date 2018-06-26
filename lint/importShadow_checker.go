package lint

//! Detects when imported package names shadowed in assignments.
//
// @Before:
// import (
//     "fmt"
//     "math"
// )
// func main() {
//     fmt.Println(math.Pi)
//     // shadowing math package
//     math := 10
//     fmt.Println(len)
// }
//
// @After:
// import (
//     "fmt"
//     "math"
// )
// func main() {
//     fmt.Println(math.Pi)
//     // change identificator name
//     value := 10
//     fmt.Println(value)
// }

import (
	"go/ast"
	"go/types"

	"github.com/go-critic/go-critic/lint/internal/astwalk"
)

func init() {
	addChecker(&importShadowChecker{}, attrExperimental)
}

type importShadowChecker struct {
	checkerBase
}

func (c *importShadowChecker) VisitLocalDef(def astwalk.Name, _ ast.Expr) {
	for _, v := range c.ctx.Context.typesInfo.Defs {
		if v == nil || v.Parent() == nil {
			continue
		}
		for _, name := range v.Parent().Names() {
			pkg, ok := v.Parent().Lookup(name).(*types.PkgName)
			if !ok {
				continue
			}
			if name == def.ID.Name {
				c.warnImportShadowed(def.ID, name, pkg.Imported())
			}
		}
	}
}

func (c *importShadowChecker) warnImportShadowed(id ast.Node, importedName string, pkg *types.Package) {
	if pkg.Path() == pkg.Name() {
		// check for standart library packages
		c.ctx.Warn(id, "shadow of imported package '%s'", importedName)
	} else {
		c.ctx.Warn(id, "shadow of imported from '%s' package '%s'", pkg.Path(), importedName)
	}
}
