package checker_test

func f() {
	a := 10
	switch a {
	case 5:
		// ...
	/*! consider to make `default` case as first or as last case */
	default:
		// ...
	case 42:
		// ...
	}
}
