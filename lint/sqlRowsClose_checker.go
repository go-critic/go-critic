package lint

import (
	"go/ast"
	"go/types"
)

const (
	rowsTypePTR = "*database/sql.Rows"
)

func init() {
	addChecker(&sqlRowsCloseChecker{}, attrExperimental)
}

type sqlRowsCloseChecker struct {
	checkerBase
}

func (c *sqlRowsCloseChecker) InitDocumentation(d *Documentation) {
	d.Summary = "Detects uses of *sql.Rows without call Close method"
	d.Before = `
rows, _ := db.Query( /**/ )
for rows.Next {
}`
	d.After = `
rows, _ := db.Query( /**/ )
for rows.Next {
}
rows.Close()`
}

// Warning if sql.Rows local variables (including function parameters):
// 1. Not using as parameter in other functions call;
// 2. Not returning in functions results;
// 3. Not call Close method for variable;

func (c *sqlRowsCloseChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	localVars := make([]types.Object, 0)
	returnVars := make([]types.Object, 0)
	closeVars := make([]types.Object, 0)
	for _, b := range decl.Body.List {
		switch b := b.(type) {
		case *ast.AssignStmt:
			// Detect local vars with sql.Rows types
			if b.Lhs != nil {
				for _, l := range b.Lhs {
					if c.typeString(l) == rowsTypePTR {
						localVars = append(localVars, c.ctx.typesInfo.ObjectOf(identOf(l)))
					}
				}
			}
		case *ast.ReturnStmt:
			// Detect return vars with sql.Rows types
			if b.Results != nil && len(b.Results) > 0 {
				for _, r := range b.Results {
					t := c.ctx.typesInfo.TypeOf(r)
					if t.String() == rowsTypePTR {
						returnVars = append(returnVars, c.ctx.typesInfo.ObjectOf(identOf(r)))
					}
				}
			}
		case *ast.ExprStmt:
			if b, ok := b.X.(*ast.CallExpr); ok {
				if bb, ok := b.Fun.(*ast.SelectorExpr); ok {
					// Detect call Close for sql.Rows variables
					funcName := qualifiedName(bb.Sel)
					if funcName == "Close" {
						closeVars = append(closeVars, c.ctx.typesInfo.ObjectOf(identOf(bb.X)))
					}
				}
			}
		case *ast.DeferStmt:
			// Detect call Close for sql.Rows variables over defer declaration
			if b, ok := b.Call.Fun.(*ast.SelectorExpr); ok {
				funcName := qualifiedName(b.Sel)
				if funcName == "Close" {
					closeVars = append(closeVars, c.ctx.typesInfo.ObjectOf(identOf(b.X)))
				}
			}
		}
	}

	// Check local variables
	for _, l := range localVars {
		// If local variable present in return or Close present - PASS
		if !c.varInList(l, returnVars) && !c.varInList(l, closeVars) {
			c.ctx.Warn(l.Parent(), "local variable db.Rows have not Close call")
		}
	}
}

func (c *sqlRowsCloseChecker) typeString(x ast.Expr) string {
	if typ := c.ctx.typesInfo.TypeOf(x); typ != nil {
		return typ.String()
	}
	return ""
}

func (c *sqlRowsCloseChecker) varInList(v types.Object, list []types.Object) bool {
	for _, r := range list {
		if v == r {
			return true
		}
	}
	return false
}
