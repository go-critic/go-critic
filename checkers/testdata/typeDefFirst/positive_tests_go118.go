//go:build go1.18
// +build go1.18

package checker_test

func (_ genericPositiveStruct[T]) Method() {}

/*! definition of type 'genericPositiveStruct' should appear before its methods */
type genericPositiveStruct[T any] struct{ _ []T }

func (_ *multiGenericPositiveStruct[T, X]) MethodWithRef() {}

/*! definition of type 'multiGenericPositiveStruct' should appear before its methods */
type multiGenericPositiveStruct[T any, X any] struct{ _ []T }
