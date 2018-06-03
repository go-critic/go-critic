package checker_test

func withWarning() {
	var s string

	/// could simplify s[:] to s
	_ = s[:]
}

func noWarning() {
	var s string

	_ = s[1:]
	_ = s[:1]
	_ = s
}
