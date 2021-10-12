package checker_test

func sliceArrayMultipleTimes() {
	var xs [3]int

	/*! could simplify xs[:][:] to xs[:] */
	_ = xs[:][:]

	/*! could simplify xs[:][:][:] to xs[:][:] */
	/*! could simplify xs[:][:] to xs[:] */
	_ = xs[:][:][:]
}

func dullStringSlicing() {
	var s string

	/*! could simplify s[:] to s */
	_ = s[:]

	/*! could simplify s[:][:] to s[:] */
	/*! could simplify s[:] to s */
	_ = s[:][:]

	/*! could simplify s[:][:][:] to s[:][:] */
	/*! could simplify s[:][:] to s[:] */
	/*! could simplify s[:] to s */
	_ = s[:][:][:]
}

func dullSlicing() {
	{
		var xs []byte
		var ys []byte
		/*! could simplify xs[:] to xs */
		/*! could simplify ys[:] to ys */
		copy(xs[:], ys[:])
	}
	{
		var xs []int
		/*! could simplify xs[:] to xs */
		_ = xs[:]
	}
	{
		var xs [][]int
		/*! could simplify xs[0][:] to xs[0] */
		_ = xs[0][:]
	}
	{
		var xs []string
		/*! could simplify xs[:] to xs */
		_ = xs[:]
	}
	{
		var xs []struct{}
		/*! could simplify xs[:] to xs */
		_ = xs[:]

		/*! could simplify xs[:][:] to xs[:] */
		/*! could simplify xs[:] to xs */
		_ = xs[:][:]

		/*! could simplify xs[:][:][:] to xs[:][:] */
		/*! could simplify xs[:][:] to xs[:] */
		/*! could simplify xs[:] to xs */
		_ = xs[:][:][:]
	}
	{
		var xs map[string][][]int
		/*! could simplify xs["0"][0][:] to xs["0"][0] */
		_ = xs["0"][0][:]
	}
}
