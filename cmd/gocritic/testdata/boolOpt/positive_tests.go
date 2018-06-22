package checker_tests

var (
	x, y, z bool
)

func doubleNegation() {
	/// can simplify `!!x` to `x`
	_ = !!x

	/// can simplify `!!!x` to `!x`
	_ = !!!x

	/// can simplify `!!!!x` to `x`
	_ = !!!!x

	/// can simplify `!!!!!x` to `!x`
	_ = !!!!!x

	/// can simplify `!(!x)` to `x`
	_ = !(!x)

	/// can simplify `!(!(!(!(x))))` to `x`
	_ = !(!(!(!(x))))
}

func negatedEquals() {
	/// can simplify `!(x) == !(y)` to `(x) == (y)`
	_ = !(x) == !(y)

	/// can simplify `!x == !x == !x` to `x == x == !x`
	_ = !x == !x == !x

	// TODO: should probably simplify other 2 expressions as well.
	/// can simplify `!x == !y == !x == !y` to `x == y == !x == !y`
	_ = !x == !y == !x == !y
}

func combined() {
	/// can simplify `!(!!x == y)` to `x != y`
	_ = !(!!x == y)

	{
		x := 1
		y := 2
		z := 3

		/// can simplify `!(x > y) == !!!(y < z)` to `x <= y == (y >= z)`
		_ = !(x > y) == !!!(y < z)
	}
}

func invertComparison() {
	/// can simplify `!(x == y)` to `x != y`
	_ = !(x == y)

	/// can simplify `!((x || y) == (z && x))` to `(x || y) != (z && x)`
	_ = !((x || y) == (z && x))

	/// can simplify `!(x != y)` to `x == y`
	_ = !(x != y)

	/// can simplify `!((x || y) != (z && x))` to `(x || y) == (z && x)`
	_ = !((x || y) != (z && x))

	{
		x := 1
		y := 2
		z := 3

		/// can simplify `!(x < y)` to `x >= y`
		_ = !(x < y)

		/// can simplify `!((x + y) < (z - x))` to `(x + y) >= (z - x)`
		_ = !((x + y) < (z - x))

		/// can simplify `!(x > y)` to `x <= y`
		_ = !(x > y)

		/// can simplify `!((x + y) > (z - x))` to `(x + y) <= (z - x)`
		_ = !((x + y) > (z - x))

		/// can simplify `!(x <= y)` to `x > y`
		_ = !(x <= y)

		/// can simplify `!((x + y) <= (z - x))` to `(x + y) > (z - x)`
		_ = !((x + y) <= (z - x))

		/// can simplify `!(x >= y)` to `x < y`
		_ = !(x >= y)

		/// can simplify `!(!((x + y) >= (z - x)))` to `(x + y) >= (z - x)`
		_ = !(!((x + y) >= (z - x)))
	}
}
