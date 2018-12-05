package checker_test

func makeSlice() []int {
	return []int{}
}

func goodLenIndex(xs []int, ys []string) {
	_ = xs[len(xs)-1]
	_ = ys[len(ys)-1]

	// Conservative with function call.
	// Might return different lengths for both calls.
	_ = makeSlice()[len(makeSlice())]

	var m map[int]int

	// Not an error. Doesn't panic.
	_ = m[len(m)]
}
