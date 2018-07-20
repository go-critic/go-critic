package lintutil

import (
	"go/ast"

	"github.com/go-toolsmith/astequal"
)

type astMapItem struct {
	key ast.Node
	val interface{}
}

// AstMap is a simple ast.Node map.
// Zero value is ready to use set.
// Can be reused after Clear call.
// In a simplest case, can be used as AST set.
type AstMap struct {
	items []astMapItem
}

// Contains reports whether m contains key.
func (m *AstMap) Contains(key ast.Node) bool {
	// Can't use Find because it's fine to store nil val
	// for set-like usage of map.
	for i := range m.items {
		if astequal.Node(m.items[i].key, key) {
			return true
		}
	}
	return false
}

// Find returns value associated with key or nil if there are none.
func (m *AstMap) Find(key ast.Node) interface{} {
	for i := range m.items {
		if astequal.Node(m.items[i].key, key) {
			return m.items[i].val
		}
	}
	return nil
}

// ValueAt returns value stored at index.
// Index must be non-negative and less than m.Len().
func (m *AstMap) ValueAt(index int) interface{} {
	return m.items[index].val
}

// KeyAt returns key stored at index.
// Index must be non-negative and less than m.Len().
func (m *AstMap) KeyAt(index int) ast.Node {
	return m.items[index].key
}

// Insert pushes <key, val> in m if key is not already there.
// Returns true if element was inserted.
// If only set-like behavior is needed, one may use nil as a value.
func (m *AstMap) Insert(key ast.Node, val interface{}) bool {
	if m.Contains(key) {
		return false
	}
	m.items = append(m.items, astMapItem{key: key, val: val})
	return true
}

// Clear removes all element from map.
func (m *AstMap) Clear() {
	m.items = m.items[:0]
}

// Len returns the number of elements contained inside m.
func (m *AstMap) Len() int {
	return len(m.items)
}
