package checker_test

func noWarnings1() {
	if 5 == 5 {
		return
	}
}

func noWarnings2() {
	var a int
	if a == 5 {
		a++
		a++
		a++
		a++
		a++
	}
}

func noWarnings3(a interface{}) {
	if a, ok := a.(string); ok {
		noWarnings3(a)
		noWarnings3(a)
		noWarnings3(a)
		noWarnings3(a)
		noWarnings3(a)
		noWarnings3(a)
		noWarnings3(a)
		noWarnings3(a)
	}
}

func noWarnings4(a int) {
	if a == 5 {
		a++
		a++
		a++
		a++
		a++
		a++
	}
	a++
	a++
	a++
	a++
	a++
	a++
	a++
}
