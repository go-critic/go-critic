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
