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
