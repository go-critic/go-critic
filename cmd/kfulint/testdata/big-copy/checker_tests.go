package checker_test

import (
	"testing"
)

type bigObject struct {
	// Fields are carefuly selected to get equal struct size
	// for both AMD64 and 386.

	body [1024]byte
	x    int32
	y    int32
}

func TestFoo(t *testing.T) {
	// OK: Test functions are skipped.
	var xs []bigObject
	for _, x := range xs {
		_ = x.x
	}
}

func BenchmarkFoo(b *testing.B) {
	var xs []bigObject
	/// each iteration copies 1032 bits (consider pointers or indexing)
	for _, x := range xs {
		_ = x.x
	}
}

func bigCopy(xs []bigObject) int32 {
	v := int32(0)
	/// each iteration copies 1032 bits (consider pointers or indexing)
	for _, x := range xs {
		v += x.x
	}
	return v
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
