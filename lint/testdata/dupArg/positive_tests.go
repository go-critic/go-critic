package checker_test

import (
	"bytes"
	"strings"
)

func duplicatedArgs() {
	var dstSlice []int
	var s string
	var b []byte

	/// suspicious duplicated args in `copy(dstSlice, dstSlice)`
	copy(dstSlice, dstSlice)

	/// suspicious duplicated args in `strings.Contains(s, s)`
	_ = strings.Contains(s, s)
	/// suspicious duplicated args in `strings.Compare(s, s)`
	_ = strings.Compare(s, s)
	/// suspicious duplicated args in `strings.EqualFold(s, s)`
	_ = strings.EqualFold(s, s)
	/// suspicious duplicated args in `strings.HasPrefix(s, s)`
	_ = strings.HasPrefix(s, s)
	/// suspicious duplicated args in `strings.HasSuffix(s, s)`
	_ = strings.HasSuffix(s, s)
	/// suspicious duplicated args in `strings.Index(s, s)`
	_ = strings.Index(s, s)
	/// suspicious duplicated args in `strings.LastIndex(s, s)`
	_ = strings.LastIndex(s, s)
	/// suspicious duplicated args in `strings.Split(s, s)`
	_ = strings.Split(s, s)
	/// suspicious duplicated args in `strings.SplitAfter(s, s)`
	_ = strings.SplitAfter(s, s)
	/// suspicious duplicated args in `strings.SplitAfterN(s, s, 2)`
	_ = strings.SplitAfterN(s, s, 2)
	/// suspicious duplicated args in `strings.SplitN(s, s, 2)`
	_ = strings.SplitN(s, s, 2)

	/// suspicious duplicated args in `bytes.Contains(b, b)`
	_ = bytes.Contains(b, b)
	/// suspicious duplicated args in `bytes.Compare(b, b)`
	_ = bytes.Compare(b, b)
	/// suspicious duplicated args in `bytes.Equal(b, b)`
	_ = bytes.Equal(b, b)
	/// suspicious duplicated args in `bytes.EqualFold(b, b)`
	_ = bytes.EqualFold(b, b)
	/// suspicious duplicated args in `bytes.HasPrefix(b, b)`
	_ = bytes.HasPrefix(b, b)
	/// suspicious duplicated args in `bytes.HasSuffix(b, b)`
	_ = bytes.HasSuffix(b, b)
	/// suspicious duplicated args in `bytes.LastIndex(b, b)`
	_ = bytes.LastIndex(b, b)
	/// suspicious duplicated args in `bytes.Split(b, b)`
	_ = bytes.Split(b, b)
	/// suspicious duplicated args in `bytes.SplitAfter(b, b)`
	_ = bytes.SplitAfter(b, b)
	/// suspicious duplicated args in `bytes.SplitAfterN(b, b, 2)`
	_ = bytes.SplitAfterN(b, b, 2)
	/// suspicious duplicated args in `bytes.SplitN(b, b, 2)`
	_ = bytes.SplitN(b, b, 2)
}
