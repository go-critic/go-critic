package checker_test

type foo struct{}

/// rename `x` to `_`
func f1(x int, y float64) {
	_ = y
}

/// rename `x` to `_`
/// rename `y` to `_`
func (*foo) f2(x int, y int) {
}

/// rename `y` to `_`
func (f *foo) f3(_, y int, z float64) {
	_ = z
}

/// consider to name parameters as `_`
func (f *foo) f4(int, float64) {
}

/// rename `x` to `_`
func (f *foo) f5(x int, _ float64) {
}
