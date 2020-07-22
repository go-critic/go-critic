package checker_test

func foo1() {
	/*! defer foo2() is placed just before return */
	defer foo2()
	return
}

func foo2() int {
	/*! defer func(){...}(...) is placed just before return */
	defer func() {}()
	return 0
}

func foo3() {
	/*! defer func(){...}(...) is placed just before return */
	defer func() {}()
}

func foo4() {
	/*! defer func(){...}(...) is placed just before return */
	defer func() {
		/*! defer foo1() is placed just before return */
		defer foo1()
		return
	}()
}

func foo5() {
	func() {
		/*! defer foo1() is placed just before return */
		defer foo1()
		return
	}()
	foo1()
	return
}

func foo6() {
	/*! defer func(){...}(...) is placed just before return */
	defer func() {
		for {
			/*! defer foo1() is placed just before return */
			defer foo1()
			return
		}
	}()
	return
}

func foo7() {
	if true {
		/*! defer foo1() is placed just before return */
		defer foo1()
		return
	}
	return
}

func returnConstExpr(s *sharedData) (string, bool, int) {
	const foo = "12"
	s.Lock()
	/*! defer s.Unlock() is placed just before return */
	defer s.Unlock()
	return foo + "3", false, len(foo) + 1
}
