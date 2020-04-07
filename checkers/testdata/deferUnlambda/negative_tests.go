package checker_test

func negativeTests() {
	var v int

	defer f(10)

	defer f()

	// OK: can't remove function literal because of arguments
	defer func(x int) {
		println("deferred")
	}(f())

	// OK: recover() can't be moved out of lambda
	defer func() {
		recover()
	}()

	// OK: skip panic() to avoid changing the stack trace
	defer func() {
		panic("whoa!")
	}()

	// OK: function has non-const arguments
	defer func() {
		f(v)
	}()

	// OK: more than 1 statement
	defer func() {
		f()
		f()
	}()

	var o object
	objects := []object{o}

	// OK: don't report method calls.
	defer func() {
		o.f()
	}()
	defer func() {
		objects[0].f()
	}()
}

func todoTests() {
	// TODO: should be reported, because called func args
	// are already evaluated.
	defer func(v int) { f(v) }(10)
}

type object struct{}

func (object) f() {}
