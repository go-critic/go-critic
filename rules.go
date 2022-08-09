//nolint // it's not a Go code file
package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

func nilSafeTypeOf(m dsl.Matcher) {
	m.Match(`$_.ctx.TypesInfo.TypeOf($x)`).
		Report(`use ctx.TypeOf($x) instead, it's nil-safe`)
}

func stdSizeof(m dsl.Matcher) {
	m.Match(`$_.ctx.SizesInfo.Sizeof($_)`).
		Report(`use ctx.SizeOf instead`)
}
