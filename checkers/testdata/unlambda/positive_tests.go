package checker_test

type Foo struct{}

func (Foo) Method() int     { return 1 }
func (*Foo) PtrMethod() int { return 1 }

func methodExpr() {
	/*! replace `func(f Foo) int { return Foo.Method(f) }` with `Foo.Method` */
	_ = func(f Foo) int { return Foo.Method(f) }

	// TODO: should generate warning too.
	_ = func(f *Foo) int { return (*Foo).PtrMethod(f) }
}

func returnIntError(x int) (int, error) {
	return x, nil
}

func returnInt(x int) int {
	return x
}

func functionLiterals() {
	/*! replace `func(x int) int { return returnInt(x) }` with `returnInt` */
	_ = func(x int) int { return returnInt(x) }

	/*! replace `func(x int) (int, error) { return returnIntError(x) }` with `returnIntError` */
	_ = func(x int) (int, error) { return returnIntError(x) }

	/*! replace `func(x, y int) int { return add(x, y) }` with `add` */
	_ = func(x, y int) int { return add(x, y) }

	/*! replace `func(x int, y int) int { return add(x, y) }` with `add` */
	_ = func(x int, y int) int { return add(x, y) }
}

func variadicInt(xs ...int) int { return 0 }

func variadicTest() {
	_ = func(x int) int { return variadicInt(x) }
	_ = func(x int) int { return variadicInt(x, 1) }
	_ = func(x, y int) int { return variadicInt(x, y) }
	_ = func(x, y int) int { return variadicInt(x) }

	/*! replace `func(xs ...int) int { return variadicInt(xs...) }` with `variadicInt` */
	_ = func(xs ...int) int { return variadicInt(xs...) }

	_ = func(x int, ys ...int) int { return variadicInt(1, 2) }
	_ = func(x int, y int, _ ...int) int { return variadicInt(x, y) }

	/*! replace `func(options ...string) error { return wrap(append(append(append(options)))...) }` with `wrap` */
	_ = func(options ...string) error { return wrap(append(append(append(options)))...) }
}

func variadicInterfaces(x int, y interface{}, ys ...interface{}) int { return 0 }

func TestSomething() {
	// See #991
	_ = func(x int, y interface{}, _ ...interface{}) int {
		return variadicInterfaces(x, y)
	}
	_ = func(x int, y interface{}, _ ...interface{}) int {
		return variadicInterfaces(x, y, 5, "?")
	}

	/*! replace `func(x int, y interface{}, zs ...interface{}) int { return variadicInterfaces(x, y, zs...) }` with `variadicInterfaces` */
	_ = func(x int, y interface{}, zs ...interface{}) int { return variadicInterfaces(x, y, zs...) }
}

type object struct{}

func (object) returnInt(x int) int { return x }

func methodValues() {
	var o object

	/*! replace `func(x int) int { return o.returnInt(x) }` with `o.returnInt` */
	_ = func(x int) int { return o.returnInt(x) }
}
