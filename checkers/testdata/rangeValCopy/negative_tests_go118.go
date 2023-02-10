//go:build go1.18
// +build go1.18

package checker_test

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
