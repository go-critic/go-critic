package checker_test

/// consider `in' name instead of `IN'
/// consider `out' name instead of `OUT'
func f1(IN int) (OUT int) {
	return 0
}

/// consider `in' name instead of `IN'
/// `X' should not be capitalized
/// `Y' should not be capitalized
/// `Z' should not be capitalized
func f2(IN, X int) (Y, Z int) {
	return 0, 0
}

type empty struct{}

/// consider `in' name instead of `IN'
/// consider `out' name instead of `OUT'
func (IN empty) method1(OUT *int) {}

/// `PN' should not be capitalized
func (PN empty) method2() {}
