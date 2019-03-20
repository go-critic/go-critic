package checker_test

import (
	"bytes"
	"strings"
)

func goodStringsReplace(s, from, to string) {
	_ = strings.Replace(s, from, to, -1)
}

func goodStringsSplitN(s, sep string) {
	_ = strings.SplitN(s, sep, -1)
}

func goodBytesReplace(s, from, to []byte) {
	_ = bytes.Replace(s, from, to, -1)
}

func goodBytesSplitN(s, sep []byte) {
	_ = bytes.SplitN(s, sep, -1)
}

func goodAppend(xs []int) {
	_ = append(xs, xs[0])
	_ = append(xs[2:], 10)
}
