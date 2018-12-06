package checker_test

func g() {
	if x := foo(); x == 0 {
	}

	switch x := foo(); x {
	default:
	}

	for _ = foo(); ; {
	}
}

func foo() int {
	return 0
}
