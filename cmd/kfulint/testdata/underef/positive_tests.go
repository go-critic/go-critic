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

func sampleCase() {
	var k *sampleStruct
	/// could simplify (*k).field to k.field
	(*k).field = 5
	//TODO: could simplify (*k).method() to k.method()
	/// could simplify (*k).method to k.method
	(*k).method()
	//TODO: could simplify (*k).nestedStruct.nestedField to k.nestedStruct.nestedField
	/// could simplify (*k).nestedStruct to k.nestedStruct
	(*k).nestedStruct.nestedField = 6
}

func sampleCase3() {
	var k *[5]int

	/// could simplify (*k)[2] to k[2]
	(*k)[2] = 3
}
