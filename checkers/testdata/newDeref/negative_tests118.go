//go:build go1.18
// +build go1.18

package checker_test

func genericNew[T any]() T {
	return *new(T)
}
