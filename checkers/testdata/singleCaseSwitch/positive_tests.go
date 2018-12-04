package checker_test

// TODO: add tests for "no warnings" cases.
// TODO: add tests for expression switch statement.

func intValue(x interface{}) int {
	/*! should rewrite switch statement to if statement */
	switch x := x.(type) {
	case int:
		return x
	}
	return 0
}

func switchDefault(x interface{}) {
	/*! found switch with default case only */
	switch x.(type) {
	default:
	}
}

func switchWithOneCase(x int) {
	/*! should rewrite switch statement to if statement */
	switch x {
	case 1:
	}
}
