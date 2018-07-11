package checker_test

func f() {
	a := []int{}

	/// useless len comparison len(a) >= 0, always true
	_ = len(a) >= 0
	/// useless len comparison len(a) < 0, always false
	_ = len(a) < 0
	/// useless len comparison len(a) <= 0, can be len(a) == 0
	_ = len(a) <= 0
}
