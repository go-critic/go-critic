package examplepkg

import "errors"

type StructType struct {
	A int
	B int
}

type InterfaceType interface {
	Method() int
}

var FooError = errors.New("foo")
var BarError = errors.New("bar")

func ReassignFoo(err error) {
	FooError = err
}
