package checker_test

func good() {
	r := []uint64{10}
	_ = r[0]
	r2 := []rune{10, 12, 34}
	_ = r2[0]
}
