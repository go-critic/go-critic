package checker_test

// OK: regular reference types
func ok1(m map[int]string, c []string) (k chan float64) {
	return nil
}

/// OK: just a regular chan
func ok2(ch chan string) {}

/// OK: fine
func ok3(a *int, ch chan string) {}

/// OK: pointers to underlaying types are acceptable
func ok4(ch chan *string) chan *int {
	return nil
}

/// OK: just a func
func ok5() {}
