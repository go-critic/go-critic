package checker_test

import (
	"sort"
)

func goodSorting() {
	{
		var xs []int
		var ys []int
		sort.Slice(xs, func(i, j int) bool {
			return (xs[i] < xs[j])
		})
		sort.Slice(ys, func(i, j int) bool {
			return (ys[i]) >= (ys[j])
		})
	}

	{
		var xs *[][]int
		sort.Slice((*xs), func(i, j int) bool {
			return ((*xs)[i][2] < (*xs)[j][2])
		})
	}

	{
		var xs [][]int
		sort.Slice(xs[0], func(i, j int) bool {
			return xs[0][i] <= xs[0][j]
		})
	}

	{
		type elem struct {
			val string
			key string
		}
		type object struct {
			elems []elem
		}
		var o object
		sort.Slice(o.elems, func(i, j int) bool {
			return o.elems[i].key < o.elems[j].key
		})
		var iface interface{} = o
		sort.Slice(iface.(object).elems, func(i, j int) bool {
			return iface.(object).elems[i].val < iface.(object).elems[j].val
		})
	}

	{
		// It's OK to sort a part of the slice.
		var xs []int
		var n int
		sort.Slice(xs[:n], func(i, j int) bool {
			return xs[i] < xs[j]
		})
	}

	{
		// OK to use type conversions.
		var xs [][]byte
		sort.Slice(xs, func(i, j int) bool {
			return string(xs[i]) > string(xs[j])
		})
	}

	{
		// OK to use len() func.
		var xs [][]byte
		sort.Slice(xs, func(i, j int) bool {
			return len(xs[i]) <= len(xs[j])
		})
	}
}

var globalSlice = []int{1, 2, 3}

func getSlice() []int { return globalSlice }

func getElem(i int) int { return globalSlice[i] }

func elemKey(i int) string { return "" }

func ignore() {
	// Skip due to the getSlice() being a function call.
	sort.Slice(getSlice(), func(i, j int) bool {
		return globalSlice[i] < globalSlice[j]
	})

	// This is also suspicious and should probably be reported,
	// but we don't analyze this case yet.
	var keys []int
	sort.Slice(globalSlice, func(i, j int) bool {
		if globalSlice[i] < globalSlice[j] {
			return true
		}
		return keys[i] < keys[j]
	})

	sort.Slice(globalSlice, func(i, j int) bool {
		return getElem(i) > getElem(j)
	})

	sort.Slice(globalSlice, func(i, j int) bool {
		return elemKey(globalSlice[i]) > elemKey(globalSlice[j])
	})
	sort.Slice(globalSlice, func(i, j int) bool {
		return elemKey(keys[i]) > elemKey(keys[j])
	})
	sort.SliceStable(globalSlice, func(i, j int) bool {
		return globalSlice[i]+globalSlice[j] > 0
	})
}
