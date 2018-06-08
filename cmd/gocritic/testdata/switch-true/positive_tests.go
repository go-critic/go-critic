package checker_test

func warnings() {
	/// replace 'switch true {}' with 'switch {}'
	switch true {
	}

	/// replace 'switch true {}' with 'switch {}'
	switch true {
	case true && false:
		println("1")
	case false && true:
		fallthrough
	default:
		println("2")
	}
}
