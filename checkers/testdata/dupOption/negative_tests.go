package checker_test

func doSome2(w, h int) {
	_ = newPanel("hello",
		withWidth(w),
		withHeight(h),
	)

	export(nil,
		withImage(nil, ""),
	)
}
