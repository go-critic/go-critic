package checker_test

func foo() float32 {
	return 9.0
}

func f1() bool {
	var p float32 = 10.0
	var q float32 = 5.0
	/// change `(p + q) == foo()` to `math.Abs((p + q) - foo()) < eps`
	return (p + q) == foo()
}

func f2() bool {
	var a float32 = 10.0
	var b float32 = 5.0
	/// change `2*a+4*b == 0.5` to `math.Abs(2*a + 4*b - 0.5) < eps`
	return 2*a+4*b == 0.5
}

func f3() bool {
	var a float32 = 10.0
	var c float32 = 5.0
	/// change `c == (a + 5)` to `math.Abs(c - (a + 5)) < eps`
	return c == (a + 5)
}

func f4() bool {
	var a float32 = 10.0
	/// change `foo()+foo() == a+a` to `math.Abs(foo() + foo() - (a + a)) < eps`
	return foo()+foo() == a+a
}
