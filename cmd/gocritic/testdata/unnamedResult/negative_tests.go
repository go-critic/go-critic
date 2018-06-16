package checker_test

func (f foo) g1() int {
	return 0
}

func g2() (int, error) {
	return 0, nil
}

func g3() (x int, err error) {
	return 0, nil
}

func g4() (x int, ok bool) {
	return 0, false
}

func g5() (*foo, bool) {
	return nil, false
}

func g6() ([]int, bool) {
	return nil, false
}

func (f *foo) g7() (func(int), bool) {
	return nil, false
}

func (f foo) g8() (x int) {
	return 0
}

func (f foo) g9() (x, y int, z float64) {
	return 0, 0, 0
}

func (f foo) g10() (x, y int) {
	return 0, 0
}

func g11() (x, y, z int) {
	return 0, 0, 0
}

func g12() (x, y int, _, w float64) {
	return 0, 0, 0, 0
}
