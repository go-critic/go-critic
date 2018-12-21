package checker_test

func warnings() {
	/*! replace 'switch true {}' with 'switch {}' */
	switch true {
	}

	/*! replace 'switch true {}' with 'switch {}' */
	switch true {
	case true && false:
		println("1")
	case false && true:
		fallthrough
	default:
		println("2")
	}

	/*! replace 'switch true := false; true {}' with 'switch true := false; {}' */
	switch true := false; true {
	case 1 < 0:
	case -1 > 0:
	}

	/*! replace 'switch _ = true; true {}' with 'switch _ = true; {}' */
	switch _ = true; true {
	}
}
