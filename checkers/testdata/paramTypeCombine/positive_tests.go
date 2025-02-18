package checker_test

/*! func(a int, b int, c int) could be replaced with func(a, b, c int) */
func extern(a int, b int, c int) {}

/*! func(a int, b int, c int) could be replaced with func(a, b, c int) */
func simple1(a int, b int, c int) {}

/*! func() (a int, b int) could be replaced with func() (a, b int) */
func simple2() (a int, b int) { return 0, 0 }

/*! func() (a int, b int, c int) could be replaced with func() (a, b, c int) */
func simple3() (a int, b int, c int) { return 0, 0, 0 }

/*! func(a, b int, c int) could be replaced with func(a, b, c int) */
func mixedStyle1(a, b int, c int) {}

/*! func(a, b int, c, d int) could be replaced with func(a, b, c, d int) */
func mixedStyle2(a, b int, c, d int) {}

/*! func(a, b, c int, d int) could be replaced with func(a, b, c, d int) */
func mixedStyle3(a, b, c int, d int) {}

/*! func(a int, b, c, d int) could be replaced with func(a, b, c, d int) */
func mixedStyle4(a int, b, c, d int) {}

/*! func(a, b int, c, d int, e, f int, g int) could be replaced with func(a, b, c, d, e, f, g int) */
func mixedStyle5(a, b int, c, d int, e, f int, g int) {}

/*! func() (a, b int, c int) could be replaced with func() (a, b, c int) */
func mixedStyle6() (a, b int, c int) { return 0, 0, 0 }

/*! func() (a, b int, c, d int) could be replaced with func() (a, b, c, d int) */
func mixedStyle7() (a, b int, c, d int) { return 0, 0, 0, 0 }

/*! func(a int, b, c int) (d int, e int) could be replaced with func(a, b, c int) (d, e int) */
func mixedStyle8(a int, b, c int) (d int, e int) { return a, c }

/*! func(a int, b int, c int64, d int, e, f int64, _, g int64, h int, k int) could be replaced with func(a, b int, c int64, d int, e, f, _, g int64, h, k int) */
func mixedTypeWarn(a int, b int, c int64, d int, e, f int64, _, g int64, h int, k int) {}

/*! func(_, _ int, _ int, _ int32) could be replaced with func(_, _, _ int, _ int32) */
func withBlank1(_, _ int, _ int, _ int32) {}

/*! func() (_, _ int, _ int, _ int32) could be replaced with func() (_, _, _ int, _ int32) */
func withBlank2() (_, _ int, _ int, _ int32) { return }

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
