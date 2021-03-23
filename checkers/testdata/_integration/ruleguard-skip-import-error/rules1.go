// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func badLock(m dsl.Matcher) {
	m.Import("foo.bar")
	// DSL invalid predicate is intentional. This is for test purpose.
	// This is specifically used to test the '-failOnError' flag when a
	// DSL syntax error or import error is encountered while parsing ruleguard files.
	// Use case: a ruleguard file contains an import that may or may not exist depending
	// on the module being analyzed.
	m.Match(`$mu.RLock(); defer $mu.Unlock()`).
		Where(m["mu].Type.Implements("foo.bar.MyInterface")).
		Report(`maybe $mu.Lock() was intended?`)
}
