package checker_test

func f1() {
	{
		var m map[int]int
		/// m != nil check is redundant
		if m != nil && len(m) == 10 {
		}
	}

	{
		var ch chan string
		/// nil != ch check is redundant
		if nil != ch && 0 != len(ch) {
		}
	}

	{
		var s []float64
		/// s != nil check is redundant
		if len(s) == 1 || s != nil {
		}
		/// s == nil check is redundant
		if s == nil || len(s) == 1 {
		}
	}

	{
		var ch chan int
		/// ch == nil check is redundant
		/// ch != nil check is redundant
		if (ch == nil || len(ch) == 1) != (ch != nil && len(ch) != 1) {
		}
	}
	{
		var ch chan int
		/// ch != nil check is redundant
		switch ch != nil && len(ch)+10 == 100 {
		default:
		}
	}
	{
		var ch chan int
		switch {
		/// ch != nil check is redundant
		case ch != nil || len(ch) == 0:
		/// ch == nil check is redundant
		case ch == nil && len(ch) == 0:
		default:
		}
	}
}

func f2() bool {
	var ch chan int
	/// ch != nil check is redundant
	return ch != nil && len(ch)+10 == 100
}
