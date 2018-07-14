package checker_test

func g() {
	a := []int{}

	_ = len(a) > 0
	_ = len(a) >= 10
	_ = len(a) <= 10
	_ = len(a) == 0
	_ = len(a) == 10
}
