package checker_test

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
