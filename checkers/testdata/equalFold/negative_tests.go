package checker_test

import (
	"bytes"
	"strings"
)

func stringsEqualFold(x, y string) {
	_ = strings.EqualFold(x, y)
	_ = strings.EqualFold(x, concat(y, "123"))
	_ = strings.EqualFold(concat(y, "123"), x)
}

func bytesEqualFold(x, y []byte) {
	_ = bytes.EqualFold(x, y)
	_ = bytes.EqualFold(x, append(y, 'a'))
	_ = bytes.EqualFold(append(y, 'a'), x)
}
