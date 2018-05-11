package checker_tests

type sampleStruct struct {
	field        int
	nestedStruct struct {
		nestedField int
	}
}

type sampleInterface interface {
	method()
}

func (s *sampleStruct) method() {
}

func sampleCase() {
	var k *sampleStruct
	///: could simplify (*k).field to k.field
	(*k).field = 5
	//TODO: could simplify (*k).method() to k.method()
	///: could simplify (*k).method to k.method
	(*k).method()
	//TODO: could simplify (*k).nestedStruct.nestedField to k.nestedStruct.nestedField
	///: could simplify (*k).nestedStruct to k.nestedStruct
	(*k).nestedStruct.nestedField = 6
}

func sampleCase2(k sampleInterface) {
	k.(*sampleStruct).method()
}

func sampleCase3() {
	var k *[5]int

	//TODO: could simplify (*k)[5] to k[5]
	(*k)[2] = 3
}

func multipleIndir1() {
	type point struct{ x, y int }
	pt := point{}

	pt1 := &pt
	pt2 := &pt1
	pt3 := &pt2
	pt4 := &pt3

	//TODO: should not trigger (#68)
	///: could simplify (*pt2).x to pt2.x
	_ = (*pt2).x

	//TODO: should not trigger (#68)
	///: could simplify (**pt3).x to *pt3.x
	_ = (**pt3).x

	//TODO: should not trigger (#68)
	///: could simplify (***pt4).x to **pt4.x
	_ = (***pt4).x
}
