//go:build go1.18
// +build go1.18

package checker_test

func genericGood1[T any](a T, b, c int)                {}
func genericGood2[T any](a, b T, c int)                {}
func genericGood3[T any](a T)                          {}
func genericGood4[T any]()                             {}
func genericGood5[T any, Y any](a, b T, c int, d, e Y) {}

func genericUnnamed[T any](T, T) {}

func genericUnnamedResults[T any]() (T, T) {
	var t T
	return t, t
}
