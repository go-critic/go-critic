package checker_test

func warn() {
	/*! empty slice should be declared as var x []int */
	x := []int{}
	x = append(x, 1)

	/*! empty slice should be declared as var y []int */
	y := make([]int, 0)
	y = append(y, 1)

	/*! empty slice should be declared as var z []int */
	z := make([]int, 0)
	z = append(z, 1)
}
