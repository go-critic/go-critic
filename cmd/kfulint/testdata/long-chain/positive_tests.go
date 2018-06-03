package checker_test

import "unsafe"

type t struct {
	a struct {
		c struct {
			d struct {
				e struct {
					p1, p2, p3, p4 int
				}
			}
		}

		c1 struct {
			d struct {
				e struct {
					p1, p2, p3, p4 int
				}
			}
		}

		c2 struct {
			d struct {
				e struct {
					p1, p2, p3, p4 int
				}
			}
		}
	}
}

func add1(x int) int { return x + 1 }

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

func binaryExpr() {
	_ = add1(1 + 2 + 3 + 4)
	_ = add1(1 + 2 + 3 + 4)
	/// 1 + 2 + 3 + 4 repeated multiple times, consider assigning it to local variable
	_ = add1(1 + 2 + 3 + 4)
}

func indexExpr(x [][][][]int) {
	i := 0
	_ = x[0][1][2][i+0]
	_ = x[0][1][2][i+1]
	/// x[0][1][2] repeated multiple times, consider assigning it to local variable
	_ = x[0][1][2][i+2]
}

func selectorExpr() {
	b := t{}

	switch 1 {
	case b.a.c.d.e.p1:
	case b.a.c.d.e.p2:
	/// b.a.c.d.e repeated multiple times, consider assigning it to local variable
	case b.a.c.d.e.p3:
	case b.a.c.d.e.p4:
	}
}

func callExpr() {
	_ = uintptr(1 + 2 + 3 + 4)
	_ = uintptr(1 + 2 + 3 + 4)
	/// uintptr(1 + 2 + 3 + 4) repeated multiple times, consider assigning it to local variable
	_ = uintptr(1 + 2 + 3 + 4)

	_ = (int)(1 + 2 + 3 + 4)
	_ = (int)(1 + 2 + 3 + 4)
	/// (int)(1 + 2 + 3 + 4) repeated multiple times, consider assigning it to local variable
	_ = (int)(1 + 2 + 3 + 4)

	type myInt int

	_ = myInt(1) + myInt(2)
	_ = myInt(1) + myInt(2)
	/// myInt(1) + myInt(2) repeated multiple times, consider assigning it to local variable
	_ = myInt(1) + myInt(2)

	type intAlias = myInt

	_ = intAlias(1) + intAlias(2)
	_ = intAlias(1) + intAlias(2)
	/// intAlias(1) + intAlias(2) repeated multiple times, consider assigning it to local variable
	_ = intAlias(1) + intAlias(2)

	type struct1 struct{}
	type struct2 struct{}

	_ = struct2(struct1(struct2{}))
	_ = struct2(struct1(struct2{}))
	/// struct2(struct1(struct2{})) repeated multiple times, consider assigning it to local variable
	_ = struct2(struct1(struct2{}))

	var x int

	_ = uintptr(unsafe.Pointer(&x))
	_ = uintptr(unsafe.Pointer(&x))
	/// uintptr(unsafe.Pointer(&x)) repeated multiple times, consider assigning it to local variable
	_ = uintptr(unsafe.Pointer(&x))
}
