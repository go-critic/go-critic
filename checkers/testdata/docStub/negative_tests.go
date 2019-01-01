package checker_test

// Bar returns string.
func Bar() {}

func foo() {}

// Issue 24791.
func MyFunc() {
}
