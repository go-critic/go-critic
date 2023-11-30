package checker_test

// OK: regular reference types.
func ok1(m map[int]string, c []string) (k chan float64) {
	return nil
}

// OK: just a regular chan.
func ok2(ch chan string) {}

// OK: primitive type pointer + regular chan.
func ok3(a *int, ch chan string) {}

// OK: pointers to underlying types are acceptable.
func ok4(ch chan *string) chan *int {
	return nil
}

// OK: just a func
func ok5() {}

type typ map[int]int
type typ2 chan string
type typ3 []string

func (k *typ) ok(kk *typ)   {}
func (k *typ2) ok(kk *typ2) {}
func (k *typ3) ok(kk *typ3) {}

type ptrTyp *map[int]int
type ptrTyp2 *chan string
type ptrTyp3 *[]string

func ptrTypOk(kk ptrTyp)   {}
func ptrTypOk2(kk ptrTyp2) {}
func ptrTypOk3(kk ptrTyp3) {}

// OK: slices are fine.
func sliceInOut(s1, s2 *[]int) *[]float32 {
	return nil
}
