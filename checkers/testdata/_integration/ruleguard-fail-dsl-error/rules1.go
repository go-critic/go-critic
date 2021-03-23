// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func badLock(m dsl.Matcher) {
	m.Match(`$mu.Lock(); defer $mu.RUnlock()`).Report(`maybe $mu.RLock() was intended?`)
	// DSL invalid predicate is intentional. This is for test purpose.
	// In this test, -@ruleguard.failOnError=dsl is specified, so go-critic should
	// return with non-zero status.
	// The file contains a DSL syntax error, -@ruleguard.failOnError=dsl specifies to exit with error.
	// Use case: suppose a directory contains multiple ruleguard rules.
	// One of these files contains a ruleguard DSL with a newer keyword which is not supported
	// by the installed version of go-critic + ruleguard.
	m.Match(`$mu.RLock(); defer $mu.Unlock()`).
		Where(version == "1.17").
		Report(`maybe $mu.Lock() was intended?`)
}
