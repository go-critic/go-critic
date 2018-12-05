package checker_test

func shouldWarn() {
	var cond1, cond2 bool

	if cond1 {
		/*! can replace 'else {if cond {}}' with 'else if cond {}' */
	} else {
		if cond2 {
			println(123)
		}
	}

	if cond1 {
	} else {
		if cond2 {
			/*! can replace 'else {if cond {}}' with 'else if cond {}' */
		} else {
			if cond1 {
			}
		}
	}
}
