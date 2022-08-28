//go:build go1.18
// +build go1.18

package checker_test

/*! func[T any]() (a T, b T, c T) could be replaced with func[T any]() (a, b, c T) */
func genericSimple1[T any]() (a T, b T, c T) { return a, b, c }

/*! func[T any](a T, b T, c T) could be replaced with func[T any](a, b, c T) */
func genericSimple2[T any](a T, b T, c T) {}

/*! func[T any, Y any](a T, b T, c int, d Y, e Y) could be replaced with func[T any, Y any](a, b T, c int, d, e Y) */
func genericMixed1[T any, Y any](a T, b T, c int, d Y, e Y) {}

/*! func[T any, Y any]() (a T, b T, c int, d Y, e Y) could be replaced with func[T any, Y any]() (a, b T, c int, d, e Y) */
func genericMixed2[T any, Y any]() (a T, b T, c int, d Y, e Y) { return a, b, c, d, e }

/*! func[T any](a T, b T, c int, d int32, e int32) (f int64, g int64, h T, k T) could be replaced with func[T any](a, b T, c int, d, e int32) (f, g int64, h, k T) */
func genericMixed3[T any](a T, b T, c int, d int32, e int32) (f int64, g int64, h T, k T) {
	return f, g, h, k
}
