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

	///: could simplify (*k)[2] to k[2]
	(*k)[2] = 3
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
