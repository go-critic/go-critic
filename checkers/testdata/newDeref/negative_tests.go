package checker_test

import (
	"github.com/go-critic/go-critic/checkers/testdata/_importable/examplepkg"
)

func fixedNewExpressions() {
	_ = false
	_ = ""
	_ = 0
	_ = 0.0
	_ = int32(0)
	_ = float32(0.0)
	_ = []int(nil)
	_ = myInt(0)
	_ = point{}
	_ = []*point(nil)
	_ = point{}
	_ = [4][2]int{}
	_ = map[int]int(nil)
	_ = []map[int][]int(nil)
	_ = (*int)(nil)
	_ = examplepkg.StructType{}

	_ = interface{}(nil)
	_ = myEface(nil)
	_ = nonEmptyIface(nil)
	_ = underlyingIface(nil)
	_ = interface{}(nil)
	_ = examplepkg.InterfaceType(nil)
}
