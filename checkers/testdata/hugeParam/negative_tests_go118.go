//go:build go1.18
// +build go1.18

package checker_test

func genericFunc[T comparable](x T) {}
