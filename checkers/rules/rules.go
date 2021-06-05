package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

//doc:summary Detects deprecated io/ioutil package usages
//doc:tags    style experimental
//doc:before  ioutil.ReadAll(r)
//doc:after   io.ReadAll(r)
func ioutilDeprecated(m dsl.Matcher) {
	m.Match(`ioutil.ReadAll($_)`).
		Report(`ioutil.ReadAll is deprecated, use io.ReadAll instead`)

	m.Match(`ioutil.ReadFile($_)`).
		Report(`ioutil.ReadFile is deprecated, use os.ReadFile instead`)

	m.Match(`ioutil.WriteFile($_, $_, $_)`).
		Report(`ioutil.WriteFile is deprecated, use os.WriteFile instead`)

	m.Match(`ioutil.ReadDir($_)`).
		Report(`ioutil.ReadDir is deprecated, use os.ReadDir instead`)

	m.Match(`ioutil.NopCloser($_)`).
		Report(`ioutil.NopCloser is deprecated, use io.NopCloser instead`)

	m.Match(`ioutil.Discard`).
		Report(`ioutil.NopCloser is deprecated, use io.Discard instead`)
}

//doc:summary Detects suspicious mutex lock/unlock operations
//doc:tags    diagnostic experimental
//doc:before  mu.Lock(); mu.Unlock()
//doc:after   mu.Lock(); defer mu.Unlock()
func badLock(m dsl.Matcher) {
	// `mu1` and `mu2` are added to make possible report a line where `m2` is used (with a defer)

	// no defer
	m.Match(`$mu1.Lock(); $mu2.Unlock()`).
		Where(m["mu1"].Text == m["mu2"].Text).
		Report(`defer is missing, mutex is unlocked immediately`).
		At(m["mu2"])

	m.Match(`$mu1.RLock(); $mu2.RUnlock()`).
		Where(m["mu1"].Text == m["mu2"].Text).
		Report(`defer is missing, mutex is unlocked immediately`).
		At(m["mu2"])

	// different lock operations
	m.Match(`$mu1.Lock(); defer $mu2.RUnlock()`).
		Where(m["mu1"].Text == m["mu2"].Text).
		Report(`suspicious unlock, maybe Unlock was intended?`).
		At(m["mu2"])

	m.Match(`$mu1.RLock(); defer $mu2.Unlock()`).
		Where(m["mu1"].Text == m["mu2"].Text).
		Report(`suspicious unlock, maybe RUnlock was intended?`).
		At(m["mu2"])

	// double locks
	m.Match(`$mu1.Lock(); defer $mu2.Lock()`).
		Where(m["mu1"].Text == m["mu2"].Text).
		Report(`maybe defer $mu1.Unlock() was intended?`).
		At(m["mu2"])

	m.Match(`$mu1.RLock(); defer $mu2.RLock()`).
		Where(m["mu1"].Text == m["mu2"].Text).
		Report(`maybe defer $mu1.RUnlock() was intended?`).
		At(m["mu2"])
}

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

//doc:summary Detects usage of `len` when result is obvious or doesn't make sense
//doc:tags    style
//doc:before  len(arr) <= 0
//doc:after   len(arr) == 0
func sloppyLen(m dsl.Matcher) {
	m.Match(`len($_) >= 0`).Report(`$$ is always true`)
	m.Match(`len($_) < 0`).Report(`$$ is always false`)
	m.Match(`len($x) <= 0`).Report(`$$ can be len($x) == 0`)
}

//doc:summary Detects value swapping code that are not using parallel assignment
//doc:tags    style
//doc:before  *tmp = *x; *x = *y; *y = *tmp
//doc:after   *x, *y = *y, *x
func valSwap(m dsl.Matcher) {
	m.Match(`$tmp := $y; $y = $x; $x = $tmp`).
		Report("can re-write as `$y, $x = $x, $y`")
}

//doc:summary Detects switch-over-bool statements that use explicit `true` tag value
//doc:tags    style
//doc:before  switch true {...}
//doc:after   switch {...}
func switchTrue(m dsl.Matcher) {
	m.Match(`switch true { $*_ }`).
		Report(`replace 'switch true {}' with 'switch {}'`)
	m.Match(`switch $x; true { $*_ }`).
		Report(`replace 'switch $x; true {}' with 'switch $x; {}'`)
}

//doc:summary Detects immediate dereferencing of `flag` package pointers
//doc:tags    diagnostic
//doc:before  b := *flag.Bool("b", false, "b docs")
//doc:after   var b bool; flag.BoolVar(&b, "b", false, "b docs")
func flagDeref(m dsl.Matcher) {
	m.Match(`*flag.Bool($*_)`).Report(`immediate deref in $$ is most likely an error; consider using flag.BoolVar`)
	m.Match(`*flag.Duration($*_)`).Report(`immediate deref in $$ is most likely an error; consider using flag.DurationVar`)
	m.Match(`*flag.Float64($*_)`).Report(`immediate deref in $$ is most likely an error; consider using flag.Float64Var`)
	m.Match(`*flag.Int($*_)`).Report(`immediate deref in $$ is most likely an error; consider using flag.IntVar`)
	m.Match(`*flag.Int64($*_)`).Report(`immediate deref in $$ is most likely an error; consider using flag.Int64Var`)
	m.Match(`*flag.String($*_)`).Report(`immediate deref in $$ is most likely an error; consider using flag.StringVar`)
	m.Match(`*flag.Uint($*_)`).Report(`immediate deref in $$ is most likely an error; consider using flag.UintVar`)
	m.Match(`*flag.Uint64($*_)`).Report(`immediate deref in $$ is most likely an error; consider using flag.Uint64Var`)
}

//doc:summary Detects empty string checks that can be written more idiomatically
//doc:tags    style experimental
//doc:before  len(s) == 0
//doc:after   s == ""
func emptyStringTest(m dsl.Matcher) {
	m.Match(`len($s) != 0`).
		Where(m["s"].Type.Is(`string`)).
		Report("replace `$$` with `$s != \"\"`")

	m.Match(`len($s) == 0`).
		Where(m["s"].Type.Is(`string`)).
		Report("replace `$$` with `$s == \"\"`")
}

//doc:summary Detects redundant conversions between string and []byte
//doc:tags    style
//doc:before  copy(b, []byte(s))
//doc:after   copy(b, s)
func stringXbytes(m dsl.Matcher) {
	m.Match(`copy($_, []byte($s))`).Report("can simplify `[]byte($s)` to `$s`")
}

//doc:summary Detects strings.Index calls that may cause unwanted allocs
//doc:tags    performance
//doc:before  strings.Index(string(x), y)
//doc:after   bytes.Index(x, []byte(y))
//doc:note    See Go issue for details: https://github.com/golang/go/issues/25864
func indexAlloc(m dsl.Matcher) {
	m.Match(`strings.Index(string($x), $y)`).
		Where(m["x"].Pure && m["y"].Pure).
		Report(`consider replacing $$ with bytes.Index($x, []byte($y))`)
}

//doc:summary Detects function calls that can be replaced with convenience wrappers
//doc:tags    style
//doc:before  wg.Add(-1)
//doc:after   wg.Done()
func wrapperFunc(m dsl.Matcher) {
	m.Match(`$wg.Add(-1)`).
		Where(m["wg"].Type.Is(`sync.WaitGroup`)).
		Report("use WaitGroup.Done method in `$$`")

	m.Match(`$buf.Truncate(0)`).
		Where(m["buf"].Type.Is(`bytes.Buffer`)).
		Report("use Buffer.Reset method in `$$`")

	m.Match(`http.HandlerFunc(http.NotFound)`).Report("use http.NotFoundHandler method in `$$`")

	m.Match(`strings.SplitN($_, $_, -1)`).Report("use strings.Split method in `$$`")
	m.Match(`strings.Replace($_, $_, $_, -1)`).Report("use strings.ReplaceAll method in `$$`")
	m.Match(`strings.Map(unicode.ToTitle, $_)`).Report("use strings.ToTitle method in `$$`")

	m.Match(`bytes.SplitN(b, []byte("."), -1)`).Report("use bytes.Split method in `$$`")
	m.Match(`bytes.Replace($_, $_, $_, -1)`).Report("use bytes.ReplaceAll method in `$$`")
	m.Match(`bytes.Map(unicode.ToUpper, $_)`).Report("use bytes.ToUpper method in `$$`")
	m.Match(`bytes.Map(unicode.ToLower, $_)`).Report("use bytes.ToLower method in `$$`")
	m.Match(`bytes.Map(unicode.ToTitle, $_)`).Report("use bytes.ToTitle method in `$$`")

	m.Match(`draw.DrawMask($_, $_, $_, $_, nil, image.Point{}, $_)`).
		Report("use draw.Draw method in `$$`")
}

//doc:summary Detects `regexp.Compile*` that can be replaced with `regexp.MustCompile*`
//doc:tags    style
//doc:before  re, _ := regexp.Compile("const pattern")
//doc:after   re := regexp.MustCompile("const pattern")
func regexpMust(m dsl.Matcher) {
	m.Match(`regexp.Compile($pat)`).
		Where(m["pat"].Const).
		Report(`for const patterns like $pat, use regexp.MustCompile`)

	m.Match(`regexp.CompilePOSIX($pat)`).
		Where(m["pat"].Const).
		Report(`for const patterns like $pat, use regexp.MustCompilePOSIX`)
}

//doc:summary Detects suspicious function calls
//doc:tags    diagnostic
//doc:before  strings.Replace(s, from, to, 0)
//doc:after   strings.Replace(s, from, to, -1)
func badCall(m dsl.Matcher) {
	m.Match(`strings.Replace($_, $_, $_, $zero)`).
		Where(m["zero"].Value.Int() == 0).
		Report(`suspicious arg 0, probably meant -1`).At(m["zero"])
	m.Match(`bytes.Replace($_, $_, $_, $zero)`).
		Where(m["zero"].Value.Int() == 0).
		Report(`suspicious arg 0, probably meant -1`).At(m["zero"])

	m.Match(`strings.SplitN($_, $_, $zero)`).
		Where(m["zero"].Value.Int() == 0).
		Report(`suspicious arg 0, probably meant -1`).At(m["zero"])
	m.Match(`bytes.SplitN($_, $_, $zero)`).
		Where(m["zero"].Value.Int() == 0).
		Report(`suspicious arg 0, probably meant -1`).At(m["zero"])

	m.Match(`append($_)`).Report(`no-op append call, probably missing arguments`)
}

//doc:summary Detects assignments that can be simplified by using assignment operators
//doc:tags    style
//doc:before  x = x * 2
//doc:after   x *= 2
func assignOp(m dsl.Matcher) {
	m.Match(`$x = $x + 1`).Where(m["x"].Pure).Report("replace `$$` with `$x++`")
	m.Match(`$x = $x - 1`).Where(m["x"].Pure).Report("replace `$$` with `$x--`")

	m.Match(`$x = $x + $y`).Where(m["x"].Pure).Report("replace `$$` with `$x += $y`")
	m.Match(`$x = $x - $y`).Where(m["x"].Pure).Report("replace `$$` with `$x -= $y`")

	m.Match(`$x = $x * $y`).Where(m["x"].Pure).Report("replace `$$` with `$x *= $y`")
	m.Match(`$x = $x / $y`).Where(m["x"].Pure).Report("replace `$$` with `$x /= $y`")
	m.Match(`$x = $x % $y`).Where(m["x"].Pure).Report("replace `$$` with `$x %= $y`")
	m.Match(`$x = $x & $y`).Where(m["x"].Pure).Report("replace `$$` with `$x &= $y`")
	m.Match(`$x = $x | $y`).Where(m["x"].Pure).Report("replace `$$` with `$x |= $y`")
	m.Match(`$x = $x ^ $y`).Where(m["x"].Pure).Report("replace `$$` with `$x ^= $y`")
	m.Match(`$x = $x << $y`).Where(m["x"].Pure).Report("replace `$$` with `$x <<= $y`")
	m.Match(`$x = $x >> $y`).Where(m["x"].Pure).Report("replace `$$` with `$x >>= $y`")
	m.Match(`$x = $x &^ $y`).Where(m["x"].Pure).Report("replace `$$` with `$x &^= $y`")
}
