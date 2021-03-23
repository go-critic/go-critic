// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func badLock(m dsl.Matcher) {
	m.Match(`$mu.Lock(); defer $mu.RUnlock()`).Report(`maybe $mu.RLock() was intended?`)
	// DSL invalid predicate is intentional. This is for test purpose.
	// This is specifically used to test the '-failOnError' flag when a
	// DSL syntax error or import error is encountered while parsing ruleguard files.
	// Use case 1: suppose a directory contains multiple ruleguard rules.
	// One of these files contains a ruleguard DSL with a newer keyword which is not supported
	// by the installed version of go-critic + ruleguard.
	// Use case 2: a ruleguard file contains an import that may or may not exist depending
	// on the module being analyzed.
	m.Match(`$mu.RLock(); defer $mu.Unlock()`).
		Where(version == "1.17").
		Report(`maybe $mu.Lock() was intended?`)
}
