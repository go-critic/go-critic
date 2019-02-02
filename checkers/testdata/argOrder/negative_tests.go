package checker_test

import (
	"bytes"
	"strings"
)

func nonConstArgs(s1, s2 string, b1, b2 []byte) {
	_ = strings.HasPrefix(s1, s2)
	_ = bytes.HasPrefix(b1, b2)

	x := byte('x')
	_ = bytes.HasPrefix([]byte{x}, b1)
	_ = bytes.HasPrefix([]byte(s1), b1)
}

func properArgsOrder(s string, b []byte) {
	_ = strings.HasPrefix(s, "http://")
	_ = bytes.HasPrefix(b, []byte("http://"))
	_ = bytes.HasPrefix(b, []byte{'h', 't', 't', 'p', ':', '/', '/'})
	_ = strings.Contains(s, ":")
	_ = bytes.Contains(b, []byte(":"))
}
