package checker_test

func f1() bool {
	var a float32 = 10.0
	var b int = 5.0
	/// consider to change way to compare floats in expression to math.Abs(a - float32(b)) < eps
	return a == float32(b)
}

func f2() bool {
	var a float32 = 10.0
	var b int = 5.0
	/// consider to change way to compare floats in expression to math.Abs(a - float32(b)) >= eps
	return a != float32(b)
}

func f3() bool {
	var a float32 = 10.0
	var b int = 5.0
	/// consider to change way to compare floats in expression to math.Abs(float32(a) - float32(b)) >= eps
	return float32(a) != float32(b)
}

func f4() {
	var a, b float64 = 10.0, 20.0

	/// consider to change way to compare floats in expression to math.Abs(a - b) < eps
	_ = a == b

	/// consider to change way to compare floats in expression to math.Abs(a - b) < eps
	_ = !(a == b)

	/// consider to change way to compare floats in expression to math.Abs(a - (b + a)) < eps
	_ = a == (b + a)

	/// consider to change way to compare floats in expression to math.Abs(a - 40.0) >= eps
	_ = a != 40.0

	/// consider to change way to compare floats in expression to math.Abs(a * 2 - b) < eps
	_ = a*2 == b

	/// consider to change way to compare floats in expression to math.Abs(a - (b + 4)) < eps
	_ = a == (b + 4)

	/// consider to change way to compare floats in expression to math.Abs(a - b) < eps
	/// consider to change way to compare floats in expression to math.Abs(b - a) >= eps
	_ = a == b && b != a

	/// consider to change way to compare floats in expression to math.Abs(a - b) < eps
	/// consider to change way to compare floats in expression to math.Abs(b - a) >= eps
	/// consider to change way to compare floats in expression to math.Abs(a - b) >= eps
	_ = a == b && b != a || !(a != b)

}
