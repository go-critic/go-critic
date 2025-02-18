package checker_test

func goodExtern(a, b, c int) {}

func good1(a, b int)                {}
func good2(a, b int, c int32)       {}
func good3(a, b, c int)             {}
func good4(a int, b int32, c int64) {}
func good5(a int)                   {}
func good6()                        {}

func mixedTypes(a, b int, c, d int64) {}

func unnamedParams(uint32, uint32) {}

func unnamedResults(sep []byte) (uint32, uint32) {
	return 0, 0
}

func multiline(
	a string,
	b int,
	c int,
	d []byte,
) int {
	return 0
}

func genericGood1[T any](a T, b, c int)                {}
func genericGood2[T any](a, b T, c int)                {}
func genericGood3[T any](a T)                          {}
func genericGood4[T any]()                             {}
func genericGood5[T any, Y any](a, b T, c int, d, e Y) {}

func genericUnnamed[T any](T, T) {}

func genericUnnamedResults[T any]() (T, T) {
	var t T
	return t, t
}
