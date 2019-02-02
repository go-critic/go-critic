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

func constOnlyArgs() {
	_ = strings.HasPrefix("", "http://")
	_ = bytes.HasPrefix([]byte{}, []byte("http://"))
	_ = bytes.HasPrefix([]byte{}, []byte{'h', 't', 't', 'p', ':', '/', '/'})
	_ = strings.Contains("", ":")
	_ = bytes.Contains([]byte{}, []byte(":"))
	_ = strings.TrimPrefix("", ":")
	_ = bytes.TrimPrefix([]byte{}, []byte(":"))
	_ = strings.TrimSuffix("", ":")
	_ = bytes.TrimSuffix([]byte{}, []byte(":"))
	_ = strings.Split("", "/")
	_ = bytes.Split([]byte{}, []byte("/"))
}

func properArgsOrder(s string, b []byte) {
	_ = strings.HasPrefix(s, "http://")
	_ = bytes.HasPrefix(b, []byte("http://"))
	_ = bytes.HasPrefix(b, []byte{'h', 't', 't', 'p', ':', '/', '/'})
	_ = strings.Contains(s, ":")
	_ = bytes.Contains(b, []byte(":"))
	_ = strings.TrimPrefix(s, ":")
	_ = bytes.TrimPrefix(b, []byte(":"))
	_ = strings.TrimSuffix(s, ":")
	_ = bytes.TrimSuffix(b, []byte(":"))
	_ = strings.Split(s, "/")
	_ = bytes.Split(b, []byte("/"))
}
