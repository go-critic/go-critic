//go:build go1.18
// +build go1.18

package checker_test

type genericNegativeStruct[T any] struct{ _ []T }

func (_ genericNegativeStruct[T]) Method()         {}
func (_ *genericNegativeStruct[T]) MethodWithRef() {}

type multiGenericNegativeStruct[T any, X any] struct{ _ []T }

func (_ multiGenericNegativeStruct[T, X]) Method()         {}
func (_ *multiGenericNegativeStruct[T, X]) MethodWithRef() {}
