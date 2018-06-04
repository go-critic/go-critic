package checker_test

func warnings() {
	{
		var xs [256]byte
		/// copy of xs (256 bytes) can be avoided with &xs
		for _, x := range xs {
			_ = x
		}
	}

	{
		var foo struct {
			arr [400]byte
		}
		/// copy of foo.arr (400 bytes) can be avoided with &foo.arr
		for _, x := range foo.arr {
			_ = x
		}
	}

	{
		xsList := make([][512]byte, 1)
		/// copy of xsList[0] (512 bytes) can be avoided with &xsList[0]
		for _, x := range xsList[0] {
			_ = x
		}
	}
}

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
