package checker_test

/*! func(a int, b int, c int) could be replaced with func(a, b, c int) */
func extern(a int, b int, c int)

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
