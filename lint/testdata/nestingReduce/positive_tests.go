package checker_test

func warnings1(a int) {
	/// nesting level could be reduced
	if a == 5 {
		warnings1(a)
		warnings1(a)
		warnings1(a)
		warnings1(a)
		warnings1(a)
		warnings1(a)
	}
}

func warnings2(a []int) {
	for _, v := range a {
		/// nesting level could be reduced
		if v == 5 {
			warnings2(a)
			warnings2(a)
			warnings2(a)
			warnings2(a)
			warnings2(a)
			warnings2(a)
		}
	}
}
