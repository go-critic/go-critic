package checker_test

///: could simplify (func()) to func()
func badReturn() [](func()) {
	return nil
}

//TODO: could simplify [](func[](func())) to []func([]func())
///: could simplify (func([](func()))) to func([](func()))
func veryBadReturn() [](func([](func()))) {
	return nil
}

///: could simplify (func()) to func()
var _ [](func())

///: could simplify (*int) to *int
var _ [5](*int)

///: could simplify (func()) to func()
var _ [](func())

var (
	_ int
	///: could simplify (*int) to *int
	_ [5](*int)
	///: could simplify (func()) to func()
	_ [](func())
)

///: could simplify (int) to int
const _ (int) = 5

//TODO: could simplify (int) to int
type _ (int)

const ALAST = 1

var opindex [(ALAST + 1) & 3]*int

///: could simplify (ALAST + 100 - 20*5) to ALAST + 100 - 20*5
var _ [(ALAST + 100 - 20*5)]*int

var _ func(int, string)

type myString string

func convertPtr(x string) *myString {
	return (*myString)(&x)
}
