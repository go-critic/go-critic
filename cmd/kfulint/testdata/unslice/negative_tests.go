package checker_test

func noWarning() {
	var s string

	_ = s[1:]
	_ = s[:1]
	_ = s
}
