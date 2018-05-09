// Package astfilter provides composable filters and predicates for astutil package.
package astfilter

import (
	"go/ast"

	"golang.org/x/tools/go/ast/astutil"
)

// Or joins filters with OR operation and returns composed filter.
// If any of filters return true, composed filter returns true.
func Or(filters ...astutil.ApplyFunc) astutil.ApplyFunc {
	return func(cur *astutil.Cursor) bool {
		for _, f := range filters {
			if f(cur) {
				return true
			}
		}
		return false
	}
}

// FuncDecl reports whether current node is *ast.FuncDecl.
func FuncDecl(cur *astutil.Cursor) bool {
	_, ok := cur.Node().(*ast.FuncDecl)
	return ok
}

// Stmt reports whether current node is ast.Stmt.
func Stmt(cur *astutil.Cursor) bool {
	_, ok := cur.Node().(ast.Stmt)
	return ok
}
