package checker_test

func newError() error { return nil }

func bad1(retVal []int, start int) {
	/*! `i > start` in loop; probably meant `i < start`? */
	for i := 0; i > start; i++ {
		retVal[i] = 0
	}
}

func bad2(x int) {
	/*! `x < -10 && x > 10` condition is always false */
	if x < -10 && x > 10 {
	}

	/*! `(x < -10) && x > 10` condition is always false */
	_ = (x < -10) && x > 10
	/*! `x < -10 && (x > 10)` condition is always false */
	_ = x < -10 && (x > 10)
	/*! `(x < -10) && (x > 10)` condition is always false */
	_ = (x < -10) && (x > 10)
}

func bad3(x int) {
	/*! `x == 10 && x == 20` condition is suspicious */
	_ = x == 10 && x == 20
	/*! `(x == 10) && x == 20` condition is suspicious */
	_ = (x == 10) && x == 20
	/*! `x == 10 && (x == 20)` condition is suspicious */
	_ = x == 10 && (x == 20)
	/*! `(x == 10) && (x == 20)` condition is suspicious */
	_ = (x == 10) && (x == 20)

	var err error
	/*! `err == nil && err == newError()` condition is suspicious */
	_ = err == nil && err == newError()

	// This one is (probably) not an error, but can be written
	// in another way, like `x == 10 && y == 10`.
	var y int
	/*! `x == 10 && x == y` condition is suspicious */
	_ = x == 10 && x == y
}
