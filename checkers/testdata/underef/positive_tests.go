package checker_test

type sampleStruct struct {
	field        int
	nestedStruct struct {
		nestedField int
	}
}

type sampleInterface interface {
	method()
}

var (
	globalStruct = &sampleStruct{}
	globalArray  = &[5]int{}
)

var (
	/*! could simplify (*globalStruct).field to globalStruct.field */
	_ = (*globalStruct).field

	/*! could simplify (*globalArray)[0] to globalArray[0] */
	_ = (*globalArray)[0]
)

func sampleCase() {
	var k *sampleStruct
	/*! could simplify (*k).field to k.field */
	(*k).field = 5
	//TODO: could simplify (*k).nestedStruct.nestedField to k.nestedStruct.nestedField
	/*! could simplify (*k).nestedStruct to k.nestedStruct */
	(*k).nestedStruct.nestedField = 6
}

func sampleCase3() {
	var k *[5]int

	/*! could simplify (*k)[2] to k[2] */
	(*k)[2] = 3
}

type underlyingPtr *sampleStruct

func withUnderlyingPtr(p underlyingPtr) {
	/*! could simplify (*p).field to p.field */
	_ = (*p).field

	ptr2 := &p

	/*! could simplify (**ptr2).field to (*ptr2).field */
	_ = (**ptr2).field

	ptr3 := &ptr2

	/*! could simplify (***ptr3).field to (**ptr3).field */
	_ = (***ptr3).field
}

func multiArrayDeref(xs **[2]int) {
	/*! could simplify (**xs)[0] to (*xs)[0] */
	_ = (**xs)[0]

	xsPtr := &xs

	/*! could simplify (***xsPtr)[1] to (**xsPtr)[1] */
	_ = (***xsPtr)[1]
}
