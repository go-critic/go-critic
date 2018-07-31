package checker_test

func terseAssignments(x, y int, z *int) {
	var o object

	x *= 2
	x += (x * 2)
	x -= (y - y)
	y &= 1
	y |= 2
	y ^= y
	y <<= 3
	y >>= uint(x)
	y &^= (1 << 10)

	for {
		o.count /= 2
		o.count %= 2
		*z %= 2
	}
}
