package checker_test

func g1() bool {
	var ch chan int
	return ch != nil && len(ch)+10 == 100
}

func g2() {
	var ch chan int
	if ch == nil {
		return
	}
	if len(ch) == 0 {
		return
	}

	switch ch != nil && len(ch)+10 == 100 {
	default:
	}

	switch {
	case ch != nil || len(ch) == 0:
	case ch != nil && len(ch) == 0:
	default:
	}
}

func BUG() {
	var ch chan int
	/// ch == nil check is redundant
	/// ch != nil check is redundant
	if (ch == nil || len(ch) == 1) != (ch != nil && len(ch) != 1) {
	}
}
