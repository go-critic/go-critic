package checker_test

func boolSwitch(x, y int, v []int) {
	switch {
	case true:
	/*! 'case true' is duplicated */
	case true:
	}

	switch true {
	/*! 'case true' is duplicated */
	case true, true:
	}

	switch {
	case x > 1:
	case y > 1:
	case x > 2:
	/*! 'case y > 1' is duplicated */
	case y > 1:
	}

	switch true {
	/*! 'case x > v[0]' is duplicated */
	case x > v[0], x > v[1], x > v[2], x > v[0], x > v[4]:
	/*! 'case x > v[1]' is duplicated */
	case x > v[1]:
	}
}

func intSwitch(x, y int, v []int) {
	switch x + 1 {
	case y:
	case v[0]:
	/*! 'case v[2]' is duplicated */
	case v[1], v[2], v[3], v[2]:
	}

	switch 10 {
	case x + y + v[0]:
	case x + y + v[1]:
	/*! 'case x + y + v[0]' is duplicated */
	case x + y + v[0]:
	case x + y + v[2]:
	}
}

func structSwitch() {
	type point struct{ x, y int }

	switch (point{}) {
	case point{1, 2}:
	case point{2, 1}:
	case point{3, 3}:
	/*! 'case point{1, 2}' is duplicated */
	/*! 'case point{2, 1}' is duplicated */
	case point{0, 0}, point{1, 2}, point{2, 1}:
	default:
	}
}
