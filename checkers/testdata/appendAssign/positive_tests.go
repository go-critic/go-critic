package checker_test

func suspiciousAppends() {
	var xs []int
	var ys []int

	/*! append result not assigned to the same slice */
	xs = append(ys, 1)
	/*! append result not assigned to the same slice */
	ys = append(xs, 1)

	/*! append result not assigned to the same slice */
	xs, xs[0] = append(ys, 1), ys[9]
	/*! append result not assigned to the same slice */
	ys, xs = append(ys, 1), append(ys[:])

	var withSlices struct {
		a []int
		b []int
	}
	/*! append result not assigned to the same slice */
	withSlices.a = append(withSlices.b, 1)
	/*! append result not assigned to the same slice */
	withSlices.b = append(withSlices.a, 1)

	var xsMap map[string][]int
	/*! append result not assigned to the same slice */
	xsMap["10"] = append(xsMap["100"], 1, 2)

	{
		xs2 := xs
		/*! append result not assigned to the same slice */
		xs = append(xs2, 1)
		/*! append result not assigned to the same slice */
		xs2 = append(xs, 1)
	}

	/*! append result not assigned to the same slice */
	zs := append(xs, 1)

	_ = zs
}
