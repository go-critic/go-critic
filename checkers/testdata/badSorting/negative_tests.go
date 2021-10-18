package checker_test

import "sort"

func _(is []int, fs []float64, xs []string) {
	sort.Ints(is)
	sort.IntSlice(is).Sort()

	sort.Float64s(fs)
	sort.Float64Slice(fs).Sort()

	sort.Strings(xs)
	sort.StringSlice(xs).Sort()
}
