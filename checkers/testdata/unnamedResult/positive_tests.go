package checker_test

type foo struct{}

/*! consider giving a name to these results */
func (f *foo) f1() (float64, float64) {
	return 0, 0
}

/*! consider giving a name to these results */
func f2() (int, float64) {
	return 0, 0
}

/*! consider giving a name to these results */
func f3() (int, int, error) {
	return 0, 0, nil
}

/*! consider giving a name to these results */
func f4() (int, int, error) {
	return 0, 0, nil
}

/*! consider giving a name to these results */
func f5() (int, float32, bool) {
	return 0, 0, false
}

/*! consider giving a name to these results */
func f6() (bool, bool) {
	return false, false
}

/*! consider giving a name to these results */
func f7() (int, float32, *foo) {
	return 0, 0, nil
}

/*! consider giving a name to these results */
func (f *foo) f8() (bool, bool) {
	return false, false
}

/*! consider giving a name to these results */
func (f *foo) f9() (bool, func() int) {
	return false, nil
}

/*! consider giving a name to these results */
func f10() (int, int, float64, float64) {
	return 0, 0, 0, 0
}

/*! consider giving a name to these results */
func doubleError() (error, error) {
	return nil, nil
}

type namedInt int

type namedStruct struct{}

/*! consider giving a name to these results */
func named2ptr() (namedInt, *namedInt) {
	return 0, nil
}

/*! consider giving a name to these results */
func named2() (namedInt, namedInt) {
	return 0, 0
}

/*! consider giving a name to these results */
func named3() (namedInt, namedStruct, namedInt) {
	return 0, namedStruct{}, 0
}

/*! consider giving a name to these results */
func namedAndPrimitive() (namedInt, int) {
	return 0, 0
}
