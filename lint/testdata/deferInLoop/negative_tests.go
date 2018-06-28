package checker_test

func g() {
	ff := func(i int) {}
	defer ff(0)

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			println(i)
		}()
	}

	for i := range []int{}[3:7] {
		i := i
		func() {
			defer func() {
				println(i)
			}()
		}()
	}
}
