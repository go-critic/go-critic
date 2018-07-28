package checker_test

func f1() bool {
	var a float32 = 10.0
	var b int = 5.0
	/// change `a == float32(b)` to `math.Abs(a - float32(b)) < eps`
	return a == float32(b)
}

func f2() bool {
	var a float32 = 10.0
	var b int = 5.0
	/// change `a != float32(b)` to `math.Abs(a - float32(b)) >= eps`
	return a != float32(b)
}

func f3() bool {
	var a float32 = 10.0
	var b int = 5.0
	/// change `float32(a) != float32(b)` to `math.Abs(float32(a) - float32(b)) >= eps`
	return float32(a) != float32(b)
}

func f4(a, b float64) {

	/// change `a == b` to `math.Abs(a - b) < eps`
	_ = a == b

	/// change `a == b` to `math.Abs(a - b) < eps`
	_ = !(a == b)

	/// change `a == (b + a)` to `math.Abs(a - (b + a)) < eps`
	_ = a == (b + a)

	/// change `a != 40.0` to `math.Abs(a - 40.0) >= eps`
	_ = a != 40.0

	/// change `a*2 == b` to `math.Abs(a * 2 - b) < eps`
	_ = a*2 == b

	/// change `a == (b + 4)` to `math.Abs(a - (b + 4)) < eps`
	_ = a == (b + 4)

	/// change `a == b` to `math.Abs(a - b) < eps`
	/// change `b != a` to `math.Abs(b - a) >= eps`
	_ = a == b && b != a

	/// change `a == b` to `math.Abs(a - b) < eps`
	/// change `b != a` to `math.Abs(b - a) >= eps`
	/// change `a != b` to `math.Abs(a - b) >= eps`
	_ = a == b && b != a || !(a != b)

	// TODO: change `a == b+a` to `math.Abs(a - b - a) < eps`
	/// change `a == b+a` to `math.Abs(a - b + a) < eps`
	_ = a == b+a
}
