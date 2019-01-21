package checker_test

func warnings() {
	var b []byte
	var s string

	/*! can simplify `[]byte(s)` to `s` */
	copy(b, []byte(s))
}
