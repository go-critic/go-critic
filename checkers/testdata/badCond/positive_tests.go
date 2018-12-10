package checker_test

func bad1(retVal []int, start int) {
	/*! `i > start` in loop; probably meant `i < start`? */
	for i := 0; i > start; i++ {
		retVal[i] = 0
	}
}

func bad2(x int) {
	/*! `x < -10 && x > 10` is always false; probably meant `x < -10 || x > 10`? */
	if x < -10 && x > 10 {
	}
}
