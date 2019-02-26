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
	/*! case *myReader must go before the reader case */
	case *myReader:
	default:
	}

	switch x.(type) {
	/*! type is not defined myType */
	case myType:
	}
}
