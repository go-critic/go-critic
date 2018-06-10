package checker_test

func noWarnings1() {
	var xs []int
	var ys []int

	// OK: can't combine if there is "..." argument.
	xs = append(xs, ys...)
	xs = append(xs, 1)

	_ = 0

	// OK: appends to different slices.
	xs = append(xs, 1)
	ys = append(ys, 2)
}

func noWarnings2() {
	xs := map[string][]int{}

	// OK: different keys.
	xs["k1"] = append(xs["k1"], 1)
	xs["k2"] = append(xs["k2"], 2)
}

func noWarnings3() {
	var xs []int

	// OK: different blocks.
	xs = append(xs, 1)
	{
		xs = append(xs, 2)
		{
			xs = append(xs, 3)
		}
		// OK: chain interrupted by block above.
		xs = append(xs, 4)
	}
}
