// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func errorUnderlying(m dsl.Matcher) {
	m.Match(`type $x error`).
		Report(`error as an underlying type is probably a mistake`).
		Suggest(`type $x struct { error }`)
}

func sprintfConcat(m dsl.Matcher) {
	m.Match(`fmt.Sprintf("%s%s", $a, $b)`).
		Where(m["a"].Type.Is(`string`) && m["b"].Type.Is(`string`)).
		Suggest(`$a+$b`)
}
