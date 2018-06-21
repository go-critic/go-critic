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
	for _, imp := range c.ctx.Context.pkg.Imports() {
		if imp.Name() == def.ID.Name {
			c.warnImportShadowed(def.ID, imp)
		}
	}
}

func (c *importShadowChecker) warnImportShadowed(id ast.Node, pkg *types.Package) {
	if pkg.Path() == pkg.Name() {
		// check for standart library packages
		c.ctx.Warn(id, "shadow of imported package '%s'", pkg.Name())
	} else {
		c.ctx.Warn(id, "shadow of imported from '%s' package '%s'", pkg.Path(), pkg.Name())
	}
}
