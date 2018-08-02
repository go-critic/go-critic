package linter_test

func goodFunctionLiterals() {
	_ = returnInt
}

func goodMethodValues() {
	var o object

	_ = o.returnInt
}

func complexCalls() {
	_ = func(x int) int {
		// Call result is used for something else.
		return returnInt(x) + 1
	}

	_ = func(x int) int {
		// The argument is not just forwarded.
		return returnInt(x + 1)
	}

	_ = func(x int) int {
		// Creates object as a part of expression
		return object{}.returnInt(x)
	}
}
