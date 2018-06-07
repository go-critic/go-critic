package checker_test

func withWarning() {
	var s string

	/// could simplify s[:] to s
	_ = s[:]
}
