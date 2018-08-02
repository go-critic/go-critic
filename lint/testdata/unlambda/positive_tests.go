package linter_test

func returnInt(x int) int {
	return x
}

func functionLiterals() {
	/// replace `func(x int) int { return returnInt(x) }` with `returnInt`
	_ = func(x int) int { return returnInt(x) }
}

type object struct{}

func (object) returnInt(x int) int { return x }

func methodValues() {
	var o object

	/// replace `func(x int) int { return o.returnInt(x) }` with `o.returnInt`
	_ = func(x int) int { return o.returnInt(x) }
}
