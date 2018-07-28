package checker_test

func g1() bool {
	var a, b float64 = 10.0, 20.0
	return a < b
}

func g2() bool {
	var a, b int = 10, 20
	return a == b
}

func g3() {
	var a, b float64 = 10.0, 20.0

	_ = a >= b

	_ = a - b

	_ = b < a && a >= b
}
