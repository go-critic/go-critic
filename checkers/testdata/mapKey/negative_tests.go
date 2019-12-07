package checker_test

func negativeTests() {
	// Non-string keys -- OK.
	_ = map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	// No spaces -- OK.
	_ = map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	_ = map[string]int{}

	_ = map[string]int{"": 1}

	// Single element.
	_ = map[string]int{
		"foo ": 1,
	}

	// Tests below check that we check underlying maps as well.

	type myMap map[string]int

	// More than 1 space.
	_ = myMap{
		"a ": 1,
		"b ": 2,
		"c ": 3,
	}

	// More than 1 space.
	_ = myMap{
		" a": 1,
		" b": 2,
		" c": 3,
	}

	// More than 1 space.
	_ = myMap{
		" a ": 1,
		" b ": 2,
		" c ": 3,
	}

	// Single-value map is an exception.
	_ = myMap{
		"a  ": 1,
	}

	// More than 1 space is not suspicious.
	_ = myMap{
		"a  ": 1,
		"b  ": 2,
	}

	// A whitespace itself is not suspicious.
	_ = myMap{
		"a": 1,
		" ": 2,
	}

	_ = myMap{
		"a":  1,
		"a ": 2,
		"  ": 3,
	}

	// Function call inside a key disables our duplicated key check.
	_ = map[string]int{
		getKeys()[0]: 1,
		getKeys()[1]: 2,
		getKeys()[0]: 3,
	}
}

func getKeys() []string {
	return []string{"a", "b"}
}
