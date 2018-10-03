package checker_test

func (s *sampleStruct) method() {
}

func sampleCase2(k sampleInterface) {
	k.(*sampleStruct).method()
}

func methodCall(v *sampleStruct) {
	(*v).method() // Call on copy
	v.method()
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

func withUnderlyingPtrOK(p underlyingPtr) {
	_ = p.field

	ptr2 := &p

	_ = (*ptr2).field

	ptr3 := &ptr2

	_ = (**ptr3).field
}

func multiArrayDerefOK(xs **[2]int) {
	_ = (*xs)[0]

	xsPtr := &xs

	_ = (**xsPtr)[1]
}
