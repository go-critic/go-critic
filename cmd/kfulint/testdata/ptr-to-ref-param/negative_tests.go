package checker_test

// OK: returned valus is not addressible, can't take address.
func ok1(m map[int]string) (k chan float64) {
	return nil
}

/// OK: 1
func ok2(ch chan string) {}

/// OK: 1231
func ok3(a *int, ch chan string) {}

/// OK: 13123
func ok4(ch chan *string) chan *int {
	return nil
}

/// OK: 131
func ok5() {}
