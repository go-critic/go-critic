package checker_test

func badInt8(x int8, y int16) {
	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = int8(y) < x
	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = x < int8(y)

	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = int8(y) <= x
	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = x <= int8(y)

	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = int8(y) > x
	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = x > int8(y)

	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = int8(y) >= x
	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = x >= int8(y)

	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = int8(y) == x
	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = x == int8(y)

	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = int8(y) != x
	/*! truncation in comparison 16->8 bit; cast the other operand to int16 instead */
	_ = x != int8(y)
}

func badInt16(x1 int8, x2 int16, y int32) {
	/*! truncation in comparison 32->8 bit; cast the other operand to int32 instead */
	_ = int8(y) == x1
	/*! truncation in comparison 32->8 bit; cast the other operand to int32 instead */
	_ = x1 == int8(y)

	/*! truncation in comparison 32->16 bit; cast the other operand to int32 instead */
	_ = int16(y) == x2
	/*! truncation in comparison 32->16 bit; cast the other operand to int32 instead */
	_ = x2 == int16(y)
}

func badInt32(x1 int8, x2 int16, x3 int32, y int64) {
	/*! truncation in comparison 64->8 bit; cast the other operand to int64 instead */
	_ = int8(y) == x1
	/*! truncation in comparison 64->8 bit; cast the other operand to int64 instead */
	_ = x1 == int8(y)

	/*! truncation in comparison 64->16 bit; cast the other operand to int64 instead */
	_ = int16(y) == x2
	/*! truncation in comparison 64->16 bit; cast the other operand to int64 instead */
	_ = x2 == int16(y)

	/*! truncation in comparison 64->32 bit; cast the other operand to int64 instead */
	_ = int32(y) == x3
	/*! truncation in comparison 64->32 bit; cast the other operand to int64 instead */
	_ = x3 == int32(y)
}

func badUint8(x uint8, y uint16) {
	/*! truncation in comparison 16->8 bit; cast the other operand to uint16 instead */
	_ = uint8(y) == x
	/*! truncation in comparison 16->8 bit; cast the other operand to uint16 instead */
	_ = x == uint8(y)
}

func badUint16(x1 uint8, x2 uint16, y uint32) {
	/*! truncation in comparison 32->8 bit; cast the other operand to uint32 instead */
	_ = uint8(y) == x1
	/*! truncation in comparison 32->8 bit; cast the other operand to uint32 instead */
	_ = x1 == uint8(y)

	/*! truncation in comparison 32->16 bit; cast the other operand to uint32 instead */
	_ = uint16(y) == x2
	/*! truncation in comparison 32->16 bit; cast the other operand to uint32 instead */
	_ = x2 == uint16(y)
}

func badUint32(x1 uint8, x2 uint16, x3 uint32, y uint64) {
	/*! truncation in comparison 64->8 bit; cast the other operand to uint64 instead */
	_ = uint8(y) == x1
	/*! truncation in comparison 64->8 bit; cast the other operand to uint64 instead */
	_ = x1 == uint8(y)

	/*! truncation in comparison 64->16 bit; cast the other operand to uint64 instead */
	_ = uint16(y) == x2
	/*! truncation in comparison 64->16 bit; cast the other operand to uint64 instead */
	_ = x2 == uint16(y)

	/*! truncation in comparison 64->32 bit; cast the other operand to uint64 instead */
	_ = uint32(y) == x3
	/*! truncation in comparison 64->32 bit; cast the other operand to uint64 instead */
	_ = x3 == uint32(y)
}
