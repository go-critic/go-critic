package checker_test

///Loud: consider `in' name instead of `IN'
///Loud: consider `out' name instead of `OUT'
func f1(IN int) (OUT int) {
	return 0
}

///Loud: consider `in' name instead of `IN'
///Capitalized: `X' should not be capitalized
///Capitalized: `Y' should not be capitalized
///Capitalized: `Z' should not be capitalized
func f2(IN, X int) (Y, Z int) {
	return 0, 0
}

type empty struct{}

///Loud: consider `in' name instead of `IN'
///Loud: consider `out' name instead of `OUT'
func (IN empty) method1(OUT *int) {}

///Capitalized: `PN' should not be capitalized
func (PN empty) method2() {}
