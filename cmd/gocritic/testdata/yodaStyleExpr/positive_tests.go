package checker_test

func f1() {
	var m map[int]int
	/// consider to change order of expression in nil == m
	if nil == m {
	}

	var a int
	/// consider to change order of expression in 10 == a
	if 10 == a {
	}

	var s string
	/// consider to change order of expression in "" == s
	if "" == s {
	}
}

func f2() bool {
	var ch chan int
	switch {
	/// consider to change order of expression in nil == ch
	case nil == ch:
		//
	}
	/// consider to change order of expression in nil == ch
	return nil == ch
}

type kek struct {
	a int
}

func f3() {
	var k kek
	/// consider to change order of expression in 0 == k.a
	if 0 == k.a {
	}
}

func f4() {
	var a int
	f := func(bool) {}
	/// consider to change order of expression in 10 == a
	f(10 == a)
}
