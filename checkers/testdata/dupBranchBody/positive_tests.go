package checker_test

func duplicatedIfBranches(cond1, cond2 bool) {
	/*! both branches in if statement have same body */
	if cond1 {
		println("cond=true")
	} else {
		println("cond=true")
	}

	if cond1 {
		println(1)
		/*! both branches in if statement have same body */
	} else if cond2 {
		println(1)
	} else {
		println(1)
	}
}
