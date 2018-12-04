package checker_test

type point struct{ x, y int }

func lhsRhsDuplicates() {
	var p point
	var xs [2]int

	/*! suspicious identical LHS and RHS for `|` operator */
	if p.x|p.x == 0 {
	}

	/*! suspicious identical LHS and RHS for `&` operator */
	if xs[0]&xs[0] == 1 {
	}

	/*! suspicious identical LHS and RHS for `<` operator */
	if xs[1] < xs[1] {
	}

	/*! suspicious identical LHS and RHS for `>` operator */
	if xs[1] > xs[1] {
	}

	/*! suspicious identical LHS and RHS for `==` operator */
	/*! suspicious identical LHS and RHS for `<` operator */
	if p == p || 1 < 1 {
	}

	/*! suspicious identical LHS and RHS for `>=` operator */
	/*! suspicious identical LHS and RHS for `^` operator */
	if (1+p.x+3) >= (1+p.x+3) && p.y^p.y != 0 {
	}
}
