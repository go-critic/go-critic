package checker_test

import "bytes"

func _() {
	/*! avoid bytes.Repeat([]byte{0}, 5); consider using make([]byte, 5) instead */
	_ = bytes.Repeat([]byte{0}, 5)

	const i = 0
	/*! avoid bytes.Repeat with a const 0; use make([]byte, 5) instead */
	_ = bytes.Repeat([]byte{i}, 5)

	const r = 0
	/*! avoid bytes.Repeat with a const 0; use make([]byte, 5) instead */
	_ = bytes.Repeat([]byte{byte(r)}, 5)

}
