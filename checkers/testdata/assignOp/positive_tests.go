package checker_test

type object struct {
	count int
}

func verboseAssignments(x, y int, z *int) {
	var o object

	/*! replace `x = x * 2` with `x *= 2` */
	x = x * 2

	/*! replace `x = x + (x * 2)` with `x += (x * 2)` */
	x = x + (x * 2)

	/*! replace `x = x - (y - y)` with `x -= (y - y)` */
	x = x - (y - y)

	/*! replace `y = y & 1` with `y &= 1` */
	y = y & 1
	/*! replace `y = y | 2` with `y |= 2` */
	y = y | 2
	/*! replace `y = y ^ y` with `y ^= y` */
	y = y ^ y
	/*! replace `y = y << 3` with `y <<= 3` */
	y = y << 3
	/*! replace `y = y >> uint(x)` with `y >>= uint(x)` */
	y = y >> uint(x)
	/*! replace `y = y &^ (1 << 10)` with `y &^= (1 << 10)` */
	y = y &^ (1 << 10)
	/*! replace `y = y + 1` with `y++` */
	y = y + 1
	/*! replace `y = y - 1` with `y--` */
	y = y - 1

	for {
		/*! replace `o.count = o.count / 2` with `o.count /= 2` */
		o.count = o.count / 2
		/*! replace `*z = *z % 2` with `*z %= 2` */
		*z = *z % 2
	}
}
