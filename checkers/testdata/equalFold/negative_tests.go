package checker_test

import (
	"bytes"
	"strings"
)

func changeCaseOfSameExpr(x string, b []byte) {
	_ = strings.ToLower(x) == x
	_ = x == strings.ToLower(x)
	_ = strings.ToLower(x) != x
	_ = x != strings.ToLower(x)

	_ = strings.ToUpper(x) == x
	_ = x == strings.ToUpper(x)
	_ = strings.ToUpper(x) != x
	_ = x != strings.ToUpper(x)

	_ = bytes.Equal(bytes.ToLower(b), b)
	_ = bytes.Equal(b, bytes.ToLower(b))
	_ = bytes.Equal(bytes.ToUpper(b), b)
	_ = bytes.Equal(b, bytes.ToUpper(b))
}

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

func stringsSideEffects(x, y string) {
	_ = strings.ToLower(x) == concat(y, "123")
}

func bytesSideEffects(x, y []byte) {
	_ = bytes.Equal(bytes.ToLower(x), append(y, 'a'))
}
