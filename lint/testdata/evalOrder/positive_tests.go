package checker_test

func mayMutateInt1(x *int) int { return 0 }

func mayMutateInt2(x *int, y *int) int { return 0 }

func mayMutateSlice(xs *[]int) bool { return false }

func multiReturn() {
	// See https://github.com/golang/go/issues/25609.

	_ = func(xs []int) (bool, int) {
		/// potential dependency on evaluation order (low)
		return mayMutateSlice(&xs), xs[0]
	}

	_ = func(xs []int) (int, bool) {
		/// potential dependency on evaluation order (low)
		return xs[0], mayMutateSlice(&xs)
	}

	_ = func(xs []int) (int, bool, int) {
		/// potential dependency on evaluation order (average)
		return xs[0], mayMutateSlice(&xs), xs[0]
	}

	_ = func(xs []int) (int, bool, int, int) {
		/// potential dependency on evaluation order (high)
		return xs[0], mayMutateSlice(&xs), xs[0], xs[1]
	}

	_ = func(xs []int) (bool, bool) {
		/// potential dependency on evaluation order (low)
		return mayMutateSlice(&xs), mayMutateSlice(&xs)
	}

	_ = func(x int) (int, int) {
		/// potential dependency on evaluation order (low)
		return mayMutateInt1(&x), mayMutateInt1(&x)
	}

	_ = func(x *int) (int, int) {
		/// potential dependency on evaluation order (low)
		return mayMutateInt1(x), mayMutateInt1(x)
	}

	_ = func(x, y, z int) (int, int) {
		/// potential dependency on evaluation order (low)
		return mayMutateInt2(&x, &y), mayMutateInt2(&z, &x)
	}

	_ = func(x, y, z int) (int, int, int) {
		/// potential dependency on evaluation order (average)
		return mayMutateInt1(&x), mayMutateInt2(&x, &z), mayMutateInt1(&z)
	}

	_ = func(x *int, y, z int) (int, int, int) {
		/// potential dependency on evaluation order (high)
		return mayMutateInt2(&y, x), mayMutateInt2(x, &z), mayMutateInt2(&z, &y)
	}

	type funcBag struct {
		a, b func() (int, bool)
	}
	_ = func(xs []int) (int, int, funcBag) {
		return xs[0], xs[0], funcBag{
			a: func() (int, bool) {
				/// potential dependency on evaluation order (low)
				return xs[0], mayMutateSlice(&xs)
			},
			b: func() (int, bool) {
				/// potential dependency on evaluation order (low)
				return xs[0], mayMutateSlice(&xs)
			},
		}
	}
}
