package checker_test

import (
	"bytes"
	"strings"
)

func fixedCode(x []byte, y string) {
	_ = bytes.Index(x, []byte(y))
	_ = bytes.Index([]byte("12"), []byte(y))
	_ = bytes.Index([]byte{'1', '2'}, []byte(y))
	_ = bytes.Index(x, []byte("a"+y))
}

func getBytes() []byte  { return nil }
func getString() string { return "" }

func unsafeArgs(x []byte, y string) {
	_ = bytes.Index(getBytes(), []byte(y))
	_ = bytes.Index(getBytes(), []byte("a"+y))

	_ = bytes.Index(x, []byte(getString()))
	_ = bytes.Index([]byte("12"), []byte(getString()))
	_ = bytes.Index([]byte{'1', '2'}, []byte(getString()))
	_ = bytes.Index(x, []byte("a"+getString()))

	_ = strings.Index(string(getBytes()), y)
	_ = strings.Index(string(getBytes()), "a"+y)

	_ = strings.Index(string(x), getString())
	_ = strings.Index(string([]byte(getString())), getString())
	_ = strings.Index(string(x), "a"+getString())
}
