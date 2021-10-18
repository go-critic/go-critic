package checker_test

import (
	"sort"
)

func _(is []int, fs []float64, xs []string) {
	/*! suspicious sort.IntSlice usage, maybe sort.Ints was intended? */
	is = sort.IntSlice(is)
	ii := sort.IntSlice(is)

	/*! suspicious sort.Float64s usage, maybe sort.Float64s was intended? */
	fs = sort.Float64Slice(fs)
	ff := sort.Float64Slice(fs)

	/*! suspicious sort.StringSlice usage, maybe sort.Strings was intended? */
	xs = sort.StringSlice(xs)
	ss := sort.StringSlice(xs)

	_, _, _ = ii, ff, ss
}
