package checker_test

func good() {
	{
		r := []uint64{10}
		_ = r[0]
		r2 := []rune{10, 12, 34}
		_ = r2[0]
	}

	{
		// OK: 'runes' is not string-typed.
		var runes []rune
		_ = []rune(runes)[0]
	}

	{
		// OK: not a 0 index.
		var s string
		_ = []rune(s)[1]
	}

	{
		// OK: let's allow using int32 for now?
		//
		// We could add a []int32($x) pattern, but it should
		// probably be a concern of other checker that would suggest
		// using rune type here if the context implies the rune type.
		// And when the type is changed to 'rune', we could report
		// preferDecodeRune in a more meaningful way.
		var s string
		_ = []int32(s)[0]
	}
}
