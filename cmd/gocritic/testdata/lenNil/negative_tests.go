package checker_test

func g1() bool {
	var m map[int]int
	if m == nil && len(m) == 0 {
		return false
	}
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
	case true:
	default:
	}
}
