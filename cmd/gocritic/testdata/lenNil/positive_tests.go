package checker_test

func f() {
	{
		var m map[int]int
		/// consider to remove redundant nil check
		if m != nil && len(m) == 10 {
		}
	}

	{
		var ch chan string
		/// consider to remove redundant nil check
		if nil != ch && 0 != len(ch) {
		}
	}

	{
		var s []float64
		/// consider to remove redundant nil check
		if len(s) == 1 && s != nil {
		}
	}
}
