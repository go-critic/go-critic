package checkers

func badEmptyStringChecks(s string) {
	sptr := &s

	/*! replace `len(s) != 0` with `s != ""` */
	_ = len(s) != 0
	/*! replace `len(s) > 0` with `s != ""` */
	_ = len(s) > 0

	/*! replace `len(s) == 0` with `s == ""` */
	_ = len(s) == 0
	/*! replace `len(s) <= 0` with `s == ""` */
	_ = len(s) <= 0

	/*! replace `len(*sptr) == 0` with `*sptr == ""` */
	_ = len(*sptr) == 0
	/*! replace `len(*sptr) != 0` with `*sptr != ""` */
	_ = len(*sptr) != 0
}
