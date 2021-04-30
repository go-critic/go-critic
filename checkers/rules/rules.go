package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

//doc:summary Detects nil usages in http.NewRequest calls, suggesting http.NoBody as an alternative
//doc:tags    style experimental
//doc:before  http.NewRequest("GET", url, nil)
//doc:after   http.NewRequest("GET", url, http.NoBody)
func httpNoBody(m dsl.Matcher) {
	m.Match("http.NewRequest($method, $url, $nil)").
		Where(m["nil"].Text == "nil").
		Suggest("http.NewRequest($method, $url, http.NoBody)").
		Report("http.NoBody should be preferred to the nil request body")
}

//doc:summary Detects expressions like []rune(s)[0] that may cause unwanted rune slice allocation
//doc:tags    performance experimental
//doc:before  r := []rune(s)[0]
//doc:after   r, _ := utf8.DecodeRuneInString(s)
//doc:note    See Go issue for details: https://github.com/golang/go/issues/45260
func preferDecodeRune(m dsl.Matcher) {
	m.Match(`[]rune($s)[0]`).
		Where(m["s"].Type.Is(`string`)).
		Report(`consider replacing $$ with utf8.DecodeRuneInString($s)`)
}
