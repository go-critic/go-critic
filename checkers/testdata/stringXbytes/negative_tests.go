package checker_test

func noWarnings() {
	var b []byte
	var s string

	copy(b, s)
}

func anotherCopyFunc() {
	copy := func(int) {}

	copy(1)
}
