package checker_test

func noWarnings() {
	// OK: complexity lower than a threshold.
	_ = 1 + 2 + 3
	_ = 1 + 2 + 3
	_ = 1 + 2 + 3
	_ = 1 + 2 + 3

	// OK: complexity lower than a threshold.
	_ = uintptr(1 + 2)
	_ = uintptr(1 + 2)
	_ = uintptr(1 + 2)
	_ = uintptr(1 + 2)

	// OK: no warnings for function calls because they may have side-effects.
	_ = add1(1 + 2 + 3)
	_ = add1(1 + 2 + 3)
	_ = add1(1 + 2 + 3)
	_ = add1(1 + 2 + 3)

	// OK: different access paths.
	switch b := (t{}); 2 {
	case b.a.c.d.e.p1:
	case b.a.c1.d.e.p1:
	case b.a.c2.d.e.p1:
	}
}
