package checker_test

const ALAST = 1

var opindex [(ALAST + 1) & 3]*int

var _ [(ALAST + 100 - 20*5)]*int

var _ func(int, string)

type myString string

func convertPtr(x string) *myString {
	return (*myString)(&x)
}

func multipleReturn() (int, bool) {
	return 1, true
}
