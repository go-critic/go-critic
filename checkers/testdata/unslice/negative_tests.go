package checker_test

func sliceArray() {
	var xs [3]int
	_ = xs[:]
}

func noWarning() {
	var s string

	_ = s[1:]
	_ = s[:1]
	_ = s
}

func slicing() {
	{
		var xs []byte
		var ys []byte
		copy(xs[1:], ys[:2])
	}
	{
		var xs []int
		_ = xs[:len(xs)-1]
	}
	{
		var xs [][]int
		_ = xs[0][1:]
	}
	{
		var xs []string
		_ = xs[:0]
	}
	{
		var xs []struct{}
		_ = xs[0:]
	}
	{
		var xs map[string][][]int
		_ = xs["0"][0][:10]
	}
}
