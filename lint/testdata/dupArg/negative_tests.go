package checker_test

import (
	"bytes"
	"strings"
)

func differentArgs() {
	var dstSlice, srcSlice []int
	var s, s2 string
	var b, b2 []byte

	copy(dstSlice, srcSlice)

	_ = strings.Contains(s, s2)
	_ = strings.Compare(s, s2)
	_ = strings.EqualFold(s, s2)
	_ = strings.HasPrefix(s, s2)
	_ = strings.HasSuffix(s, s2)
	_ = strings.Index(s, s2)
	_ = strings.LastIndex(s, s2)
	_ = strings.Split(s, s2)
	_ = strings.SplitAfter(s, s2)
	_ = strings.SplitAfterN(s, s2, 2)
	_ = strings.SplitN(s, s2, 2)

	_ = bytes.Contains(b, b2)
	_ = bytes.Compare(b, b2)
	_ = bytes.Equal(b, b2)
	_ = bytes.EqualFold(b, b2)
	_ = bytes.HasPrefix(b, b2)
	_ = bytes.HasSuffix(b, b2)
	_ = bytes.LastIndex(b, b2)
	_ = bytes.Split(b, b2)
	_ = bytes.SplitAfter(b, b2)
	_ = bytes.SplitAfterN(b, b2, 2)
	_ = bytes.SplitN(b, b2, 2)
}
