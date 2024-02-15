package gorules

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

//doc:summary Detects redundant fmt.Sprint calls
//doc:tags    style experimental
//doc:before  fmt.Sprint(x)
//doc:after   x.String()
func redundantSprint(m dsl.Matcher) {
	m.Match(`fmt.Sprint($x)`, `fmt.Sprintf("%s", $x)`, `fmt.Sprintf("%v", $x)`).
		Where(!m["x"].Type.Is(`reflect.Value`) && m["x"].Type.Implements(`fmt.Stringer`)).
		Suggest(`$x.String()`).
		Report(`use $x.String() instead`)

	m.Match(`fmt.Sprint($x)`, `fmt.Sprintf("%s", $x)`, `fmt.Sprintf("%v", $x)`).
		Where(m["x"].Type.Is(`string`)).
		Suggest(`$x`).
		Report(`$x is already string`)
}

//doc:summary Detects deferred function literals that can be simplified
//doc:tags    style experimental
//doc:before  defer func() { f() }()
//doc:after   defer f()
func deferUnlambda(m dsl.Matcher) {
	m.Match(`defer func() { $f($*args) }()`).
		Where(m["f"].Node.Is(`Ident`) && m["f"].Text != "panic" && m["f"].Text != "recover" && m["args"].Const).
		Report("can rewrite as `defer $f($args)`")

	m.Match(`defer func() { $pkg.$f($*args) }()`).
		Where(m["f"].Node.Is(`Ident`) && m["args"].Const && m["pkg"].Object.Is(`PkgName`)).
		Report("can rewrite as `defer $pkg.$f($args)`")
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

	m.Match("http.NewRequestWithContext($ctx, $method, $url, $nil)").
		Where(m["nil"].Text == "nil").
		Suggest("http.NewRequestWithContext($ctx, $method, $url, http.NoBody)").
		Report("http.NoBody should be preferred to the nil request body")

	m.Match("httptest.NewRequest($method, $url, $nil)").
		Where(m["nil"].Text == "nil").
		Suggest("httptest.NewRequest($method, $url, http.NoBody)").
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
//doc:tags    diagnostic
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
	m.Match(`len($s) > 0`).
		Where(m["s"].Type.Is(`string`)).
		Report("replace `$$` with `$s != \"\"`")

	m.Match(`len($s) == 0`).
		Where(m["s"].Type.Is(`string`)).
		Report("replace `$$` with `$s == \"\"`")
	m.Match(`len($s) <= 0`).
		Where(m["s"].Type.Is(`string`)).
		Report("replace `$$` with `$s == \"\"`")
}

//doc:summary Detects redundant conversions between string and []byte
//doc:tags    performance
//doc:before  copy(b, []byte(s))
//doc:after   copy(b, s)
func stringXbytes(m dsl.Matcher) {
	m.Match(`copy($_, []byte($s))`).Report("can simplify `[]byte($s)` to `$s`")

	m.Match(`string($b) == ""`).Where(m["b"].Type.Is(`[]byte`)).Suggest(`len($b) == 0`)
	m.Match(`string($b) != ""`).Where(m["b"].Type.Is(`[]byte`)).Suggest(`len($b) != 0`)

	m.Match(`len(string($b))`).Where(m["b"].Type.Is(`[]byte`)).Suggest(`len($b)`)

	m.Match(`string($x) == string($y)`).
		Where(m["x"].Type.Is(`[]byte`) && m["y"].Type.Is(`[]byte`)).
		Suggest(`bytes.Equal($x, $y)`)

	m.Match(`string($x) != string($y)`).
		Where(m["x"].Type.Is(`[]byte`) && m["y"].Type.Is(`[]byte`)).
		Suggest(`!bytes.Equal($x, $y)`)

	m.Match(`$re.Match([]byte($s))`).
		Where(m["re"].Type.Is(`*regexp.Regexp`) && m["s"].Type.Is(`string`)).
		Suggest(`$re.MatchString($s)`)

	m.Match(`$re.FindIndex([]byte($s))`).
		Where(m["re"].Type.Is(`*regexp.Regexp`) && m["s"].Type.Is(`string`)).
		Suggest(`$re.FindStringIndex($s)`)

	m.Match(`$re.FindAllIndex([]byte($s), $n)`).
		Where(m["re"].Type.Is(`*regexp.Regexp`) && m["s"].Type.Is(`string`)).
		Suggest(`$re.FindAllStringIndex($s, $n)`)
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
	m.Match(`strings.Index($s1, $s2) >= 0`, `strings.Index($s1, $s2) != -1`).Suggest(`strings.Contains($s1, $s2)`)
	m.Match(`strings.IndexAny($s1, $s2) >= 0`, `strings.IndexAny($s1, $s2) != -1`).Suggest(`strings.ContainsAny($s1, $s2)`)
	m.Match(`strings.IndexRune($s1, $s2) >= 0`, `strings.IndexRune($s1, $s2) != -1`).Suggest(`strings.ContainsRune($s1, $s2)`)

	m.Match(`$i := strings.Index($s, $sep); $*_; $x, $y = $s[:$i], $s[$i+1:]`,
		`$i := strings.Index($s, $sep); $*_; $x = $s[:$i]; $*_; $y = $s[$i+1:]`).
		Where(m.GoVersion().GreaterEqThan("1.18")).
		Suggest("$x, $y, _ = strings.Cut($s, $sep)")

	m.Match(
		`if $i := strings.Index($s, $sep); $i != -1 { $*_; $x, $y = $s[:$i], $s[$i+1:]; $*_ }`,
		`if $i := strings.Index($s, $sep); $i != -1 { $*_; $x = $s[:$i]; $*_; $y = $s[$i+1:]; $*_ }`,
		`if $i := strings.Index($s, $sep); $i >= 0 { $*_; $x, $y = $s[:$i], $s[$i+1:]; $*_ }`,
		`if $i := strings.Index($s, $sep); $i >= 0 { $*_; $x = $s[:$i]; $*_; $y = $s[$i+1:]; $*_ }`).
		Where(m.GoVersion().GreaterEqThan("1.18")).
		Suggest("if $x, $y, ok = strings.Cut($s, $sep); ok { ... }")

	m.Match(`bytes.SplitN(b, []byte("."), -1)`).Report("use bytes.Split method in `$$`")
	m.Match(`bytes.Replace($_, $_, $_, -1)`).Report("use bytes.ReplaceAll method in `$$`")
	m.Match(`bytes.Map(unicode.ToUpper, $_)`).Report("use bytes.ToUpper method in `$$`")
	m.Match(`bytes.Map(unicode.ToLower, $_)`).Report("use bytes.ToLower method in `$$`")
	m.Match(`bytes.Map(unicode.ToTitle, $_)`).Report("use bytes.ToTitle method in `$$`")
	m.Match(`bytes.Index($b1, $b2) >= 0`, `bytes.Index($b1, $b2) != -1`).Suggest(`bytes.Contains($b1, $b2)`)
	m.Match(`bytes.IndexAny($b1, $b2) >= 0`, `bytes.IndexAny($b1, $b2) != -1`).Suggest(`bytes.ContainsAny($b1, $b2)`)
	m.Match(`bytes.IndexRune($b1, $b2) >= 0`, `bytes.IndexRune($b1, $b2) != -1`).Suggest(`bytes.ContainsRune($b1, $b2)`)

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

	m.Match(`filepath.Join($_)`).Report(`suspicious Join on 1 argument`)
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

//doc:summary Detects WriteRune calls with rune literal argument that is single byte and reports to use WriteByte instead
//doc:tags    performance experimental opinionated
//doc:before  w.WriteRune('\n')
//doc:after   w.WriteByte('\n')
func preferWriteByte(m dsl.Matcher) {
	// utf8.RuneSelf:
	// characters below RuneSelf are represented as themselves in a single byte.
	const runeSelf = 0x80
	m.Match(`$w.WriteRune($c)`).Where(
		m["w"].Type.Implements("io.ByteWriter") && (m["c"].Const && m["c"].Value.Int() < runeSelf),
	).Report(`consider writing single byte rune $c with $w.WriteByte($c)`)
}

//doc:summary Detects fmt.Sprint(f/ln) calls which can be replaced with fmt.Fprint(f/ln)
//doc:tags    performance experimental
//doc:before  w.Write([]byte(fmt.Sprintf("%x", 10)))
//doc:after   fmt.Fprintf(w, "%x", 10)
func preferFprint(m dsl.Matcher) {
	m.Match(`$w.Write([]byte(fmt.Sprint($*args)))`).
		Where(m["w"].Type.Implements("io.Writer")).
		Suggest("fmt.Fprint($w, $args)").
		Report(`fmt.Fprint($w, $args) should be preferred to the $$`)

	m.Match(`$w.Write([]byte(fmt.Sprintf($*args)))`).
		Where(m["w"].Type.Implements("io.Writer")).
		Suggest("fmt.Fprintf($w, $args)").
		Report(`fmt.Fprintf($w, $args) should be preferred to the $$`)

	m.Match(`$w.Write([]byte(fmt.Sprintln($*args)))`).
		Where(m["w"].Type.Implements("io.Writer")).
		Suggest("fmt.Fprintln($w, $args)").
		Report(`fmt.Fprintln($w, $args) should be preferred to the $$`)

	m.Match(`io.WriteString($w, fmt.Sprint($*args))`).Suggest(`fmt.Fprint($w, $args)`)
	m.Match(`io.WriteString($w, fmt.Sprintf($*args))`).Suggest(`fmt.Fprintf($w, $args)`)
	m.Match(`io.WriteString($w, fmt.Sprintln($*args))`).Suggest(`fmt.Fprintln($w, $args)`)

	m.Match("$w.WriteString(fmt.Sprint($*args))").
		Where(m["w"].Type.Implements("io.Writer") && m["w"].Type.Implements("io.StringWriter")).
		Suggest("fmt.Fprint($w, $args)")
	m.Match("$w.WriteString(fmt.Sprintf($*args))").
		Where(m["w"].Type.Implements("io.Writer") && m["w"].Type.Implements("io.StringWriter")).
		Suggest("fmt.Fprintf($w, $args)")
	m.Match("$w.WriteString(fmt.Sprintln($*args))").
		Where(m["w"].Type.Implements("io.Writer") && m["w"].Type.Implements("io.StringWriter")).
		Suggest(`fmt.Fprintln($w, $args)`)
}

//doc:summary Detects suspicious duplicated arguments
//doc:tags    diagnostic
//doc:before  copy(dst, dst)
//doc:after   copy(dst, src)
func dupArg(m dsl.Matcher) {
	m.Match(`$x.Equal($x)`, `$x.Equals($x)`, `$x.Compare($x)`, `$x.Cmp($x)`).
		Where(m["x"].Pure).
		Report(`suspicious method call with the same argument and receiver`)

	m.Match(`copy($x, $x)`,
		`math.Max($x, $x)`,
		`math.Min($x, $x)`,
		`reflect.Copy($x, $x)`,
		`reflect.DeepEqual($x, $x)`,
		`strings.Contains($x, $x)`,
		`strings.Compare($x, $x)`,
		`strings.EqualFold($x, $x)`,
		`strings.HasPrefix($x, $x)`,
		`strings.HasSuffix($x, $x)`,
		`strings.Index($x, $x)`,
		`strings.LastIndex($x, $x)`,
		`strings.Split($x, $x)`,
		`strings.SplitAfter($x, $x)`,
		`strings.SplitAfterN($x, $x, $_)`,
		`strings.SplitN($x, $x, $_)`,
		`strings.Replace($_, $x, $x, $_)`,
		`strings.ReplaceAll($_, $x, $x)`,
		`bytes.Contains($x, $x)`,
		`bytes.Compare($x, $x)`,
		`bytes.Equal($x, $x)`,
		`bytes.EqualFold($x, $x)`,
		`bytes.HasPrefix($x, $x)`,
		`bytes.HasSuffix($x, $x)`,
		`bytes.Index($x, $x)`,
		`bytes.LastIndex($x, $x)`,
		`bytes.Split($x, $x)`,
		`bytes.SplitAfter($x, $x)`,
		`bytes.SplitAfterN($x, $x, $_)`,
		`bytes.SplitN($x, $x, $_)`,
		`bytes.Replace($_, $x, $x, $_)`,
		`bytes.ReplaceAll($_, $x, $x)`,
		`types.Identical($x, $x)`,
		`types.IdenticalIgnoreTags($x, $x)`,
		`draw.Draw($x, $_, $x, $_, $_)`).
		Where(m["x"].Pure).
		Report(`suspicious duplicated args in $$`)
}

//doc:summary Detects suspicious http.Error call without following return
//doc:tags    diagnostic experimental
//doc:before  if err != nil { http.Error(...); }
//doc:after   if err != nil { http.Error(...); return; }
func returnAfterHttpError(m dsl.Matcher) {
	m.Match(`if $_ { $*_; http.Error($w, $err, $code) }`).
		Report("Possibly return is missed after the http.Error call").
		At(m["w"])
}

//doc:summary Detects concatenation with os.PathSeparator which can be replaced with filepath.Join
//doc:tags    style experimental
//doc:before  x + string(os.PathSeparator) + y
//doc:after   filepath.Join(x, y)
func preferFilepathJoin(m dsl.Matcher) {
	m.Match(`$x + string(os.PathSeparator) + $y`).
		Where(m["x"].Type.Is(`string`) && m["y"].Type.Is(`string`)).
		Suggest("filepath.Join($x, $y)").
		Report(`filepath.Join($x, $y) should be preferred to the $$`)
}

//doc:summary Detects w.Write or io.WriteString calls which can be replaced with w.WriteString
//doc:tags    performance experimental
//doc:before  w.Write([]byte("foo"))
//doc:after   w.WriteString("foo")
func preferStringWriter(m dsl.Matcher) {
	m.Match(`$w.Write([]byte($s))`).
		Where(m["w"].Type.Implements("io.StringWriter")).
		Suggest("$w.WriteString($s)").
		Report(`$w.WriteString($s) should be preferred to the $$`)

	m.Match(`io.WriteString($w, $s)`).
		Where(m["w"].Type.Implements("io.StringWriter")).
		Suggest("$w.WriteString($s)").
		Report(`$w.WriteString($s) should be preferred to the $$`)
}

//doc:summary Detects slice clear loops, suggests an idiom that is recognized by the Go compiler
//doc:tags    performance experimental
//doc:before  for i := 0; i < len(buf); i++ { buf[i] = 0 }
//doc:after   for i := range buf { buf[i] = 0 }
func sliceClear(m dsl.Matcher) {
	m.Match(`for $i := 0; $i < len($xs); $i++ { $xs[$i] = $zero }`).
		Where(m["zero"].Value.Int() == 0).
		Report(`rewrite as for-range so compiler can recognize this pattern`)
}

//doc:summary Detects sync.Map load+delete operations that can be replaced with LoadAndDelete
//doc:tags    diagnostic experimental
//doc:before  v, ok := m.Load(k); if ok { m.Delete($k); f(v); }
//doc:after   v, deleted := m.LoadAndDelete(k); if deleted { f(v) }
func syncMapLoadAndDelete(m dsl.Matcher) {
	m.Match(`$_, $ok := $m.Load($k); if $ok { $m.Delete($k); $*_ }`).
		Where(m.GoVersion().GreaterEqThan("1.15") &&
			m["m"].Type.Is(`*sync.Map`)).
		Report(`use $m.LoadAndDelete to perform load+delete operations atomically`)
}

//doc:summary Detects "%s" formatting directives that can be replaced with %q
//doc:tags    diagnostic experimental
//doc:before  fmt.Sprintf(`"%s"`, s)
//doc:after   fmt.Sprintf(`%q`, s)
func sprintfQuotedString(m dsl.Matcher) {
	m.Match(`fmt.Sprintf($s, $*_)`).
		Where(m["s"].Text.Matches("^`.*\"%s\".*`$") ||
			m["s"].Text.Matches(`^".*\\"%s\\".*"$`)).
		Report(`use %q instead of "%s" for quoted strings`)
}

//doc:summary Detects various off-by-one kind of errors
//doc:tags    diagnostic
//doc:before  xs[len(xs)]
//doc:after   xs[len(xs)-1]
func offBy1(m dsl.Matcher) {
	m.Match(`$x[len($x)]`).
		Where(m["x"].Pure && m["x"].Type.Is(`[]$_`)).
		Suggest(`$x[len($x)-1]`).
		Report(`index expr always panics; maybe you wanted $x[len($x)-1]?`)

	// TODO: use $slicing[$i:$*_] form when we'll update go-ruleguard
	// version so it includes https://github.com/quasilyte/go-ruleguard/pull/284

	m.Match(
		`$i := strings.Index($s, $_); $_ := $slicing[$i:]`,
		`$i := strings.Index($s, $_); $_ = $slicing[$i:]`,
		`$i := bytes.Index($s, $_); $_ := $slicing[$i:]`,
		`$i := bytes.Index($s, $_); $_ = $slicing[$i:]`).
		Where(m["s"].Text == m["slicing"].Text).
		Report(`Index() can return -1; maybe you wanted to do $s[$i+1:]`).
		At(m["slicing"])

	m.Match(
		`$i := strings.Index($s, $_); $_ := $slicing[:$i]`,
		`$i := strings.Index($s, $_); $_ = $slicing[:$i]`,
		`$i := bytes.Index($s, $_); $_ := $slicing[:$i]`,
		`$i := bytes.Index($s, $_); $_ = $slicing[:$i]`).
		Where(m["s"].Text == m["slicing"].Text).
		Report(`Index() can return -1; maybe you wanted to do $s[:$i+1]`).
		At(m["slicing"])

	m.Match(
		`$s[strings.Index($s, $_):]`,
		`$s[:strings.Index($s, $_)]`,
		`$s[bytes.Index($s, $_):]`,
		`$s[:bytes.Index($s, $_)]`).
		Report(`Index() can return -1; maybe you wanted to do Index()+1`)
}

//doc:summary Detects slice expressions that can be simplified to sliced expression itself
//doc:tags    style
//doc:before  copy(b[:], values...)
//doc:after   copy(b, values...)
func unslice(m dsl.Matcher) {
	m.Match(`$s[:]`).
		Where(m["s"].Type.Is(`string`) || m["s"].Type.Is(`[]$_`)).
		Suggest(`$s`).
		Report(`could simplify $$ to $s`)
}

//doc:summary Detects Yoda style expressions and suggests to replace them
//doc:tags    style experimental
//doc:before  return nil != ptr
//doc:after   return ptr != nil
func yodaStyleExpr(m dsl.Matcher) {
	m.Match(`$constval != $x`).Where(m["constval"].Node.Is(`BasicLit`) && !m["x"].Node.Is(`BasicLit`)).
		Report(`consider to change order in expression to $x != $constval`)
	m.Match(`$constval == $x`).Where(m["constval"].Node.Is(`BasicLit`) && !m["x"].Node.Is(`BasicLit`)).
		Report(`consider to change order in expression to $x == $constval`)

	m.Match(`nil != $x`).Where(!m["x"].Node.Is(`BasicLit`)).
		Report(`consider to change order in expression to $x != nil`)
	m.Match(`nil == $x`).Where(!m["x"].Node.Is(`BasicLit`)).
		Report(`consider to change order in expression to $x == nil`)
}

//doc:summary Detects unoptimal strings/bytes case-insensitive comparison
//doc:tags    performance experimental
//doc:before  strings.ToLower(x) == strings.ToLower(y)
//doc:after   strings.EqualFold(x, y)
func equalFold(m dsl.Matcher) {
	// We specify so many patterns to avoid too generic
	// patterns that would match things like
	// `strings.ToLower(x) == strings.ToUpper(y)`
	// While it could be an EqualFold candidate,
	// it just looks wrong and should probably be
	// marked by some other checker.

	// string == patterns
	m.Match(
		`strings.ToLower($x) == $y`,
		`strings.ToLower($x) == strings.ToLower($y)`,
		`$x == strings.ToLower($y)`,
		`strings.ToUpper($x) == $y`,
		`strings.ToUpper($x) == strings.ToUpper($y)`,
		`$x == strings.ToUpper($y)`).
		Where(m["x"].Pure && m["y"].Pure && m["x"].Text != m["y"].Text).
		Suggest(`strings.EqualFold($x, $y)`).
		Report(`consider replacing with strings.EqualFold($x, $y)`)

	// string != patterns
	m.Match(
		`strings.ToLower($x) != $y`,
		`strings.ToLower($x) != strings.ToLower($y)`,
		`$x != strings.ToLower($y)`,
		`strings.ToUpper($x) != $y`,
		`strings.ToUpper($x) != strings.ToUpper($y)`,
		`$x != strings.ToUpper($y)`).
		Where(m["x"].Pure && m["y"].Pure && m["x"].Text != m["y"].Text).
		Suggest(`!strings.EqualFold($x, $y)`).
		Report(`consider replacing with !strings.EqualFold($x, $y)`)

	// bytes.Equal patterns
	m.Match(
		`bytes.Equal(bytes.ToLower($x), $y)`,
		`bytes.Equal(bytes.ToLower($x), bytes.ToLower($y))`,
		`bytes.Equal($x, bytes.ToLower($y))`,
		`bytes.Equal(bytes.ToUpper($x), $y)`,
		`bytes.Equal(bytes.ToUpper($x), bytes.ToUpper($y))`,
		`bytes.Equal($x, bytes.ToUpper($y))`).
		Where(m["x"].Pure && m["y"].Pure && m["x"].Text != m["y"].Text).
		Suggest(`bytes.EqualFold($x, $y)`).
		Report(`consider replacing with bytes.EqualFold($x, $y)`)
}

//doc:summary Detects suspicious arguments order
//doc:tags    diagnostic
//doc:before  strings.HasPrefix("#", userpass)
//doc:after   strings.HasPrefix(userpass, "#")
func argOrder(m dsl.Matcher) {
	m.Match(
		`strings.HasPrefix($lit, $s)`,
		`bytes.HasPrefix($lit, $s)`,
		`strings.HasSuffix($lit, $s)`,
		`bytes.HasSuffix($lit, $s)`,
		`strings.Contains($lit, $s)`,
		`bytes.Contains($lit, $s)`,
		`strings.TrimPrefix($lit, $s)`,
		`bytes.TrimPrefix($lit, $s)`,
		`strings.TrimSuffix($lit, $s)`,
		`bytes.TrimSuffix($lit, $s)`,
		`strings.Split($lit, $s)`,
		`bytes.Split($lit, $s)`).
		Where((m["lit"].Const || m["lit"].ConstSlice) &&
			!(m["s"].Const || m["s"].ConstSlice) &&
			!m["lit"].Node.Is(`Ident`)).
		Report(`$lit and $s arguments order looks reversed`)
}

//doc:summary Detects string concat operations that can be simplified
//doc:tags    style experimental
//doc:before  strings.Join([]string{x, y}, "_")
//doc:after   x + "_" + y
func stringConcatSimplify(m dsl.Matcher) {
	m.Match(`strings.Join([]string{$x, $y}, "")`).Suggest(`$x + $y`)
	m.Match(`strings.Join([]string{$x, $y, $z}, "")`).Suggest(`$x + $y + $z`)
	m.Match(`strings.Join([]string{$x, $y}, $glue)`).Suggest(`$x + $glue + $y`)
}

//doc:summary Detects manual conversion to milli- or microseconds
//doc:tags    style experimental
//doc:before  t.Unix() / 1000
//doc:after   t.UnixMilli()
func timeExprSimplify(m dsl.Matcher) {
	isTime := func(v dsl.Var) bool {
		return v.Type.Is(`time.Time`) || v.Type.Is(`*time.Time`)
	}

	m.Match(`$t.Unix() / 1000`).
		Where(m.GoVersion().GreaterEqThan("1.17") && isTime(m["t"])).
		Suggest("$t.UnixMilli()").
		Report(`use $t.UnixMilli() instead of $$`)

	m.Match(`$t.UnixNano() * 1000`).
		Where(m.GoVersion().GreaterEqThan("1.17") && isTime(m["t"])).
		Suggest("$t.UnixMicro()").
		Report(`use $t.UnixMicro() instead of $$`)
}

//doc:summary Detects exposed methods from sync.Mutex and sync.RWMutex
//doc:tags    style experimental
//doc:before  type Foo struct{ ...; sync.Mutex; ... }
//doc:after   type Foo struct{ ...; mu sync.Mutex; ... }
func exposedSyncMutex(m dsl.Matcher) {
	isExported := func(v dsl.Var) bool {
		return v.Text.Matches(`^\p{Lu}`)
	}

	m.Match(`type $x struct { $*_; sync.Mutex; $*_ }`).
		Where(isExported(m["x"])).
		Report("don't embed sync.Mutex")

	m.Match(`type $x struct { $*_; *sync.Mutex; $*_ }`).
		Where(isExported(m["x"])).
		Report("don't embed *sync.Mutex")

	m.Match(`type $x struct { $*_; sync.RWMutex; $*_ }`).
		Where(isExported(m["x"])).
		Report("don't embed sync.RWMutex")

	m.Match(`type $x struct { $*_; *sync.RWMutex; $*_ }`).
		Where(isExported(m["x"])).
		Report("don't embed *sync.RWMutex")
}

//doc:summary Detects bad usage of sort package
//doc:tags    diagnostic experimental
//doc:before  xs = sort.StringSlice(xs)
//doc:after   sort.Strings(xs)
func badSorting(m dsl.Matcher) {
	m.Match(`$x = sort.IntSlice($x)`).
		Where(m["x"].Type.Is(`[]int`)).
		Suggest(`sort.Ints($x)`).
		Report(`suspicious sort.IntSlice usage, maybe sort.Ints was intended?`)

	m.Match(`$x = sort.Float64Slice($x)`).
		Where(m["x"].Type.Is(`[]float64`)).
		Suggest(`sort.Float64s($x)`).
		Report(`suspicious sort.Float64s usage, maybe sort.Float64s was intended?`)

	m.Match(`$x = sort.StringSlice($x)`).
		Where(m["x"].Type.Is(`[]string`)).
		Suggest(`sort.Strings($x)`).
		Report(`suspicious sort.StringSlice usage, maybe sort.Strings was intended?`)
}

//doc:summary Detects suspicious reassignment of error from another package
//doc:tags    diagnostic experimental
//doc:before  io.EOF = nil
//doc:after   /* don't do it */
func externalErrorReassign(m dsl.Matcher) {
	m.Match(`$pkg.$err = $x`).
		Where(m["err"].Type.Is(`error`) && m["pkg"].Object.Is(`PkgName`)).
		Report(`suspicious reassignment of error from another package`)
}

//doc:summary Detects suspicious empty declarations blocks
//doc:tags    diagnostic experimental
//doc:before  var()
//doc:after   /* nothing */
func emptyDecl(m dsl.Matcher) {
	m.Match(`var()`).Report(`empty var() block`)
	m.Match(`const()`).Report(`empty const() block`)
	m.Match(`type()`).Report(`empty type() block`)
}

//doc:summary Detects suspicious formatting strings usage
//doc:tags    diagnostic experimental
//doc:before  fmt.Errorf(msg)
//doc:after   fmt.Errorf("%s", msg)
func dynamicFmtString(m dsl.Matcher) {
	m.Match(`fmt.Errorf($f)`).
		Where(!m["f"].Const).
		Suggest("errors.New($f)").
		Report(`use errors.New($f) or fmt.Errorf("%s", $f) instead`)

	m.Match(`fmt.Errorf($f($*args))`).
		Suggest("errors.New($f($*args))").
		Report(`use errors.New($f($*args)) or fmt.Errorf("%s", $f($*args)) instead`)
}

//doc:summary Detects strings.Compare usage
//doc:tags    style experimental
//doc:before  strings.Compare(x, y)
//doc:after   x < y
func stringsCompare(m dsl.Matcher) {
	m.Match(`strings.Compare($s1, $s2) == 0`).
		Suggest(`$s1 == $s2`)

	m.Match(`strings.Compare($s1, $s2) == -1`,
		`strings.Compare($s1, $s2) < 0`).
		Suggest(`$s1 < $s2`)

	m.Match(`strings.Compare($s1, $s2) == 1`,
		`strings.Compare($s1, $s2) > 0`).
		Suggest(`$s1 > $s2`)
}

//doc:summary Detects unchecked errors in if statements
//doc:tags    diagnostic experimental
//doc:before  if err := expr(); err2 != nil { /*...*/ }
//doc:after   if err := expr(); err != nil { /*...*/ }
func uncheckedInlineErr(m dsl.Matcher) {
	m.Match(
		`if $err := $_($*_); $err2 != nil { $*_ }`,
		`if $err = $_($*_); $err2 != nil { $*_ }`,
		`if $*_, $err := $_($*_); $err2 != nil { $*_ }`,
		`if $*_, $err = $_($*_); $err2 != nil { $*_ }`).
		Where(m["err"].Type.Implements("error") && m["err2"].Type.Implements("error") &&
			m["err"].Text != m["err2"].Text).
		Report("$err error is unchecked, maybe intended to check it instead of $err2").
		At(m["err"])
}

//doc:summary Detects bad usage of sync.OnceFunc
//doc:tags    diagnostic experimental
//doc:before  sync.OnceFunc(foo)()
//doc:after   fooOnce := sync.OnceFunc(foo); ...; fooOnce()
func badSyncOnceFunc(m dsl.Matcher) {
	m.Match(`$*_; sync.OnceFunc($x); $*_;`).
		Report("possible sync.OnceFunc misuse, sync.OnceFunc($x) result is not used").
		Where(m.GoVersion().GreaterEqThan("1.21"))

	m.Match(`sync.OnceFunc($x)()`).
		Report("possible sync.OnceFunc misuse, consider to assign sync.OnceFunc($x) to a variable").
		Where(m.GoVersion().GreaterEqThan("1.21"))
}
