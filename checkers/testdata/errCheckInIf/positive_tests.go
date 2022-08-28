package checker_test

type foo string

func (foo) boo() error {
	return nil
}

func (foo) boo2() (string, error) {
	return "", nil
}

func boo() error {
	return nil
}

func boo2() (string, error) {
	return "", nil
}

func warning1() {
	var err2 error
	/*! returned error 'err' must be checked */
	if err := boo(); err2 != nil {
		print(err)
	}

	var err error
	/*! returned error 'err' must be checked */
	if err = boo(); err2 != nil {
	}

	print(err)
}

func warning2() {
	var (
		d    foo
		err2 error
	)
	/*! returned error 'err' must be checked */
	if err := d.boo(); err2 != nil {
		print(err)
	}

    /*! returned error 'err' must be checked */
    if _, err := d.boo2(); err2 != nil {
		print(err)
	}

	var err error
	/*! returned error 'err' must be checked */
	if err = d.boo(); err2 != nil {
	}

    /*! returned error 'err' must be checked */
    if _, err = d.boo2(); err2 != nil {
	}

	print(err)
}

func warning3() {
	var (
		err2 error
	)
	/*! returned error 'err' must be checked */
	if _, err := boo2(); err2 != nil {
		print(err)
	}
	/*! returned error 'err' must be checked */
	if _, err := boo2(); err2 != nil {
		print(err)
	}
}
