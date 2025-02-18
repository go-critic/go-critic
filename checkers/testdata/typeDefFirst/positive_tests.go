package checker_test

func (r recv) MethodBefore() {}

/*! definition of type 'recv' should appear before its methods */
type recv struct{}

func (_ genericPositiveStruct[T]) Method() {}

/*! definition of type 'genericPositiveStruct' should appear before its methods */
type genericPositiveStruct[T any] struct{ _ []T }

func (_ *multiGenericPositiveStruct[T, X]) MethodWithRef() {}

/*! definition of type 'multiGenericPositiveStruct' should appear before its methods */
type multiGenericPositiveStruct[T any, X any] struct{ _ []T }
