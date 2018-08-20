package checker_tests

import "fmt"

func noWarningsNonEmptyFallthrough(i int) int {
	switch i {
	case 0:
		fmt.Print("")
		fallthrough
	case 1:
		return 1
	default:
		return 2
	}
}

func noWarningsNonEmptyFallthroughToDefault(i int) int {
	switch i {
	case 0:
		return 0
	case 1:
		fmt.Print("")
		fallthrough
	default:
		return 1
	}
}
