//go:build ignore
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

//doc:tags test
func stackedIf(m dsl.Matcher) {
	m.Match(`if $*_ { if $*_ { $*_ } }`).
		Report(`may be simplified to one if`)
}

//doc:tags style
func dynamicFmtString(m dsl.Matcher) {
	m.Match(`fmt.Errorf($f($*args))`).
		Suggest("errors.New($f($args))").
		Report(`use errors.New($f($args)) or fmt.Errorf("%s", $f($args)) instead`)
}

//doc:tags experimental
func regexpCompile(m dsl.Matcher) {
	m.Match(`regexp.Match($*_)`,
		`regexp.MatchString($*_)`,
		`regexp.MatchReader($*_)`,
	).Report(`regexp compilation should be avoided on the hot paths`)
}
