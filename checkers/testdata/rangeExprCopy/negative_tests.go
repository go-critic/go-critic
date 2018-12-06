package checker_test

func returnArray() [20]int {
	return [20]int{}
}

func noWarnings() {
	// OK: returned valus is not addressible, can't take address.
	for _, x := range returnArray() {
		_ = x
	}

	{
		var xs [200]byte
		// OK: already iterating over a pointer.
		for _, x := range &xs {
			_ = x
		}
		// OK: only index is used. No copy is generated.
		for i := range xs {
			_ = xs[i]
		}
		// OK: like in case above, no copy, so it's OK.
		for range xs {
		}
	}

	{
		var xs [10]byte
		// OK: xs is a very small array that can be trivially copied.
		for _, x := range xs {
			_ = x
		}
	}
}
