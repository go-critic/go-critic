package checker_test

func f() {
	a := []int{}

	/*! len(a) >= 0 is always true */
	_ = len(a) >= 0
	/*! len(a) < 0 is always false */
	_ = len(a) < 0
	/*! len(a) <= 0 can be len(a) == 0 */
	_ = len(a) <= 0
}
