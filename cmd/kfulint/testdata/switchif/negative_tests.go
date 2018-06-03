package checker_test

func intValue1(x interface{}) int {
	switch x := x.(type) {
	case int:
		return x
	case uint:
		return 1
	}
	return 0
}

func switchWithOneCaseDefault(x int) {
	switch x {
	default:
	case 1:
	}
}

func switchWithTwoCase(x int) {
	switch x {
	case 1:
	case 2:
	}
}
