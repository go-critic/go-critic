package checker_test

import "runtime"

func goodGOOS() {
	_ = runtime.GOOS == "linux"
	_ = runtime.GOOS == "darwin"
}

func nonConstant() {
	someVar := "linux"
	_ = runtime.GOOS == someVar
}

func emptyString() {
	_ = runtime.GOOS == ""
}

func unrelatedComparison() {
	x := "hello"
	_ = x == "world"
}
