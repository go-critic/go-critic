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

func goodSwitch(x int) {
	switch x {
	case 1:
	case 2, 3:
	case 4, 5, 6:
	case 7, 8, 9:
	}

	switch {
	case x == 1:
	case x == 2, x == 3 || x == 4, x == 10:
	default:
		println(x)
	}
}
