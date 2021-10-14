package checker_test

type unimplemented interface {
	dontImplementMe()
}

func goodTypeSwitches(x interface{}) {
	// OK: good order.
	switch x.(type) {
	case myReader:
	case *myReader:
	case reader:
	default:
	}

	// OK: good order.
	switch x.(type) {
	case *myReader:
	case myReader:
	case reader:
	default:
	}

	// OK: interface is not implemented by latter cases.
	switch x.(type) {
	case unimplemented:
	case myReader:
	case *myReader:
	case reader:
	}
}

func switchWithTwoVars(x, y int) {
	switch {
	case y <= 10:
	case x == 10:
	case x == 1:
	}

	switch {
	case y <= 10:
	case y == 100 || x == 10:
	}

	switch {
	case y <= 10:
	case y < x:
	}

	switch {
	case x > 10 && y == 2:
	case x == 1:
	}

	switch {
	case x > 10 && y == 2:
	case x == 1 || y == 3:
	}

	switch {
	case x > 10 && y == 2:
	case x == 1 && y == 3:
	}

	switch {
	case x > 10 || y == 2:
	case x == 1 && y == 3:
	}

	switch {
	case x > 10 || y == 2:
	case x == 1 || y == 3:
	}
}

func switchWithDifferentRanges(x int) {
	switch {
	case x <= 10:
	case x == 11 || x == 12:
	}

	switch {
	case x > 9:
	case x == 9:
	case x < 9:
	}
}
