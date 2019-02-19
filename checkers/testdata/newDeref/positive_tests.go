package checker_test

import (
	"github.com/go-critic/go-critic/checkers/testdata/_importable/examplepkg"
)

type point struct {
	x float64
	y float64
}

type myInt int

func badNewExpressions() {
	/*! replace `*new(bool)` with `false` */
	_ = *new(bool)

	/*! replace `*new(string)` with `""` */
	_ = *new(string)

	/*! replace `*new(int)` with `0` */
	_ = *new(int)

	/*! replace `*new(float64)` with `0.0` */
	_ = *new(float64)

	/*! replace `*new(int32)` with `int32(0)` */
	_ = *new(int32)

	/*! replace `*new(float32)` with `float32(0.0)` */
	_ = *new(float32)

	/*! replace `*new([]int)` with `[]int(nil)` */
	_ = *new([]int)

	/*! replace `*new(myInt)` with `myInt(0)` */
	_ = *new(myInt)

	/*! replace `*new(point)` with `point{}` */
	_ = *new(point)

	/*! replace `*new([]*point)` with `[]*point(nil)` */
	_ = *new([]*point)

	/*! replace `*new((point))` with `point{}` */
	_ = *new((point))

	/*! replace `*new([4][2]int)` with `[4][2]int{}` */
	_ = *new([4][2]int)

	/*! replace `*new(map[int]int)` with `map[int]int(nil)` */
	_ = *new(map[int]int)

	/*! replace `*new([]map[int][]int)` with `[]map[int][]int(nil)` */
	_ = *new([]map[int][]int)

	/*! replace `*new(*int)` with `(*int)(nil)` */
	_ = *new(*int)

	/*! replace `*new(examplepkg.StructType)` with `examplepkg.StructType{}` */
	_ = *new(examplepkg.StructType)
}

type myEface interface{}

type nonEmptyIface interface {
	Foo()
	Bar()
}

type underlyingIface nonEmptyIface

func interfaceDeref() {
	/*! replace `*new(interface{})` with `interface{}(nil)` */
	_ = *new(interface{})

	/*! replace `*new(myEface)` with `myEface(nil)` */
	_ = *new(myEface)

	/*! replace `*new(nonEmptyIface)` with `nonEmptyIface(nil)` */
	_ = *new(nonEmptyIface)

	/*! replace `*new(underlyingIface)` with `underlyingIface(nil)` */
	_ = *new(underlyingIface)

	/*! replace `*new(interface{})` with `interface{}(nil)` */
	_ = *new(interface{})

	/*! replace `*new(examplepkg.InterfaceType)` with `examplepkg.InterfaceType(nil)` */
	_ = *new(examplepkg.InterfaceType)
}
