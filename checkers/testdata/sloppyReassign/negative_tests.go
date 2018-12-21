package checker_test

func ifStmtOK() (int, error) {
	x, err := returnsIntAndError()
	if err != nil {
		return 0, err
	}

	if err := returnsError(); err != nil {
		return 0, err
	}

	if err2 := err; err2 != nil {
		return x, err2
	}

	// Pointers are fine.
	var errPtr *error
	if *errPtr = returnsError(); *errPtr != nil {
		return 0, *errPtr
	}

	// Selector expressions are fine.
	var withError struct {
		err error
	}
	if withError.err = returnsError(); withError.err != nil {
		return 0, withError.err
	}

	return x, nil
}
