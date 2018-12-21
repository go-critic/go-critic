package checker_test

func lenIndex(xs []int, ys []string) {
	/*! index expr always panics; maybe you wanted xs[len(xs)-1]? */
	_ = xs[len(xs)]
	/*! index expr always panics; maybe you wanted ys[len(ys)-1]? */
	_ = ys[len(ys)]
}
