package checker_test

type rec struct{}

func (r rec) Method()         {}
func (r *rec) MethodWithRef() {}

// Considering no type declaration in file it is also ok
func (r recv) Method2() {}

func JustFunction() {}

type genericNegativeStruct[T any] struct{ _ []T }

func (_ genericNegativeStruct[T]) Method()         {}
func (_ *genericNegativeStruct[T]) MethodWithRef() {}

type multiGenericNegativeStruct[T any, X any] struct{ _ []T }

func (_ multiGenericNegativeStruct[T, X]) Method()         {}
func (_ *multiGenericNegativeStruct[T, X]) MethodWithRef() {}
