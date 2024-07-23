package checker_test

func collectBigger(ns []int, k int) []int {
	rs := make([]int, 0)
	for _, n := range ns {
		if n > k {
			/*! append all `ns` data while range it */
			rs = append(rs, ns...)
		}
	}
	return rs
}

func collectLong(ns []string, k int) []string {
	rs := make([]string, 0)
	for _, n := range ns {
		if len(n) > k {
			/*! append all `ns` data while range it */
			rs = append(rs, ns...)
		}
	}
	return rs
}

func someFuncCall(n int) {
}

func collectBasic(ns []int, k int) []int {
	rs := make([]int, 0)
	for _, n := range ns {
		someFuncCall(n)
		/*! append all `ns` data while range it */
		rs = append(rs, ns...)
	}
	return rs
}

func collectLongCorrect(ns []string, k int) []string {
	rs := make([]string, 0)
	n := ns[0]
	rs = append(rs, n)
	for _, n := range ns {
		if len(n) > k {
			rs = append(rs, n)
		}
	}
	rs = append(rs, ns...)
	return rs
}
