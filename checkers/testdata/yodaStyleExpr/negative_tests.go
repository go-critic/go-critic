package checker_test

import "unsafe"

func g1() {
	var a int
	if a == 10 {
	}
}

func g2() bool {
	f := func() interface{} { return nil }
	return f() != nil
}

func _() {
	const foo = 10
	_ = 10 != 15
	_ = foo == 15

	type myArray struct {
		data [10]int
	}
	var arr myArray
	var i int
	_ = len(arr.data) == i
	_ = i == len(arr.data)

	_ = unsafe.Sizeof(0) == 0

	var c byte
	if '0' <= c && c <= '9' {
		// character range ok
	}
	if c >= '0' && c <= '9' {
		// character range ok
	}
}
