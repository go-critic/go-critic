package checker_test

func noReturnDependency() {
	var x int
	var y int

	_ = func() (int, int) {
		v := mutateArg(&x)
		return x, v
	}

	_ = func() (int, int, int, int) {
		yv := mutateArg(&y)
		xv := mutateArg(&x)
		return yv, xv, x, y
	}

	_ = func() (int, int, int) {
		v1 := mutateArg(&x)
		v2 := mutateArg(&x)
		return v1, v2, x
	}

	var o object

	_ = func() (object, int) {
		v := o.mutate()
		return o, v
	}

	_ = func() (object, int) {
		return o, o.cantMutate()
	}
}
