package checker_test

func g1() {
	var ch chan int
	if ch == nil {
		return
	}
	if len(ch) == 0 {
		return
	}
}

func g2() {
	var x map[int]int
	var y []int
	if x != nil && len(y) != 0 {
	}
	if x != nil && len(y) != 0 {
	}
	if len(y) != 0 || x != nil {
	}
	if 0 != len(y) || nil != x {
	}
}
