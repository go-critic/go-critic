package checker_test

/// could simplify (func()) to func()
func badReturn() [](func()) {
	return nil
}

//TODO: could simplify [](func[](func())) to []func([]func())
/// could simplify (func([](func()))) to func([](func()))
func veryBadReturn() [](func([](func()))) {
	return nil
}

/// could simplify (func()) to func()
var _ [](func())

/// could simplify (*int) to *int
var _ [5](*int)

/// could simplify (func()) to func()
var _ [](func())

var (
	_ int
	/// could simplify (*int) to *int
	_ [5](*int)
	/// could simplify (func()) to func()
	_ [](func())
)

/// could simplify (int) to int
const _ (int) = 5

/// could simplify (int) to int
type _ (int)

type myStruct1 struct {
	/// could simplify (int) to int
	x (int)

	/// could simplify (int64) to int64
	y (int64)
}

type myInterface1 interface {
	/// could simplify (int) to int
	foo([](int))

	/// could simplify (func() string) to func() string
	bar() [](func() string)
}

func myFunc1() {
	type localType1 struct {
		/// could simplify (int) to int
		x (int)
	}

	/// could simplify (int) to int
	type localType2 (int)

	const (
		/// could simplify (int) to int
		localConst1 (int) = 1
		/// could simplify (string) to string
		localConst2 (string) = "1"
	)

	var (
		/// could simplify (int) to int
		localVar1 (int) = 1
		/// could simplify (string) to string
		localVar2 (string) = "1"
	)

	_ = localVar1
	_ = localVar2
}
