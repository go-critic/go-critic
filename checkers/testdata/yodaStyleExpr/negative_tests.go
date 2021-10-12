package checker_test

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
	_ = 15 == foo
}
