package checker_test

/*! shadowing of predeclared identifier: int */
type int struct{}

type (
	/*! shadowing of predeclared identifier: int8 */
	int8 = int
	/*! shadowing of predeclared identifier: int16 */
	int16 = int
)

/*! shadowing of predeclared identifier: bool */
func bool() {}

var (
	/*! shadowing of predeclared identifier: float32 */
	float32 = 1
	/*! shadowing of predeclared identifier: float64 */
	float64 = 2
)

const (
	/*! shadowing of predeclared identifier: complex64 */
	complex64 = 1
	/*! shadowing of predeclared identifier: complex128 */
	complex128 = 2
)
