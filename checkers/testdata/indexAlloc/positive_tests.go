package checker_test

import "strings"

func positiveTests(x []byte, y string) {
	/*! consider replacing strings.Index(string(x), y) with bytes.Index(x, []byte(y)) */
	_ = strings.Index(string(x), y)

	/*! consider replacing strings.Index(string([]byte("12")), y) with bytes.Index([]byte("12"), []byte(y)) */
	_ = strings.Index(string([]byte("12")), y)

	/*! consider replacing strings.Index(string([]byte{'1', '2'}), y) with bytes.Index([]byte{'1', '2'}, []byte(y)) */
	_ = strings.Index(string([]byte{'1', '2'}), y)

	/*! consider replacing strings.Index(string(x), "a"+y) with bytes.Index(x, []byte("a" + y)) */
	_ = strings.Index(string(x), "a"+y)
}
