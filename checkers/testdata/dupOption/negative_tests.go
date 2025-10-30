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

func f(a int, b string, c ...float64) {}

func g() (int, string, float64) {
	return 42, "hello", 3.14
}

func h() (int, string) {
	return 42, "hello"
}

func _() {
	f(g())
	f(h())
}
