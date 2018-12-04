package checker_test

import (
	"testing"
)

func TestFoo(t *testing.T) {
	// OK: Test functions are skipped (by default).
	var xs []bigObject
	for _, x := range xs {
		_ = x.x
	}
}

func bigIndex(xs []bigObject) int32 {
	// OK: no copies.
	v := int32(0)
	for i := range xs {
		v += xs[i].x
	}
	return v
}

func bigTakeAddr(xs []bigObject) int32 {
	// OK: manually taking pointers.
	v := int32(0)
	for i := range xs {
		x := &xs[i]
		v += x.x
	}
	return v
}

func bigPointers(xs []*bigObject) int32 {
	// OK: xs store pointers.
	v := int32(0)
	for _, x := range xs {
		v += x.x
	}
	return v
}
