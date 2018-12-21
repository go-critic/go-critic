package checker_test

import "fmt"

func noWarningsNonEmptyFallthrough(i int) bool {
	switch i {
	case 0:
		fmt.Print("")
		fallthrough
	case 1:
		return true
	default:
		return false
	}
}

func noWarningsNonEmptyFallthroughToDefault(i int) bool {
	switch i {
	case 0:
		return true
	case 1:
		fmt.Print("")
		fallthrough
	default:
		return false
	}
}

func noWarningsNonEmptyFallthroughInNestedSwitch(i, j int) bool {
	switch i {
	case 0:
		return true
	case 1:
		switch j {
		case 0:
			fmt.Print("")
			fallthrough
		case 1:
			return true
		default:
			return false
		}
	case 2:
		fmt.Print("")
		fallthrough
	default:
		return false
	}
}
