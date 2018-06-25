package checker_test

func (f *foo) g1(x int, _ float64) {
	_ = x
}

func (f *foo) g2(_ int, _ float64) {
}

func g3() (_ int, _ float64) {
	return 0, 0
}

func external(x int)
