package checker_test

func suggestTypeSwitch() {
	var x interface{}

	/*! rewrite if-else to type switch statement */
	if v, ok := x.(int8); ok {
		_ = v
	} else if v, ok := x.(int16); ok {
		_ = v
	}

	/*! rewrite if-else to type switch statement */
	if v, ok := x.(int8); ok {
		_ = v
	} else if v, ok := x.(int16); ok {
		_ = v
	} else if v, ok := x.(int32); ok {
		_ = v
	}

	/*! rewrite if-else to type switch statement */
	if v1, ok := x.(int8); ok {
		_ = v1
	} else if v2, ok := x.(int16); ok {
		_ = v2
	} else if v3, ok := x.(int32); ok {
		_ = v3
	}
}
