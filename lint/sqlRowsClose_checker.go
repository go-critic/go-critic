package lint

//! Detects use *sql.Rows without call Close method.
//
// @Before:
// rows, _ := db.Query( /**/ )
// for rows.Next {
// }
//
// @After:
// rows, _ := db.Query( /**/ )
// for rows.Next {
// }
// rows.Close()

import (
	"go/ast"
	"go/types"
)

const (
	rowsTypePTR = "*database/sql.Rows"
	rowsType    = "database/sql.Rows"
)

func init() {
	addChecker(&sqlRowsCloseChecker{}, attrExperimental)
}

type sqlRowsCloseChecker struct {
	checkerBase
}

// Warning if sql.Rows local variables (including function parameters):
// 1. Not using as parameter in other functions call;
// 2. Not returning in functions results;
// 3. Not call Close method for variable;

func (c *sqlRowsCloseChecker) VisitFuncDecl(decl *ast.FuncDecl) {
	// Function parameter variable
	params := decl.Type.Params
	var paramVal types.Object
	for _, p := range params.List {
		t := c.ctx.typesInfo.TypeOf(p.Type)
		if t != nil && t.String() == rowsTypePTR {
			paramVal = c.ctx.typesInfo.ObjectOf(identOf(p.Names[0]))
			break
		}
	}

	localVars := make([]types.Object, 0)
	returnVars := make([]types.Object, 0)
	closeVars := make([]types.Object, 0)
	callVars := make([]types.Object, 0)
	for _, b := range decl.Body.List {
		switch b := b.(type) {
		case *ast.AssignStmt:
			// Detect local vars with sql.Rows types
			if b.Lhs != nil {
				for _, l := range b.Lhs {
					t := c.ctx.typesInfo.TypeOf(l)
					if t.String() == rowsTypePTR {
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
				} else {
					// Detect call other functions with sql.Rows variable in parameters
					for _, v := range b.Args {
						t := c.ctx.typesInfo.TypeOf(v)
						if t.String() == rowsTypePTR || t.String() == rowsType {
							callVars = append(callVars, c.ctx.typesInfo.ObjectOf(identOf(v)))
						}
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

	// CHECKS

	// Check function parameter local variable
	if paramVal != nil {
		// If parameter variable present in return or in other functions call or Close present - PASS
		if !c.varInList(paramVal, returnVars) &&
			!c.varInList(paramVal, callVars) &&
			!c.varInList(paramVal, closeVars) {
			c.ctx.Warn(paramVal.Parent(), "param variable db.Rows have not Close call")
		}
	}

	// Check local variables
	for _, l := range localVars {
		// If local variable present in return or in other functions call or Close present - PASS
		if !c.varInList(l, returnVars) &&
			!c.varInList(l, callVars) &&
			!c.varInList(l, closeVars) {
			c.ctx.Warn(l.Parent(), "local variable db.Rows have not Close call")
		}
	}
}

func (c *sqlRowsCloseChecker) varInList(v types.Object, list []types.Object) bool {
	for _, r := range list {
		if v == r {
			return true
		}
	}
	return false
}
