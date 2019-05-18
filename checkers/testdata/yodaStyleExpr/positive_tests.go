package checker_test

func yodaComparisons() {
	var x int
	/*! consider to change order in expression to x <= 0 */
	_ = 0 > x
	/*! consider to change order in expression to x >= 0 */
	_ = 0 < x
	/*! consider to change order in expression to x < 0 */
	_ = 0 >= x
	/*! consider to change order in expression to x > 0 */
	_ = 0 <= x
}

func f1() {
	var m map[int]int
	/*! consider to change order in expression to m == nil */
	if nil == m {
	}

	var a int
	/*! consider to change order in expression to a == 10 */
	if 10 == a {
	}

	var s string
	/*! consider to change order in expression to s == "" */
	if "" == s {
	}
}

func f2() bool {
	var ch chan int
	switch {
	/*! consider to change order in expression to ch == nil */
	case nil == ch:
		//
	}
	/*! consider to change order in expression to ch == nil */
	return nil == ch
}

type foo struct {
	a int
}

func f3() {
	var k foo
	/*! consider to change order in expression to k.a == 0 */
	if 0 == k.a {
	}
}

func f4() {
	var a int
	f := func(bool) {}
	/*! consider to change order in expression to a == 10 */
	f(10 == a)
}

func f5() bool {
	f := func() interface{} { return nil }
	/*! consider to change order in expression to f() != nil */
	return nil != f()
}
