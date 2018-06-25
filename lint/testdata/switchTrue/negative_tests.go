package checker_test

func noWarnings() {
	switch {
	}

	switch {
	case true && false:
		println("1")
	case false && true:
		fallthrough
	default:
		println("2")
	}

	switch x := 0; {
	case x < 0:
	case x > 0:
	}
}
