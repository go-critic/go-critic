package checker_test

func add1(x int) int { return x + 1 }

func badNilAnd(xs, ys []int) {
	var i int

	/*! suspicious `xs != nil && xs[0] == 10`; nil check may not be enough, check for len */
	_ = xs != nil && xs[0] == 10
	/*! suspicious `(xs != nil) && xs[1] == 10`; nil check may not be enough, check for len */
	_ = (xs != nil) && xs[1] == 10
	/*! suspicious `xs != nil && add1(xs[0]) == 0`; nil check may not be enough, check for len */
	_ = xs != nil && add1(xs[0]) == 0
	/*! suspicious `xs != nil && (xs[0]+xs[1]) != 0`; nil check may not be enough, check for len */
	_ = xs != nil && (xs[0]+xs[1]) != 0
	/*! suspicious `xs != nil && xs[i] > 10`; nil check may not be enough, check for len */
	_ = xs != nil && xs[i] > 10
}

func badNilOr(xs, ys []int) {
	var i int

	/*! suspicious `xs == nil || xs[0] == 10`; nil check may not be enough, check for len */
	_ = xs == nil || xs[0] == 10
	/*! suspicious `(xs == nil) || xs[1] == 10`; nil check may not be enough, check for len */
	_ = (xs == nil) || xs[1] == 10
	/*! suspicious `xs == nil || add1(xs[0]) == 0`; nil check may not be enough, check for len */
	_ = xs == nil || add1(xs[0]) == 0
	/*! suspicious `xs == nil || (xs[0]+xs[1]) != 0`; nil check may not be enough, check for len */
	_ = xs == nil || (xs[0]+xs[1]) != 0
	/*! suspicious `xs == nil || xs[i] > 10`; nil check may not be enough, check for len */
	_ = xs == nil || xs[i] > 10
}
