package checker_test

func ifBranches(cond1, cond2 bool) {
	if cond1 {
		println("cond=true")
	} else {
		println("cond=false")
	}

	if cond1 {
		println(1)
	} else if cond2 {
		println(2)
	} else {
		println(3)
	}

	if cond1 {
		println(1)
	} else if cond2 {
		println(1)
	}

	x := 1
	if cond1 {
		println(x)
	} else if x := 2; cond2 {
		println(x)
	}
}
