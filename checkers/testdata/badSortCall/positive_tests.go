package checker_test

import (
	"sort"
)

func _(is []int, fs []float64, xs []string) {
	/*! bad sort.IntSlice usage */
	is = sort.IntSlice(is)
	ii := sort.IntSlice(is)

	/*! bad sort.Float64Slice usage */
	fs = sort.Float64Slice(fs)
	ff := sort.Float64Slice(fs)

	/*! bad sort.StringSlice usage */
	xs = sort.StringSlice(xs)
	ss := sort.StringSlice(xs)

	_, _, _ = ii, ff, ss
}
