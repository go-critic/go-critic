package checker_test

func getIntPtr(v *int) {}

func fixed1(retVal []int, start int) {
	for i := 0; i < start; i++ {
		retVal[i] = 0
	}
}

func mutated1(retVal []int, start int) {
	for i := 0; i < start; i++ {
		x := retVal[i]
		if x%2 == 0 {
			i--
		}
		i--
	}

	for i := 0; i < start; i++ {
		i = i - 10
	}

	for i := 0; i < start; i++ {
		for j := 0; j < 10; j++ {
			i -= j
		}
	}

	for i := 0; i < start; i++ {
		getIntPtr(&i)
	}

	for i := 0; i < start; i++ {
		ptr := &i
		*ptr -= 10
	}
}

func fixed2(x int) {
	if x < -10 || x > 10 {
	}
}

func unknownCmp2(x, y int) {
	// Don't know what value `y` have.
	if x < y && x > 10 {
	}

	if x < -10 && y > 10 {
	}
}

func fixed3(x int) {
	_ = x == 10 || x == 20

	var err error
	var err2 error
	_ = err2 != nil && err == nil

	// This one is (probably) not an error, but can be written
	// in another way, like `x == 10 && y == 10`.
	var y int
	_ = x == 10 && y == 10
}
