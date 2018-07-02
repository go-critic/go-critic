package checker_test

func f() {
	ff := func(i int) {}
	for i := 0; i < 10; i++ {
		/// defer will be executed only at the end of the func's scope
		defer ff(i)
	}

	for i := range []int{}[3:7] {
		i := i
		/// defer will be executed only at the end of the func's scope
		defer func() {
			println(i)
		}()
	}

	for {
		/// defer will be executed only at the end of the func's scope
		defer f()
	}
}
