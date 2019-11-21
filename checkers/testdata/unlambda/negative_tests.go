package checker_test

import "errors"

func goodFunctionLiterals() {
	_ = returnInt
}

func goodMethodValues() {
	var o object

	_ = o.returnInt
}

func add(x, y int) int { return x + y }

func unusedArgs() {
	_ = func(string) error {
		return errors.New("123")
	}

	_ = func(s string) error {
		return errors.New("456")
	}

	_ = func(int, int) int {
		return add(1, 2)
	}

	_ = func(_ int, _ int) int {
		return add(1, 2)
	}

	_ = func(_, _ int) int {
		return add(1, 2)
	}
}

func nonMatchingCalls() {
	_ = func(x int) int {
		return add(x, 1)
	}

	_ = func(x int) int {
		return add(1, x)
	}

	_ = func() int {
		return add(1, 2)
	}

	_ = func(x, y int) int {
		return add(y, x)
	}
}

func multiStmt() {
	_ = func(x, y int) int {
		a := x
		b := y
		return add(a, b)
	}

	_ = func(x, y int) int {
		println("123")
		return add(x, y)
	}
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
		// Creates object as a part of expression.
		return object{}.returnInt(x)
	}

	_ = func(x int) (int, error) {
		// Return of multiple values.
		return returnInt(x), nil
	}

	_ = func(x int) interface{} {
		// The returnInt returns int, but enclosing func lit does
		// return interface{}.
		return returnInt(x)
	}
}

func builtins() {
	_ = func(s []int) int { return len(s) }
	_ = func(s []int) int { return cap(s) }
}

func typeConvert() {
	_ = func(x int) int32 { return int32(x) }
	_ = func(x int) int { return int(x) }
}

func varFunc() {
	var varfunc func(x int) int
	_ = func(x int) int { return varfunc(x) }
}
