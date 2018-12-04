package checker_test

type bigStruct struct {
	x1, x2, x3, x4, x5 string
}

/*! a is heavy (1600 bytes); consider passing it by pointer */
func bigArray1(a [200]int) {}

/*! a is heavy (1024 bytes); consider passing it by pointer */
/*! b is heavy (1024 bytes); consider passing it by pointer */
func bigArray2(a, b [1024]byte) {}

/*! x is heavy (80 bytes); consider passing it by pointer */
func bigStruct1(x bigStruct) {}

/*! x is heavy (80 bytes); consider passing it by pointer */
/*! y is heavy (80 bytes); consider passing it by pointer */
func bigStruct2(x, y bigStruct) {}

/*! x is heavy (80 bytes); consider passing it by pointer */
/*! y is heavy (480 bytes); consider passing it by pointer */
func mixedBigObjects(x bigStruct, y [20][]int) {}

/*! x is heavy (80 bytes); consider passing it by pointer */
/*! y is heavy (160 bytes); consider passing it by pointer */
func (x bigStruct) bigRecv(y [2]bigStruct) {}
