package lint

import (
	"fmt"
	"go/ast"
)

//! Detects when range loop is iterating over container which has pointer-type elements (that is, any *T or reference-like type as map or chan) using only key variable (no value) and container[key] occurs more than once inside loop body
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
		sX := s.X.(*ast.Ident).Obj
		// sX is always declaration of some *ast.ArrayType
		// cause we are in *ast.RangeStmt statement
		sXField := sX.Decl.(*ast.Field)
		sxFiledType := sXField.Type.(*ast.ArrayType)
		if _, ok = sxFiledType.Elt.(*ast.StarExpr); !ok {
			return
		}
		sKey := s.Key.(*ast.Ident).Obj

		count := 0
		ast.Inspect(stmt, func(n ast.Node) bool {
			if iExpr, ok := n.(*ast.IndexExpr); ok {
				x := iExpr.X.(*ast.Ident).Obj
				key := iExpr.Index.(*ast.Ident).Obj
				if x == sX && key == sKey {
					count++
				}
			}
			// stop DFS traverse if we found more then one usage
			return count < 2
		})
		if count > 1 {
			c.warn(stmt, fmt.Sprintf("for _, value := range %s", sX.Name))
		}
	}
}

func (c *indexOnlyLoopChecker) warn(x ast.Node, suggestion string) {
	c.ctx.Warn(x, "key variable occurs more then once in the loop; consider using %s", suggestion)
}
