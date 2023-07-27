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

func genericSlice[T any](original []T, f func(T)) {
	for _, v := range original {
		// OK: v is a type parameter.
		f(v)
	}
}

type Maybe[V any] struct {
	value V
}

func (m Maybe[V]) Fn() {}

type ID interface {
	IsUnique() bool
}

func testID[T ID](id T) {
	testCases := []struct {
		id T // problematic line for https://github.com/go-critic/go-critic/issues/1354
	}{
		{id},
	}

	for _, tc := range testCases {
		_ = tc.id.IsUnique()
	}
}
