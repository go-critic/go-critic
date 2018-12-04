package checker_test

type withPointer struct {
	ptr *int
}

func nilPtrDeref(x *int, obj *withPointer, arr *[5]int) {
	/*! nil ptr deref possible; probably `x != nil` was intended */
	if x == nil {
		_ = *x
	}

	// TODO(Quasilyte): warn on implicit dereferences as well.
	if obj == nil {
		_ = obj.ptr
	}

	/*! nil ptr deref possible; probably `obj.ptr != nil` was intended */
	if obj.ptr == nil {
		_ = *obj.ptr
	}

	// TODO(Quasilyte): warn on this, too.
	if arr == nil {
		_ = arr[0]
	}
	// TODO(Quasilyte): and on this one, too.
	if arr == nil {
		arr[0] = 0
	}
	// TODO(Quasilyte): range over nil array is also an implicit deref.
	// But only if key+value are present (so there is an array copy involved).
	if arr == nil {
		for _, x := range arr {
			_ = x
		}
	}

	if x != nil {
		/*! nil ptr deref possible; probably `x != nil` was intended */
		if x == nil {
			_ = 10 + *x
		}
	}

	y := x

	/*! nil ptr deref possible; probably `y != nil` was intended */
	if y == nil {
		for i := range arr {
			arr[i] = *y
		}
	}
}
