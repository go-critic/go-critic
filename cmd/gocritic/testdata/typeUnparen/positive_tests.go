package checker_test

/// could simplify [](func()) to []func()
func badReturn() [](func()) {
	return nil
}

/// could simplify [](func([](func()))) to []func([]func())
func veryBadReturn() [](func([](func()))) {
	return nil
}

/// could simplify [](func()) to []func()
var _ [](func())

/// could simplify [5](*int) to [5]*int
var _ [5](*int)

/// could simplify [](func()) to []func()
var _ [](func())

var (
	_ int
	/// could simplify [5](*int) to [5]*int
	_ [5](*int)
	/// could simplify [](func()) to []func()
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
	/// could simplify [](int) to []int
	foo([](int))

	/// could simplify [](func() string) to []func() string
	bar() [](func() string)
}

func myFunc1() {
	func() {
		type localType1 struct {
			/// could simplify ([]complex128) to []complex128
			x ([]complex128)
		}

		_ = struct {
			/// could simplify (struct{...}) to struct{...}
			_ struct{ x (struct{}) }
			_ struct {
				/// could simplify (struct{...}) to struct{...}
				y (struct {
					/// could simplify (struct{...}) to struct{...}
					_ (struct{})
				})
			}

			/// could simplify (struct{...}) to struct{...}
			_ (struct {
				x int
				y int
			})
		}{}
	}()

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

/// could simplify map[(string)](string) to map[string]string
type mapType1 map[(string)](string)

/// could simplify map[[5][5](string)]map[(string)](string) to map[[5][5]string]map[string]string
type mapType2 map[[5][5](string)]map[(string)](string)

/// could simplify [4](*int) to [4]*int
var _ = [4](*int){}

/// could simplify func() [](func()) to func() []func()
var _ = func() [](func()) { return nil }

var _ = []interface{}{
	/// could simplify (complex64) to complex64
	struct{ x (complex64) }{},

	func() {
		/// could simplify (mapType1) to mapType1
		type T (mapType1)

		var (
			/// could simplify [](interface{}) to []interface{}
			_ = [](interface{}){}
		)
	},
}
