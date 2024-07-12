package checker_test

func collectWithNil(ns []string, k int) bool {
	rs := make([]string, 0)
	rs2 := make([]string, 0)
	rs3 := make([]string, 0)
	for _, n := range ns {
		if len(n) > k {
			rs = append([]string{n}, ns...)
			rs2 = append([]string{}, ns...)
			rs3 = append([]string(nil), ns...)
		}
	}
	if len(rs) > 0 || len(rs2) > 0 || len(rs3) > 0 {
		return true
	}
	return false
}
