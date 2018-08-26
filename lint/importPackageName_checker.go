package lint

import (
	"go/ast"
	"strings"
)

func init() {
	addChecker(&importPackageNameChecker{}, attrExperimental)
}

type importPackageNameChecker struct {
	checkerBase
}

func (c *importPackageNameChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects when imported package names are unnecessary renamed"
	d.Before = `
import lint "github.com/go-critic/go-critic/lint"`
	d.After = `
import "github.com/go-critic/go-critic/lint"`
}

func (c *importPackageNameChecker) VisitFile(file *ast.File) {
	for _, imp := range file.Imports {
		var pkgName string
		for _, pkgImport := range c.ctx.pkg.Imports() {
			if pkgImport.Path() == strings.Trim(imp.Path.Value, `"`) {
				pkgName = pkgImport.Name()
				break
			}
		}

		if imp.Name != nil && imp.Name.Name == pkgName {
			c.warn(imp)
		}
	}
}

func (c *importPackageNameChecker) warn(cause ast.Node) {
	c.ctx.Warn(cause, "unnecessary rename of import package")
}
