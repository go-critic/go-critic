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

type kek map[int]int
type kek2 chan string
type kek3 []string

func (k *kek) ok(kk *kek)   {}
func (k *kek2) ok(kk *kek2) {}
func (k *kek3) ok(kk *kek3) {}

type wow *map[int]int
type wow2 *chan string
type wow3 *[]string

func wowok(kk wow)   {}
func wowok2(kk wow2) {}
func wowok3(kk wow3) {}
