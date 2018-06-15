package checker_test

const length = 1

var opindex [(length + 1) & 3]*int

var _ [(length + 100 - 20*5)]*int

var _ func(int, string)

type myString string

func convertPtr(x string) *myString {
	return (*myString)(&x)
}

func multipleReturn() (int, bool) {
	return 1, true
}

type goodMap1 map[string]string

type goodMap2 map[[5][5]string]map[string]string

var _ = [4]*int{}

var _ = func() []func() { return nil }
