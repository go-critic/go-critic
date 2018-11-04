package checker_test

type foo struct {
	a  int
	pt point
	b  int
	bar
	c int
}

type bar struct{}

type point struct {
	x int
	y int
	z int
}

/// literal key order does not match declaration key order
var _ = foo{b: 1, pt: point{}}

/// literal key order does not match declaration key order
var _ = foo{c: 1, bar: bar{}}

/// literal key order does not match declaration key order
var _ = foo{a: 1, pt: point{y: 1, x: 2}}

/// literal key order does not match declaration key order
var _ = point{y: 1, x: 0}

/// literal key order does not match declaration key order
var _ = point{z: 2, x: 0}

/// literal key order does not match declaration key order
var _ = point{z: 1, y: 0}

func chaoticFieldsOrder() {
	/// literal key order does not match declaration key order
	_ = point{y: 1, x: 0}

	/// literal key order does not match declaration key order
	_ = point{z: 2, x: 0}

	/// literal key order does not match declaration key order
	_ = point{z: 1, y: 0}
}
