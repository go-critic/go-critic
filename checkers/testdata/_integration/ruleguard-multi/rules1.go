// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func badLock(m dsl.Matcher) {
	m.Match(`$mu.Lock(); defer $mu.RUnlock()`).Report(`maybe $mu.RLock() was intended?`)
	m.Match(`$mu.RLock(); defer $mu.Unlock()`).Report(`maybe $mu.Lock() was intended?`)
}
