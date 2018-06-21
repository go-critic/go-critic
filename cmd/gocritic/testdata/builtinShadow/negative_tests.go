package checker_test

func noWarnigs() {
	var foo struct {
		len int
		cap int
	}

	foo.len = 123
	foo.cap = 321

	foo.len, foo.cap = foo.cap, foo.len
}

func noWarnParams(x int, y string) (z complex128) {
	return z
}

type someType struct{}

func (t someType) doesntShadow() {}

func genDeclNoShadow() {
	var _byte = 1
	var (
		_true  = false
		_false = 10
	)
	const _error = "error"
	const (
		_float32 = 32
		_float64 = 64
	)

	_ = _byte
	_ = _true
	_ = _false
}
