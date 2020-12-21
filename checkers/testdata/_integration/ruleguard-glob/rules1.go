// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl/fluent"

func badLock(m fluent.Matcher) {
	m.Match(`$mu.Lock(); defer $mu.RUnlock()`).Report(`maybe $mu.RLock() was intended?`)
	m.Match(`$mu.RLock(); defer $mu.Unlock()`).Report(`maybe $mu.Lock() was intended?`)
}
