package checker_test

func warnings1() {
	var xs []int

	/*! can combine chain of 2 appends into one */
	xs = append(xs, 1)
	xs = append(xs, 2)

	_ = 0

	/*! can combine chain of 2 appends into one */
	xs = append(xs, 1, 2)
	xs = append(xs, 3, 4)

	_ = 0

	/*! can combine chain of 3 appends into one */
	xs = append(xs, 1)
	xs = append(xs, 2)
	xs = append(xs, 3)

	switch len(xs) == 0 {
	case true:
		/*! can combine chain of 2 appends into one */
		xs = append(xs, 1)
		xs = append(xs, 2)
	case false:
		/*! can combine chain of 4 appends into one */
		xs = append(xs, 1)
		xs = append(xs, 2)
		xs = append(xs, 3)
		xs = append(xs, 4, 5, 6)
	default:
		// Intermixing chains and breaks.

		var ys []int
		xs = append(xs, ys...)
		/*! can combine chain of 2 appends into one */
		xs = append(xs, 1, 2)
		xs = append(xs, 3)
		xs = append(xs, ys...)
		xs = append(xs, 4)
		xs = append(xs, ys...)
		/*! can combine chain of 3 appends into one */
		xs = append(xs, 5, 6)
		xs = append(xs, 7, 8)
		xs = append(xs, 9)
	}

	ch := make(chan bool)
	select {
	case <-ch:
		/*! can combine chain of 2 appends into one */
		xs = append(xs, 1)
		xs = append(xs, 2)
		if ch != nil {
			/*! can combine chain of 2 appends into one */
			xs = append(xs, 5)
			xs = append(xs, 6)
		} else {
			/*! can combine chain of 2 appends into one */
			xs = append(xs, 7)
			xs = append(xs, 8)
		}
	default:
		/*! can combine chain of 2 appends into one */
		xs = append(xs, 3)
		xs = append(xs, 4)
	}

	/*! can combine chain of 3 appends into one */
	xs = append(xs, 1)
	// Comments can't break the chain.
	xs = append(xs, 2)
	// Even if there are multiple.
	xs = append(xs, 3)
}

func warnings2() {
	xs := map[string][]int{}

	/*! can combine chain of 2 appends into one */
	xs["k"] = append(xs["k"], 1)
	xs["k"] = append(xs["k"], 2)

	xs["k1"] = append(xs["k1"], 1)
	/*! can combine chain of 2 appends into one */
	xs["k2"] = append(xs["k2"], 2)
	xs["k2"] = append(xs["k2"], 3)
	xs["k3"] = append(xs["k3"], 4)
	xs["k2"] = append(xs["k2"], 5)
}
