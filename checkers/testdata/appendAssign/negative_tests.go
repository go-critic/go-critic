package checker_test

func normalAppends() {
	var xs, ys []int

	xs = append(xs, 1)
	ys = append(ys, 1, 2)
	xs = append(xs, ys[0], xs[0])
}

func permittedAppends() {
	var xs, ys []int

	// We're trying to detect `x = append(y, ...)` patterns
	// where y is used instead of x by mistake, so lines below
	// do not trigger a warning.

	xs = append(xs, 1)
	xs0 := xs
	xs = append(xs, 1)
	xs1 := xs
	ys = append(ys, 1)
	ys0 := ys

	ts := append([]int{1}, 1)
	_ = ts

	// Also permit to assign to "_".
	_ = append(xs, xs0[0], xs1[1], ys0[0])

	{
		var m map[int][]int
		xs := m[0]
		m[0] = append(xs, 1)
	}

	// Sliced xs is still xs.
	xs = append(xs[:0], 1)
	xs = append(xs[:], 2)

	// OK to use slice literals.
	xs = append([]int{}, 1)
	xs = append([]int{1, 2}, 1)

	// Also OK to use slices returned by a function calls.
	xs = append(*new([]int), 1)
	*(new([]int)) = append(*(new([]int)), 1)

	// This prepends ys to the xs. Common idiom.
	xs = append(ys, xs...)
	xs = append(ys, xs[1:]...)

	// Scratch array idiom.
	var scratch [10]int
	xs = append(scratch[:], 1)
	xs = append(scratch[1:5])

	{
		xs := &xs
		*xs = append((*xs)[:], 1, 2)
	}

	var withSlices struct {
		a []int
		b []int
	}
	withSlices.a = append(withSlices.a, 1)
	withSlices.b = append(withSlices.b, 1)

	var xsMap map[string][]int
	xsMap["10"] = append(xsMap["10"], 1, 2)
}

func appendNotInAssignment() {
	var xs, ys []int

	// These are somewhat weird, but has nothing
	// to do with diagnostic this checker wants to perform.

	var v1 = append(xs, 1)
	var (
		v2 = append(xs, v1[0])
		v3 = append(v2[:], ys[0])
	)
	v3 = append(v3, xs[0])
	{
		v3 = append(v3, 1)
		_ = v3
	}
}
