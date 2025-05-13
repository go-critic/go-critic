package checker_test

import "bytes"

func _() {
	const i = 1
	_ = bytes.Repeat([]byte{1}, 5)
	_ = bytes.Repeat([]byte{'x'}, 5)
	_ = bytes.Repeat([]byte{'0'}, 5)
	_ = bytes.Repeat([]byte{i}, 5)

	const s = 'a'
	_ = bytes.Repeat([]byte{s}, 5)

	const s1 = '0'
	_ = bytes.Repeat([]byte{s1}, 5)

}
