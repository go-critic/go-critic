// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func badLock(m dsl.Matcher) {
	m.Match(`$mu.Lock(); defer $mu.RUnlock()`).Report(`maybe $mu.RLock() was intended?`)
	// DSL invalid predicate is intentional. This is for test purpose.
	// In this test, -@ruleguard.failOnError=dsl is NOT specified, so go-critic should
	// ignore this file.
	// I.e. skip files that contain a DSL syntax error.
	//
	// Use case: suppose a directory contains multiple ruleguard rules.
	// One of these files contains a ruleguard DSL with a newer keyword which is not supported
	// by the installed version of go-critic + ruleguard.
	m.Match(`$mu.RLock(); defer $mu.Unlock()`).
		Where(version == "1.17").
		Report(`maybe $mu.Lock() was intended?`)
}
