package checker_test

func suspiciousWhitespace() {
	_ = map[string]int{
		/*! suspicious whitespace in `foo ` key */
		`foo `: 1,
		`bar`:  2,
		`baz`:  3,
	}

	_ = map[string]int{
		"foo": 1,
		"bar": 2,
		/*! suspicious whitespace in " baz" key */
		" baz": 3,
	}

	type myMap map[string]int

	_ = myMap{
		"foo": 1,
		/*! suspicious whitespace in "bar " key */
		"bar ": 2,
		"baz":  3,
	}
}

func suspiciousDupKey() {
	k := "abc"

	keys := []string{"a", "b"}

	_ = map[string]int{
		k: 1,
		/*! suspicious duplicate k key */
		k: 2,
	}

	_ = map[string]int{
		keys[0]: 1,
		keys[1]: 2,
		/*! suspicious duplicate keys[0] key */
		keys[0]: 3,
	}
}
