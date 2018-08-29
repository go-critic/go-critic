package checker_test

func noWarnings() {
	cond1 := true
	cond2 := true

	if cond1 {
	}

	if cond1 {
	} else {
	}

	if cond1 {
	} else if cond2 {
	}

	if cond1 {
		if cond2 {
			if true {
			} else {
			}
		} else {
		}
	} else {
	}
}

func ifelseWithInit() {
	// Don't trigger on these due to init statements.

	if true {
	} else if false {
	} else if x := 1; x > 0 {
	} else {
	}

	if x := 0; x == 0 {
	} else if y := 2; y != 0 {
	} else if true {
	}

	if x := 0; x == 0 {
		if true {
		} else if false {
		} else if x := 1; x > 0 {
		} else {
		}
	} else if y := 2; y != 0 {
		if x := 0; x == 0 {
			if x := 0; x == 0 {
			} else if y := 2; y != 0 {
			} else if true {
			}
		} else if y := 2; y != 0 {
		} else if true {
		}
	} else if true {
	}
}
