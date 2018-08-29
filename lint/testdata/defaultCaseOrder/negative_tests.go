package checker_test

func g1() {
	a := 10
	switch a {
	default:
		// ...
	case 5:
		// ...
	case 42:
		// ...
	}
}

func g2() {
	a := 10
	switch {
	case a == 5:
		// ...
	case a == 42:
		// ...
	default:
		// ...
	}
}
