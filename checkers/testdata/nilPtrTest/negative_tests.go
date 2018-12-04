package checker_test

func noNilPtrDeref(x *int, obj *withPointer, arr *[5]int) {
	if x != nil {
		_ = *x
	}

	if obj != nil {
		_ = obj.ptr
	}

	if obj.ptr != nil {
		_ = *obj.ptr
	}

	if arr != nil {
		_ = arr[0]
	}
	if arr != nil {
		arr[0] = 0
	}
	if arr != nil {
		for _, x := range arr {
			_ = x
		}
	}

	if x != nil {
		if x != nil {
			_ = 10 + *x
		}
	}

	y := x

	if y != nil {
		for i := range arr {
			arr[i] = *y
		}
	}
}

func newObject() *withPointer {
	return &withPointer{}
}

func nilPtrAssigned(x *int, obj *withPointer, arr *[5]int) {
	if x == nil {
		if true {
			x = new(int)
		}
		_ = *x
	}

	if obj == nil {
		for {
			obj = newObject()
			_ = *obj
			break
		}
		_ = *obj
	}

	if arr == nil {
		arr = new([5]int)
		_ = *arr
		_ = arr[0]
		arr[0] = 1
	}

	if arr == nil {
		// TODO(Quasilyte): The order doesn't matter right now.
		// Could be fixed in future.
		_ = *arr
		_ = arr[0]
		arr[0] = 1
		arr = new([5]int)
	}

	if obj.ptr == nil {
		if obj.ptr != nil {
			// Even if this branch is impossible without data races,
			// we're conservative and consider initialization below
			// as possible.
			obj.ptr = new(int)
		}
		_ = *obj.ptr
	}

	if arr == nil {
		indirect := &arr
		_ = *arr
		_ = indirect
	}

	if x == nil {
		intPtrPtr(&x)
		_ = *x
	}
}

func intPtrPtr(x **int) {
	*x = new(int)
}
