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

func valueSwitchesInt(x int) {
	{
		switch {
		case x < 10:
		/*! case x == 1 should go before the case x < 10 */
		case x == 1:
		}

		switch {
		case 10 > x:
		/*! case x == 1 should go before the case 10 > x */
		case x == 1:
		}

		switch {
		case 10 > x:
		/*! case 1 == x should go before the case 10 > x */
		case 1 == x:
		}

		switch {
		case x < 10:
		/*! case 1 == x should go before the case x < 10 */
		case 1 == x:
		}
	}
	{
		switch {
		case x <= 10:
		/*! case x == 10 should go before the case x <= 10 */
		case x == 10:
		}

		switch {
		case x <= 10:
		/*! case 10 == x should go before the case x <= 10 */
		case 10 == x:
		}

		switch {
		case 10 >= x:
		/*! case 10 == x should go before the case 10 >= x */
		case 10 == x:
		}

		switch {
		case 10 >= x:
		/*! case x == 10 should go before the case 10 >= x */
		case x == 10:
		}

		switch {
		case x <= 10:
		/*! case x == 5 should go before the case x <= 10 */
		case x == 5:
		}

		switch {
		case x <= 10:
		/*! case 5 == x should go before the case x <= 10 */
		case 5 == x:
		}

		switch {
		case 10 >= x:
		/*! case 5 == x should go before the case 10 >= x */
		case 5 == x:
		}

		switch {
		case 10 >= x:
		/*! case x == 5 should go before the case 10 >= x */
		case x == 5:
		}
	}

	{
		switch {
		case x > 10:
		/*! case x == 11 should go before the case x > 10 */
		case x == 11:
		}

		switch {
		case x > 10:
		/*! case 11 == x should go before the case x > 10 */
		case 11 == x:
		}

		switch {
		case 10 < x:
		/*! case 11 == x should go before the case 10 < x */
		case 11 == x:
		}
		switch {
		case 10 < x:
		/*! case x == 11 should go before the case 10 < x */
		case x == 11:
		}
	}

	{
		switch {
		case x >= 10:
		/*! case x == 10 should go before the case x >= 10 */
		case x == 10:
		}

		switch {
		case x >= 10:
		/*! case 10 == x should go before the case x >= 10 */
		case 10 == x:
		}

		switch {
		case 10 <= x:
		/*! case 10 == x should go before the case 10 <= x */
		case 10 == x:
		}

		switch {
		case 10 <= x:
		/*! case x == 10 should go before the case 10 <= x */
		case x == 10:
		}

		switch {
		case x >= 10:
		/*! case x == 12 should go before the case x >= 10 */
		case x == 12:
		}

		switch {
		case x >= 10:
		/*! case 12 == x should go before the case x >= 10 */
		case 12 == x:
		}

		switch {
		case 10 <= x:
		/*! case 12 == x should go before the case 10 <= x */
		case 12 == x:
		}

		switch {
		case 10 <= x:
		/*! case x == 12 should go before the case 10 <= x */
		case x == 12:
		}
	}
}

func valueSwitchesFloat(x float32) {
	{
		switch {
		case x < 10.1:
		/*! case x == 1.1 should go before the case x < 10.1 */
		case x == 1.1:
		}

		switch {
		case 10.1 > x:
		/*! case x == 1.1 should go before the case 10.1 > x */
		case x == 1.1:
		}

		switch {
		case 10.1 > x:
		/*! case 1.1 == x should go before the case 10.1 > x */
		case 1.1 == x:
		}

		switch {
		case x < 10.1:
		/*! case 1.1 == x should go before the case x < 10.1 */
		case 1.1 == x:
		}
	}
	{
		switch {
		case x <= 10.1:
		/*! case x == 10.1 should go before the case x <= 10.1 */
		case x == 10.1:
		}

		switch {
		case x <= 10.1:
		/*! case 10.1 == x should go before the case x <= 10.1 */
		case 10.1 == x:
		}

		switch {
		case 10.1 >= x:
		/*! case 10.1 == x should go before the case 10.1 >= x */
		case 10.1 == x:
		}

		switch {
		case 10.1 >= x:
		/*! case x == 10.1 should go before the case 10.1 >= x */
		case x == 10.1:
		}

		switch {
		case x <= 10.1:
		/*! case x == 5.2 should go before the case x <= 10.1 */
		case x == 5.2:
		}

		switch {
		case x <= 10.1:
		/*! case 5.2 == x should go before the case x <= 10.1 */
		case 5.2 == x:
		}

		switch {
		case 10.1 >= x:
		/*! case 5.2 == x should go before the case 10.1 >= x */
		case 5.2 == x:
		}

		switch {
		case 10.1 >= x:
		/*! case x == 5.2 should go before the case 10.1 >= x */
		case x == 5.2:
		}
	}

	{
		switch {
		case x > 10.1:
		/*! case x == 11.1 should go before the case x > 10.1 */
		case x == 11.1:
		}

		switch {
		case x > 10.1:
		/*! case 11.1 == x should go before the case x > 10.1 */
		case 11.1 == x:
		}

		switch {
		case 10.1 < x:
		/*! case 11.1 == x should go before the case 10.1 < x */
		case 11.1 == x:
		}
		switch {
		case 10.1 < x:
		/*! case x == 11.1 should go before the case 10.1 < x */
		case x == 11.1:
		}
	}

	{
		switch {
		case x >= 10.1:
		/*! case x == 10.1 should go before the case x >= 10.1 */
		case x == 10.1:
		}

		switch {
		case x >= 10.1:
		/*! case 10.1 == x should go before the case x >= 10.1 */
		case 10.1 == x:
		}

		switch {
		case 10.1 <= x:
		/*! case 10.1 == x should go before the case 10.1 <= x */
		case 10.1 == x:
		}

		switch {
		case 10.1 <= x:
		/*! case x == 10.1 should go before the case 10.1 <= x */
		case x == 10.1:
		}

		switch {
		case x >= 10.1:
		/*! case x == 12.2 should go before the case x >= 10.1 */
		case x == 12.2:
		}

		switch {
		case x >= 10.1:
		/*! case 12.2 == x should go before the case x >= 10.1 */
		case 12.2 == x:
		}

		switch {
		case 10.1 <= x:
		/*! case 12.2 == x should go before the case 10.1 <= x */
		case 12.2 == x:
		}

		switch {
		case 10.1 <= x:
		/*! case x == 12.2 should go before the case 10.1 <= x */
		case x == 12.2:
		}
	}
}
