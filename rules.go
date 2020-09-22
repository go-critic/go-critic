//nolint // it's not a Go code file
package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl/fluent"
)

func nilSafeTypeOf(m fluent.Matcher) {
	m.Match(`$_.ctx.TypesInfo.TypeOf($x)`).
		Report(`use ctx.TypeOf($x) instead, it's nil-safe`)
}
