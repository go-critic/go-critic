package checker_test

func lenAnd(xs, ys []int) {
	var i int

	_ = len(xs) != 0 && xs[0] == 10
	_ = (len(xs) >= 1) && xs[1] == 10
	_ = len(xs) != 0 && add1(xs[0]) == 0
	_ = len(xs) > 1 && (xs[0]+xs[1]) != 0
	_ = len(xs) > i && xs[i] > 10
}

func lenOr(xs, ys []int) {
	var i int

	_ = len(xs) == 0 || xs[0] == 10
	_ = (len(xs) < 1) || xs[1] == 10
	_ = len(xs) == 0 || add1(xs[0]) == 0
	_ = len(xs) >= 1 || (xs[0]+xs[1]) != 0
	_ = len(xs) <= i || xs[i] > 10
}
