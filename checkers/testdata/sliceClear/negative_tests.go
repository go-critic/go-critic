package checker_test

func _() {
	{
		var xs []int
		for i := range xs {
			xs[i] = 0
		}
	}

	{
		var xs []int
		var ys []int
		for i := 0; i < len(xs); i++ {
			ys[i] = 0
		}
	}
}
