package checker_test

import (
	"time"
)

func deferWithCall() {
	for {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer println("test")
		break
	}

	for range []int{1, 2, 3, 4} {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer println("test")
	}
}

func deferWithClosure() {
	for {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer func() {}()

		break
	}

	for range []int{1, 2, 3, 4} {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer func() {}()
	}
}

func innerLoops() {
	for {
		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer func() {}()

			break
		}
		break
	}

	for {
		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer println(123)

			break
		}
		break
	}

	for range []int{1, 2, 3, 4} {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer func() {}()

		for range []int{1, 2, 3, 4} {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer func() {}()
		}
	}
}

func anonFunc() {
	func() {
		for range []int{1, 2, 3, 4} {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer func() {}()

			for range []int{1, 2, 3, 4} {
				/*! Possible resource leak, 'defer' is called in the 'for' loop */
				defer func() {}()
			}
		}
	}()

	for {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer println(123)
		break
	}

	func() {
		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer func() {}()

			for range []int{1, 2, 3, 4} {
				/*! Possible resource leak, 'defer' is called in the 'for' loop */
				defer func() {}()
			}

			break
		}

		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer func() {}()
			break
		}
	}()

	go func() {
		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer println(123)
			break
		}
	}()
}

func assignStmt() {
	f, ff := func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
				}
				defer println(123)
			}
		}
	}, func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
				}
				defer println(123)
			}
		}
	}

	fff := func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
				}
				defer println(123)
			}
		}
	}

	var t = func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
				}
				defer println(123)
			}
		}
	}

	var tt func() = func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
				}
				defer println(123)
			}
		}
	}

	var _ = func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
				}
				defer println(123)
			}
		}
	}

	_ = func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
				}
				defer println(123)
			}
		}
	}

	var ttt = (func() {
		{
			defer println(123)
			for x := 0; x < 5; x++ {
				/*! Possible resource leak, 'defer' is called in the 'for' loop */
				defer println(123)
			}
			defer println(123)
		}
	})

	var _ = (func() {
		{
			defer println(123)
			for range []int{1, 2, 3, 4} {
				/*! Possible resource leak, 'defer' is called in the 'for' loop */
				defer println(123)
			}
			defer println(123)
		}
	})

	f()
	ff()
	fff()
	t()
	tt()
	ttt()
}

func contextBlock() {
	{
		go func() {
			for {
				/*! Possible resource leak, 'defer' is called in the 'for' loop */
				defer println(123)
				break
			}
		}()
	}

	go (func() {
		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer println(123)
			break
		}
	})()

	{
		func() {
			for {
				/*! Possible resource leak, 'defer' is called in the 'for' loop */
				defer println(123)
				break
			}
		}()
	}

	{
		{
			func() {
				for {
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
					break
				}
			}()
		}
	}

	{
		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer println(123)
			break
		}
	}

	for {
		{
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer println(123)
			break
		}
	}

	for range []int{1, 2, 3} {
		{
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer println(123)
			break
		}
	}

	go func() {
		{
			for {
				{
					/*! Possible resource leak, 'defer' is called in the 'for' loop */
					defer println(123)
					break
				}
			}
		}
	}()
}

type x string

func funcArgs() {
	time.AfterFunc(time.Second, func() {
		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer println(123)
		}
	})

	{
		time.AfterFunc(time.Second, func() {
			{
				for {
					{
						/*! Possible resource leak, 'defer' is called in the 'for' loop */
						defer println(123)
						break
					}
				}
				defer println(123)
			}
		})
	}

	x("").closureExec(func() {
		defer println(123)
		for {
			/*! Possible resource leak, 'defer' is called in the 'for' loop */
			defer println(123)
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
						/*! Possible resource leak, 'defer' is called in the 'for' loop */
						defer println(123)
					}
					defer println(123)
				}
			}
		})
	}
}

func (x x) closureExec(f func()) { f() }
