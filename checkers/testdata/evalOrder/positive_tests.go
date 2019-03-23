package checker_test

func mutateArg(x *int) int {
	*x += 10
	return *x
}

func identity(x int) int {
	return x
}

type object struct {
	val int
}

func (o *object) mutate() int {
	o.val++
	return o.val
}

func (o object) cantMutate() int {
	return o.val
}

func returnDepencency() {
	var x int
	var y int

	_ = func() (int, int) {
		/*! may want to evaluate mutateArg(&x) before the return statement */
		return x, mutateArg(&x)
	}

	_ = func() (int, int, int, int) {
		/*! may want to evaluate mutateArg(&x) before the return statement */
		/*! may want to evaluate mutateArg(&y) before the return statement */
		return mutateArg(&y), mutateArg(&x), x, y
	}

	_ = func() (int, int, int) {
		/*! may want to evaluate mutateArg(&x) before the return statement */
		return identity(x), mutateArg(&x), x
	}

	var o object

	_ = func() (object, int) {
		/*! may want to evaluate o.mutate() before the return statement */
		return o, o.mutate()
	}
}
