package checker_test

import "io"

type reader interface {
	Read([]byte) (int, error)
}

type myReader struct{}

func (myReader) Read(_ []byte) (int, error) { return 0, nil }

func typeSwitches(x interface{}) {
	switch x.(type) {
	case io.Reader:
	/*! case *myReader must go before the io.Reader case */
	case *myReader:
	}

	switch x.(type) {
	case reader:
	/*! case myReader must go before the reader case */
	case myReader:
	/*! case *myReader must go before the reader case */
	case *myReader:
	default:
	}

	switch x.(type) {
	case interface{}:
	/*! case reader must go before the interface{} case */
	case reader:
	/*! case myReader must go before the interface{} case */
	case myReader:
	/*! case *myReader must go before the interface{} case */
	case *myReader:
	default:
	}

	switch x.(type) {
	case reader:
	case interface{}:
	/*! case myReader must go before the reader case */
	case myReader:
		println(x)
	/*! case *myReader must go before the reader case */
	case *myReader:
		println(x)
	default:
	}

	switch x.(type) {
	/*! type is not defined myType */
	case myType:
	}
}

func valueSwitches(x int) {
	switch x {
	case 1, 2, 3:
	/*! case 4 should go before the case 1, 2, 3 */
	case 4:
	}

	switch {
	case 1 == (x - 1), 0 == x, x == 10:
		println(x)
	/*! case x != 4 should go before the case 1 == (x - 1), 0 == x, x == 10 */
	case x != 4:
	}

	switch x {
	case 1, 2, 3, 4:
	/*! case 5, 6, 7 should go before the case 1, 2, 3, 4 */
	case 5, 6, 7:
		println(x)
	/*! case 8, 9 should go before the case 1, 2, 3, 4 */
	case 8, 9:
		println(x)
	}
}
