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
