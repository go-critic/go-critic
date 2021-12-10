package checker_test

import "time"

func testForStmt() {
	defer println(111)

	for range []int{1, 2, 3, 4} {
		println(222)
		break
	}

	defer println(222)
}

func testRangeStmt() {
	defer println(222)

	for i := 0; i < 10; i++ {
		println(111)
	}

	defer println(222)
}

func testClosure() {
	func() {
		for {
			break
		}
		defer println(1)

		for {
			for {
				break
			}
			break
		}

		defer println(1)
	}()

	func() {
		defer println(123)

		for range []int{1, 2, 3, 4} {
			println(222)
		}

		defer println(123)

		for range []int{1, 2, 3, 4} {
			for range []int{1, 2, 3, 4} {
				println(222)
			}
		}

		defer println(123)
	}()

	for {
		func() {
			defer println(123)
		}()
		break
	}

	for {
		go func() {
			defer println()
		}()

		break
	}
}

func testBlock() {
	{
		for {
			func() {
				defer println()
			}()
			break
		}
	}
	{
		for {
			go func() {
				defer println()
			}()
			break
		}
	}
	{
		for range []int{1, 2, 3, 4} {
			go func() {
				defer println()
			}()
			break
		}
	}
	{
		for range []int{1, 2, 3, 4} {
			{
				func() {
					{
						defer println()
					}
				}()
			}
			break
		}
	}
}

func negativeAssign() {
	x := func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}

	var xx = func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}

	var xxx func() = func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}

	_ = func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}

	var _ = func() {
		{
			defer println(123)
			for {
				break
			}
			defer println(123)
		}
	}

	var _ func() = func() {
		{
			defer println(123)
			for {
				break
			}
			defer println(123)
		}
	}

	var _ = (func() {
		{
			defer println(123)
			for range []int{1, 2, 3} {
			}
			defer println(123)
		}
	})

	x()
	xx()
	xxx()
}

func negativeFuncArgs() {
	time.AfterFunc(time.Second, func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	})

	time.AfterFunc(time.Second, (func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}))

	{
		time.AfterFunc(time.Second, func() {
			{
				for {
					break
				}
				defer println(123)
			}
		})
	}

	x("").closureExec(func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	})

	h := x("")
	{
		go h.closureExec(func() {
			{
				defer println(123)
				{
					defer println(123)
					for range []int{1, 2, 3} {
					}
					defer println(123)
				}
			}
		})
	}
}
