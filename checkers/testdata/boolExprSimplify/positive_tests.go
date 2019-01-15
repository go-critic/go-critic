package checker_test

var (
	x, y, z bool
)

func combineChecks() {
	var x, y int

	/*! can simplify `x > y || x == y` to `x >= y` */
	_ = x > y || x == y
	/*! can simplify `x == y || x > y` to `x >= y` */
	_ = x == y || x > y

	/*! can simplify `(x > y) || (x == y)` to `x >= y` */
	_ = (x > y) || (x == y)
	/*! can simplify `(x == y) || (x > y)` to `x >= y` */
	_ = (x == y) || (x > y)

	/*! can simplify `x < y || x == y` to `x <= y` */
	_ = x < y || x == y
	/*! can simplify `x == y || x < y` to `x <= y` */
	_ = x == y || x < y

	/*! can simplify `(x < y) || (x == y)` to `x <= y` */
	_ = (x < y) || (x == y)
	/*! can simplify `(x == y) || (x < y)` to `x <= y` */
	_ = (x == y) || (x < y)
}

func doubleNegation() {
	/*! can simplify `!!x` to `x` */
	_ = !!x

	/*! can simplify `!!!x` to `!x` */
	_ = !!!x

	/*! can simplify `!!!!x` to `x` */
	_ = !!!!x

	/*! can simplify `!!!!!x` to `!x` */
	_ = !!!!!x

	/*! can simplify `!(!x)` to `x` */
	_ = !(!x)

	/*! can simplify `!(!(!(!(x))))` to `x` */
	_ = !(!(!(!(x))))
}

func negatedEquals() {
	/*! can simplify `!(x) == !(y)` to `(x) == (y)` */
	_ = !(x) == !(y)

	/*! can simplify `!x == !x == !x` to `x == x == !x` */
	_ = !x == !x == !x

	// TODO: should probably simplify other 2 expressions as well.
	/*! can simplify `!x == !y == !x == !y` to `x == y == !x == !y` */
	_ = !x == !y == !x == !y
}

func combined() {
	/*! can simplify `!(!!x == y)` to `x != y` */
	_ = !(!!x == y)

	{
		x := 1
		y := 2
		z := 3

		/*! can simplify `!(x > y) == !!!(y < z)` to `x <= y == (y >= z)` */
		_ = !(x > y) == !!!(y < z)

		/*! can simplify `!(x >= y+1)` to `x <= y` */
		_ = !(x >= y+1)
	}
}

func invertComparison() {
	/*! can simplify `!(x == y)` to `x != y` */
	_ = !(x == y)

	/*! can simplify `!((x || y) == (z && x))` to `(x || y) != (z && x)` */
	_ = !((x || y) == (z && x))

	/*! can simplify `!(x != y)` to `x == y` */
	_ = !(x != y)

	/*! can simplify `!((x || y) != (z && x))` to `(x || y) == (z && x)` */
	_ = !((x || y) != (z && x))

	{
		x := 1
		y := 2
		z := 3

		/*! can simplify `!(x < y)` to `x >= y` */
		_ = !(x < y)

		/*! can simplify `!((x + y) < (z - x))` to `(x + y) >= (z - x)` */
		_ = !((x + y) < (z - x))

		/*! can simplify `!(x > y)` to `x <= y` */
		_ = !(x > y)

		/*! can simplify `!((x + y) > (z - x))` to `(x + y) <= (z - x)` */
		_ = !((x + y) > (z - x))

		/*! can simplify `!(x <= y)` to `x > y` */
		_ = !(x <= y)

		/*! can simplify `!((x + y) <= (z - x))` to `(x + y) > (z - x)` */
		_ = !((x + y) <= (z - x))

		/*! can simplify `!(x >= y)` to `x < y` */
		_ = !(x >= y)

		/*! can simplify `!(!((x + y) >= (z - x)))` to `(x + y) >= (z - x)` */
		_ = !(!((x + y) >= (z - x)))
	}
}

func insideParens() {
	var x, y int

	/*! can simplify `!(x >= y)` to `x < y` */
	_ = (!(x >= y))
}

func returnsBool(f func()) bool { return false }

func insideLambda() {
	var x, y, z int

	_ = returnsBool(func() {
		/*! can simplify `!(x >= y)` to `x < y` */
		_ = !(x >= y)
	})

	_ = returnsBool(func() {
		/*! can simplify `!(x >= y)` to `x < y` */
		_ = !(x >= y)
		/*! can simplify `!(!((x + y) >= (z - x)))` to `(x + y) >= (z - x)` */
		_ = !(!((x + y) >= (z - x)))
	})
}

func removeIncDec(x, y, z int) {
	// `token.LSS`
	/*! can simplify `x < y+1` to `x <= y` */
	_ = x < y+1
	/*! can simplify `x+z < x+y+1` to `x+z <= x+y` */
	_ = x+z < x+y+1
	/*! can simplify `x-1 < y` to `x <= y` */
	_ = x-1 < y

	// `token.LEQ`
	/*! can simplify `x+2 <= z-1` to `x+2 < z` */
	_ = x+2 <= z-1
	/*! can simplify `x+z*2 <= x+y-1` to `x+z*2 < x+y` */
	_ = x+z*2 <= x+y-1
	/*! can simplify `x+1 <= y` to `x < y` */
	_ = x+1 <= y

	// `token.GTR`
	/*! can simplify `x+1 > y` to `x >= y` */
	_ = x+1 > y
	/*! can simplify `x > y-1` to `x >= y` */
	_ = x > y-1

	// `token.GEQ`
	/*! can simplify `x-1 >= y` to `x > y` */
	_ = x-1 >= y
	/*! can simplify `x >= y+1` to `x > y` */
	_ = x >= y+1
}

func foldRanges(x, y int) {
	/*! can simplify `x > 10 && x < 12` to `x == 11` */
	_ = x > 10 && x < 12
	/*! can simplify `x >= 11 && x < 12` to `x == 11` */
	_ = x >= 11 && x < 12
	/*! can simplify `x > 10 && x <= 11` to `x == 11` */
	_ = x > 10 && x <= 11
	/*! can simplify `x >= 11 && x <= 11` to `x == 11` */
	_ = x >= 11 && x <= 11

	/*! can simplify `x < 11 || x > 11` to `x != 11` */
	_ = x < 11 || x > 11
	/*! can simplify `x <= 10 || x > 11` to `x != 11` */
	_ = x <= 10 || x > 11
	/*! can simplify `x < 11 || x >= 12` to `x != 11` */
	_ = x < 11 || x >= 12
	/*! can simplify `x <= 10 || x >= 12` to `x != 11` */
	_ = x <= 10 || x >= 12
}
