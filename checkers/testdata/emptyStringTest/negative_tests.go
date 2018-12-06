package checkers

func goodEmptyStringChecks(s string) {
	sptr := &s

	_ = s == ""
	_ = s != ""

	_ = *sptr == ""
	_ = *sptr != ""
}

func sliceChecks(b []byte) {
	_ = len(b) == 0
	_ = len(b) != 0
}
