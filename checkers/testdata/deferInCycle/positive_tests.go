package checker_test

import "fmt"

func deferWithCall() {
	for {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer fmt.Println("test")
		break
	}

	for range []int{1, 2, 3, 4} {
		/*! Possible resource leak, 'defer' is called in the 'for' loop */
		defer fmt.Println("test")
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
			defer fmt.Println(123)

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
}
