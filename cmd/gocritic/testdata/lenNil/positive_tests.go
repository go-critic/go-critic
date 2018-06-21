package checker_test

func f() {
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
}
