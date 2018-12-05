package checker_test

func f() {
	/*! consider to move `sideEffect()` before if */
	if sideEffect(); true {
	}

	/*! consider to move `sideEffect()` before switch */
	switch sideEffect(); true {
	default:
	}
}

func sideEffect() {}
