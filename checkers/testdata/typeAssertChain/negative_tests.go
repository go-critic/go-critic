package checker_test

func negativeTests() {
	var x interface{}

	switch v := x.(type) {
	case int8:
		_ = v
	case int16:
		_ = v
	}

	switch v := x.(type) {
	case int8:
		_ = v
	case int16:
		_ = v
	case int32:
		_ = v
	}

	// Not a type assertion chain.
	if true {
	} else if true {
	} else if false {
	}

	// Only a single type assertion.
	if v, ok := x.(int8); ok {
		_ = v
	}

	// Duplicated types.
	if v, ok := x.(int8); ok {
		_ = v
	} else if v, ok := x.(int8); ok {
		_ = v
	}

	// Non-matching condition.
	if v, ok := x.(int8); ok {
		_ = v
	} else if v, ok := x.(int8); true {
		_ = v
		_ = ok
	}
	if v, ok := x.(int8); ok {
		_ = v
	} else if v, ok := x.(int8); !ok {
		_ = v
	}

	var y interface{}
	// Mixed type-asserted values.
	if v1, ok := x.(int8); ok {
		_ = v1
	} else if v2, ok := x.(int16); ok {
		_ = v2
	} else if v3, ok := y.(int32); ok {
		_ = v3
	}
}
