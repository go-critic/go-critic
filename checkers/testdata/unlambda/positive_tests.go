package checker_test

func returnIntError(x int) (int, error) {
	return x, nil
}

func returnInt(x int) int {
	return x
}

func functionLiterals() {
	/*! replace `func(x int) int { return returnInt(x) }` with `returnInt` */
	_ = func(x int) int { return returnInt(x) }

	/*! replace `func(x int) (int, error) { return returnIntError(x) }` with `returnIntError` */
	_ = func(x int) (int, error) { return returnIntError(x) }

	/*! replace `func(x, y int) int { return add(x, y) }` with `add` */
	_ = func(x, y int) int { return add(x, y) }

	/*! replace `func(x int, y int) int { return add(x, y) }` with `add` */
	_ = func(x int, y int) int { return add(x, y) }
}

type object struct{}

func (object) returnInt(x int) int { return x }

func methodValues() {
	var o object

	/*! replace `func(x int) int { return o.returnInt(x) }` with `o.returnInt` */
	_ = func(x int) int { return o.returnInt(x) }
}
