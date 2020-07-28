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

func switchWithOneCaseAndDefault(x int) {
	switch x {
	default:
	case 1:
	}
}

func switchWithTwoCases(x int) {
	switch x {
	case 1:
	case 2:
	}
}

func caseWithTwoValues(x int) {
	switch x {
	case 1, 2:
	}
}

func caseWithBreak(x interface{}) {
	switch x.(type) {
	case int:
		println(x)
		break
	}

	switch x.(int) {
	case 0:
		println(x)
		break
	}

	for {
		switch x.(int) {
		case 0:
			println(x)
			break
		}
	}
}
