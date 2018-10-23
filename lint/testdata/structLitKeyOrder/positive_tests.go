package checker_test

type point struct {
	x int
	y int
	z int
}

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
