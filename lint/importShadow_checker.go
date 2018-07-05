package lint

//! Detects when imported package names shadowed in assignments.
//
// @Before:
// // "path/filepath" is imported.
// func myFunc(filepath string) {
// }
//
// @After:
// func myFunc(filename string) {
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
		elem := v.Parent().Lookup(def.ID.Name)
		if elem == nil {
			continue
		}
		pkg, ok := elem.(*types.PkgName)
		if !ok {
			continue
		}
		c.warnImportShadowed(def.ID, def.ID.Name, pkg.Imported())
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
