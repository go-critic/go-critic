package checker_test

import (
	"io"
	"os"
)

func _(r io.Reader) {
	io.ReadAll(r)
	os.ReadFile("")
	os.WriteFile("", nil, 0)
	os.ReadDir("")
	io.NopCloser(r)
	_ = io.Discard
}
