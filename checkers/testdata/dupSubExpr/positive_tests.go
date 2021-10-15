package checker_test

type point struct{ x, y int }

func lhsRhsDuplicates() {
	var p point
	var xs [2]int

	/*! suspicious identical left and right expression near | */
	if p.x|p.x == 0 {
	}

	/*! suspicious identical left and right expression near & */
	if xs[0]&xs[0] == 1 {
	}

	/*! suspicious identical left and right expression near < */
	if xs[1] < xs[1] {
	}

	/*! suspicious identical left and right expression near > */
	if xs[1] > xs[1] {
	}

	/*! suspicious identical left and right expression near == */
	/*! suspicious identical left and right expression near < */
	if p == p || 1 < 1 {
	}

	/*! suspicious identical left and right expression near >= */
	/*! suspicious identical left and right expression near ^ */
	if (1+p.x+3) >= (1+p.x+3) && p.y^p.y != 0 {
	}
}
