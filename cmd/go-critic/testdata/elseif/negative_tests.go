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
