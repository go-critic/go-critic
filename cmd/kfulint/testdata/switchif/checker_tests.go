package checker_test

func intValue(x interface{}) int {
	///: should rewrite switch statement to if statement
	switch x := x.(type) {
	case int:
		return x
	}
	return 0
}

func switchDefault(x interface{}) {
	///: found switch with default case only
	switch x.(type) {
	default:
	}
}
