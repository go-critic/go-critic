package checker_test

import "unsafe"

func typeConversions() {
	_ = func(xs []int) (a, b unsafe.Pointer) {
		// OK: unsafe.Pointer is not a function call.
		return unsafe.Pointer(&xs[0]), unsafe.Pointer(&xs[0])
	}

	_ = func(xs []int) {
		// OK: unsafe.Pointer is not a function call.
		_ = unsafe.Sizeof(xs)
		_ = unsafe.Alignof(xs)
		type A struct {
			a int
		}
		a := A{}
		_ = unsafe.Offsetof(a.a)
	}

	_ = func(x int) (a, b, c *int) {
		return (*int)(&x), (*int)(&x), (*int)(&x)
	}
}

func onlyPassive() {
	_ = func(xs []int) (int, int, int) {
		return xs[0], xs[0], xs[0]
	}

	_ = func(xs []int) (int, int, int, int) {
		return xs[0], xs[1], xs[0], xs[1]
	}
}

func unrelated() {
	_ = func(x int) (a, b, c *int) {
		return &x, &x, &x
	}

	type funcBag struct {
		a, b, c func() bool
	}
	_ = func(xs []int) funcBag {
		return funcBag{
			a: func() bool { return mayMutateSlice(&xs) },
			b: func() bool { return mayMutateSlice(&xs) },
			c: func() bool { return mayMutateSlice(&xs) },
		}
	}
}
