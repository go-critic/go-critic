package lint

import (
	"go/ast"
	"go/types"
)

//! Detects for loops that can benefit from rewrite to range loop.
//
// Suggests to use for key, v := range container form.
//
// @Before:
// for i := range files {
// 	if files[i] != nil {
// 		files[i].Close()
// 	}
// }
//
// @After:
// for _, f := range files {
// 	if f != nil {
// 		f.Close()
// 	}
// }

func init() {
	addChecker(&indexOnlyLoopChecker{}, attrExperimental)
}

type indexOnlyLoopChecker struct {
	checkerBase
}

func (c *indexOnlyLoopChecker) VisitStmt(stmt ast.Stmt) {
	rng, ok := stmt.(*ast.RangeStmt)
	if !ok || rng.Key == nil {
		return
	}
	iterated := c.ctx.typesInfo.ObjectOf(identOf(rng.X))
	if iterated == nil || !c.elemTypeIsPtr(iterated) {
		return // To avoid redundant traversals
	}
	count := 0
	ast.Inspect(rng.Body, func(n ast.Node) bool {
		if n, ok := n.(*ast.IndexExpr); ok {
			if iterated == c.ctx.typesInfo.ObjectOf(identOf(n.X)) {
				count++
			}
		}
		// Stop DFS traverse if we found more than one usage.
		return count < 2
	})
	if count > 1 {
		c.warn(stmt, rng.Key, iterated.Name())
	}
}

func (c *indexOnlyLoopChecker) elemTypeIsPtr(obj types.Object) bool {
	switch typ := obj.Type().(type) {
	case *types.Slice:
		return typeIsPointer(typ.Elem())
	case *types.Array:
		return typeIsPointer(typ.Elem())
	default:
		return false
	}
}

func (c *indexOnlyLoopChecker) warn(x, key ast.Node, iterated string) {
	c.ctx.Warn(x, "%s occurs more than once in the loop; consider using for _, value := range %s",
		key, iterated)
}
