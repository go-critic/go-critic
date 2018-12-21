package checker_test

func returnsError() error { return nil }

func returnsIntAndError() (int, error) { return 0, nil }

func ifStmtInitReassign() (int, error) {
	x, err := returnsIntAndError()
	if err != nil {
		return 0, err
	}

	/*! re-assignment to `err` can be replaced with `err := returnsError()` */
	if err = returnsError(); err != nil {
		return 0, err
	}

	var err2 error
	/*! re-assignment to `err2` can be replaced with `err2 := err` */
	if err2 = err; err2 != nil {
		return x, err2
	}
	return x, nil
}
