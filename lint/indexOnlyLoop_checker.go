package lint

import (
	"fmt"
	"go/ast"
)

//! Detects for loops that can benefit from rewrite to range loop.
//
// Suggests to use for key, v := range container form.
//
// @Before:
// func closeFiles(files []*os.File) {
//     for i := range files {
//         if files[i] != nil {
//            files[i].Close()
//         }
//     }
// }
//
// @After:
// func closeFilesSuggester(files []*os.File) {
//     for _, f := range files {
//         if f != nil {
//             f.Close()
//         }
//     }
// }

func init() {
	addChecker(&indexOnlyLoopChecker{}, attrExperimental)
}

type indexOnlyLoopChecker struct {
	checkerBase
}

func (c *indexOnlyLoopChecker) VisitStmt(stmt ast.Stmt) {
	if s, ok := stmt.(*ast.RangeStmt); ok && s.Key != nil {

		var sXIdentObj *ast.Object
		switch sX := s.X.(type) {
		case *ast.Ident:
			sXIdentObj = sX.Obj
		case *ast.SliceExpr:
			sXIdent, ok := sX.X.(*ast.Ident)
			if !ok {
				return
			}
			sXIdentObj = sXIdent.Obj
		default:
			return
		}

		var typ ast.Expr
		switch sXFieldDecl := sXIdentObj.Decl.(type) {
		case *ast.Field: // go 1.10
			typ = sXFieldDecl.Type
		case *ast.ValueSpec: // go 1.11
			typ = sXFieldDecl.Type
		default:
			return
		}
		// sX should always be of *ast.ArrayType type cause we are in *ast.RangeStmt statement
		// and even for *ast.SliceExpr, sXIdentObj will point to underlying array which have *ast.ArrayType
		sxFieldType, ok := typ.(*ast.ArrayType)
		if !ok {
			return
		}
		if _, ok = sxFieldType.Elt.(*ast.StarExpr); !ok {
			return
		}
		sKey := s.Key.(*ast.Ident).Obj

		count := 0
		ast.Inspect(stmt, func(n ast.Node) bool {
			if iExpr, ok := n.(*ast.IndexExpr); ok {
				x := iExpr.X.(*ast.Ident).Obj
				key := iExpr.Index.(*ast.Ident).Obj
				if x == sXIdentObj && key == sKey {
					count++
				}
			}
			// stop DFS traverse if we found more than one usage
			return count < 2
		})
		if count > 1 {
			c.warn(stmt, fmt.Sprintf("for _, value := range %s", sXIdentObj.Name))
		}
	}
}

func (c *indexOnlyLoopChecker) warn(x ast.Node, suggestion string) {
	c.ctx.Warn(x, "key variable occurs more then once in the loop; consider using %s", suggestion)
}
