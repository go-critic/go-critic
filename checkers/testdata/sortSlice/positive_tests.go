package checker_test

import (
	"sort"
)

func badSorting() {
	{
		var xs []int
		var ys []int
		sort.Slice(xs, func(i, j int) bool {
			/*! cmp func must use xs slice in comparison */
			return (ys[i] < ys[j])
		})
		sort.Slice(ys, func(i, j int) bool {
			/*! cmp func must use ys slice in comparison */
			return (xs[i]) >= (xs[j])
		})

		// Same as above, but for SliceStable func.
		sort.SliceStable(xs, func(i, j int) bool {
			/*! cmp func must use xs slice in comparison */
			return (ys[i] < ys[j])
		})
		sort.SliceStable(ys, func(i, j int) bool {
			/*! cmp func must use ys slice in comparison */
			return (xs[i]) >= (xs[j])
		})
	}

	{
		var xs *[][]int
		var ys []int
		sort.Slice((*xs), func(i, j int) bool {
			/*! cmp func must use *xs slice in comparison */
			return (ys[i] < ys[j])
		})
	}

	{
		var xs [][]int
		var ys []int
		sort.Slice(xs[0], func(i, j int) bool {
			/*! cmp func must use xs[0] slice in comparison */
			return ys[i] <= ys[j]
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
		var ys []string
		sort.Slice(o.elems, func(i, j int) bool {
			/*! cmp func must use o.elems slice in comparison */
			return ys[i] < ys[j]
		})
		var iface interface{} = o
		sort.Slice(iface.(object).elems, func(i, j int) bool {
			/*! cmp func must use iface.(object).elems slice in comparison */
			return ys[i] < ys[j]
		})
	}
}

func swappedIndex() {
	{
		{
			var xs []int
			sort.Slice(xs, func(i, j int) bool {
				/*! unusual order of {i,j} params in comparison */
				return (xs[j]) >= (xs[i])
			})
		}
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
		sort.Slice(o.elems, func(i int, j int) bool {
			/*! unusual order of {i,j} params in comparison */
			return o.elems[j].key < o.elems[i].key
		})
		var iface interface{} = o
		sort.Slice(iface.(object).elems, func(a, b int) bool {
			/*! unusual order of {a,b} params in comparison */
			return iface.(object).elems[b].val < iface.(object).elems[a].val
		})
	}
}
