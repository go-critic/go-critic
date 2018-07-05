package checker_tests

func (s *sampleStruct) method() {
}

func sampleCase2(k sampleInterface) {
	k.(*sampleStruct).method()
}

func multipleIndir1() {
	type point struct{ x, y int }
	pt := point{}

	pt1 := &pt
	pt2 := &pt1
	pt3 := &pt2
	pt4 := &pt3

	_ = (*pt2).x

	_ = (**pt3).x

	_ = (***pt4).x
}

type myString string

func convertPtr(x string) *myString {
	return (*myString)(&x)
}

type myInterface interface {
	f()
}

func interfacePtr() {
	var val myInterface
	ptrInterface := &val
	(*(ptrInterface)).f()
}
