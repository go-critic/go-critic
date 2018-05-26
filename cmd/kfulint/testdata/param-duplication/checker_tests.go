package checker_test

func good1(a, b int)                {}
func good2(a, b int, c int32)       {}
func good3(a, b, c int)             {}
func good4(a int, b int32, c int64) {}
func good5(a int)                   {}
func good6()                        {}

/// a int, b int, c int could be replaced with a, b, c int
func simple1(a int, b int, c int) {}

/// a int, b int could be replaced with a, b int
func simple2() (a int, b int) { return 0, 0 }

/// a int, b int, c int could be replaced with a, b, c int
func simple3() (a int, b int, c int) { return 0, 0, 0 }

/// a, b int, c int could be replaced with a, b, c int
func mixedStyle1(a, b int, c int) {}

/// a, b int, c, d int could be replaced with a, b, c, d int
func mixedStyle2(a, b int, c, d int) {}

/// a, b, c int, d int could be replaced with a, b, c, d int
func mixedStyle3(a, b, c int, d int) {}

/// a int, b, c, d int could be replaced with a, b, c, d int
func mixedStyle4(a int, b, c, d int) {}

/// a, b int, c, d int, e, f int, g int could be replaced with a, b, c, d, e, f, g int
func mixedStyle5(a, b int, c, d int, e, f int, g int) {}

/// a, b int, c int could be replaced with a, b, c int
func mixedStyle6() (a, b int, c int) { return 0, 0, 0 }

/// a, b int, c, d int could be replaced with a, b, c, d int
func mixedStyle7() (a, b int, c, d int) { return 0, 0, 0, 0 }

/// a int, b, c int could be replaced with a, b, c int
/// d int, e int could be replaced with d, e int
func mixedStyle8(a int, b, c int) (d int, e int) { return a, c }

func mixedTypes(a, b int, c, d int64) {}

/// a int, b int, c int64, d int, e, f int64, _, g int64, h int, k int could be replaced with a, b int, c int64, d int, e, f, _, g int64, h, k int
func mixedTypeWarn(a int, b int, c int64, d int, e, f int64, _, g int64, h int, k int) {}
