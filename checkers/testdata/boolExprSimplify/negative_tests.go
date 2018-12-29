package checker_test

func idealExpressions() {
	// No negations:
	_ = true
	_ = true
}

func negationOK() {
	_ = !true
	_ = !false
}

func otherBinOps() {
	_ = true != false
	_ = true && false
	_ = true || false
}

func cantCombine() {
	fn := func() int { return 0 }

	var x, y, z int

	// OK: not safe expressions.
	_ = fn() > y || fn() == y
	_ = fn() == y || fn() > y
	_ = x > fn() || x == fn()
	_ = x == fn() || x > fn()
	_ = fn() > fn() || fn() == fn()
	_ = fn() == fn() || fn() > fn()

	// OK: different operands.
	_ = x > y || x == z
	_ = x == z || x > y
	_ = x > z || x == y
	_ = x == y || x > z

	// OK: unrelated operations.
	_ = x < y || x > z
	_ = x > z || x < y
}

func floatCompare() {
	var f1, f2 float32

	// Can't be simplified to `f1 != f2`.
	_ = !(f1 == f2)
}
