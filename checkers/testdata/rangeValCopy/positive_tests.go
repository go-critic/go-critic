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

func BenchmarkFoo(b *testing.B) {
	var xs []bigObject
	/*! each iteration copies 1032 bytes (consider pointers or indexing) */
	for _, x := range xs {
		_ = x.x
	}
}

func bigCopy(xs []bigObject) int32 {
	v := int32(0)
	/*! each iteration copies 1032 bytes (consider pointers or indexing) */
	for _, x := range xs {
		v += x.x
	}
	return v
}
