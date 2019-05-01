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

	_ = x < 11 || x > 14
	_ = x <= 10 || x >= z
	_ = x <= 10 || x >= 100

	_ = x < 11 || y > 14
	_ = x <= 10 || y >= z
	_ = x <= 10 || y >= 100

	_ = x < 11 || fn() > 14
	_ = fn() <= 10 || x >= z
	_ = fn() <= 10 || fn() >= 100
}

func floatCompare() {
	var f1, f2 float32

	// Can't be simplified to `f1 != f2`.
	_ = !(f1 == f2)

	// Can't be simplified to `p == 1`.
	_ = f1 > 0 && f1 <= 1
}

func balancedIncDec(x, y, z int) {
	// This is usually done on purpose.

	_ = x+0 < y+0
	_ = x+1 < y+1

	_ = x-0 <= y-0
	_ = x-1 <= y-1

	_ = x-0 < y-0
	_ = x-1 < y-1
}
