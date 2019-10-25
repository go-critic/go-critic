package checker_test

func bothTruncated(x int32, y int64) {
	_ = int16(x) < int16(y)
}

func cmpWithConst(x int64) {
	_ = int16(x) < 0
	_ = 0 < int16(x)
}

func sameSizes(x, y int16) {
	_ = int16(y) < x
	_ = x < int16(y)

	_ = y < x
	_ = x < y
}

func mixedSignedness1(x uint8, y int16) {
	_ = uint8(y) < x
	_ = x < uint8(y)
}

func mixedSignedness2(x int8, y uint16) {
	_ = int8(y) < x
	_ = x < int8(y)
}

func goodInt8(x int8, y int16) {
	_ = y < int16(x)
	_ = int16(x) < y

	_ = y <= int16(x)
	_ = int16(x) <= y

	_ = y > int16(x)
	_ = int16(x) > y

	_ = y >= int16(x)
	_ = int16(x) >= y

	_ = y == int16(x)
	_ = int16(x) == y

	_ = y != int16(x)
	_ = int16(x) != y
}

func goodInt16(x1 int8, x2 int16, y int32) {
	_ = y == int32(x1)
	_ = int32(x1) == y

	_ = y == int32(x2)
	_ = int32(x2) == y
}

func goodInt32(x1 int8, x2 int16, x3 int32, y int64) {
	_ = y == int64(x1)
	_ = int64(x1) == y

	_ = y == int64(x2)
	_ = int64(x2) == y

	_ = y == int64(x3)
	_ = int64(3) == y
}

func goodUint8(x uint8, y uint16) {
	_ = y == uint16(x)
	_ = uint16(x) == y
}
