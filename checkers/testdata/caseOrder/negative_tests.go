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
