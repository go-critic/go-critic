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
