package checker_test

import (
	"unsafe"
)

type myInt int

var intVar int

func sizeofBasicLit() {
	/// can replace unsafe.Sizeof(0) with math/bits.UintSize
	_ = unsafe.Sizeof(0)
	/// can replace unsafe.Sizeof(0xff) with math/bits.UintSize
	_ = unsafe.Sizeof(0xff)
}

func sizeofCallExpr() {
	/// can replace unsafe.Sizeof(uint(1)) with math/bits.UintSize
	_ = unsafe.Sizeof(uint(1))
	/// can replace unsafe.Sizeof(int(2)) with math/bits.UintSize
	_ = unsafe.Sizeof(int(2))
}

func maxValue() {
	/// can replace 1<<7 - 1 with math.MaxInt8
	_ = 1<<7 - 1
	/// can replace -1 << 7 with math.MinInt8
	_ = -1 << 7
	/// can replace 1<<15 - 1 with math.MaxInt16
	_ = 1<<15 - 1
	/// can replace -1 << 15 with math.MinInt16
	_ = -1 << 15
	/// can replace 1<<31 - 1 with math.MaxInt32
	_ = 1<<31 - 1
	/// can replace -1 << 31 with math.MinInt32
	_ = -1 << 31
	/// can replace 1<<63 - 1 with math.MaxInt64
	_ = 1<<63 - 1
	/// can replace -1 << 63 with math.MinInt64
	_ = -1 << 63
	/// can replace 1<<8 - 1 with math.MaxUint8
	_ = 1<<8 - 1
	/// can replace 1<<16 - 1 with math.MaxUint16
	_ = 1<<16 - 1
	/// can replace 1<<32 - 1 with math.MaxUint32
	_ = 1<<32 - 1
	/// can replace 1<<64 - 1 with math.MaxUint64
	var _ uint64 = 1<<64 - 1

	// Same cases, but with extra parenthesis.

	/// can replace 1<<7 - 1 with math.MaxInt8
	_ = (1<<7 - 1)
	/// can replace -1 << 7 with math.MinInt8
	_ = (-1 << 7)
	/// can replace 1<<15 - 1 with math.MaxInt16
	_ = (1<<15 - 1)
	/// can replace -1 << 15 with math.MinInt16
	_ = (-1 << 15)
	/// can replace 1<<31 - 1 with math.MaxInt32
	_ = (1<<31 - 1)
	/// can replace -1 << 31 with math.MinInt32
	_ = (-1 << 31)
	/// can replace 1<<63 - 1 with math.MaxInt64
	_ = (1<<63 - 1)
	/// can replace -1 << 63 with math.MinInt64
	_ = (-1 << 63)
	/// can replace 1<<8 - 1 with math.MaxUint8
	_ = (1<<8 - 1)
	/// can replace 1<<16 - 1 with math.MaxUint16
	_ = (1<<16 - 1)
	/// can replace 1<<32 - 1 with math.MaxUint32
	_ = (1<<32 - 1)
	/// can replace 1<<64 - 1 with math.MaxUint64
	var _ uint64 = (1<<64 - 1)
}

func mathPi() {
	// Not precise enough to trigger a warning:
	_ = 3.1415
	// But this is a special (common) case:
	/// can replace 3.14 with math.Pi
	_ = 3.14

	/// can replace 3.14159 with math.Pi
	_ = 3.14159
	/// can replace 3.141592653589793 with math.Pi
	_ = 3.141592653589793
	/// can replace 3.14159265358979323846264338327950288419716939937510582097494459 with math.Pi
	_ = 3.14159265358979323846264338327950288419716939937510582097494459
}

func mathE() {
	// But this is a special (common) case:
	/// can replace 2.71 with math.E
	_ = 2.71

	/// can replace 2.7182818284590452 with math.E
	_ = 2.7182818284590452
	/// can replace 2.71828182845904523536028747135266249775724709369995957496696763 with math.E
	_ = 2.71828182845904523536028747135266249775724709369995957496696763
}

func mathConsts() {
	// Less common math consts.

	/// can replace 1.4142135623730950488 with math.Sqrt2
	_ = 1.4142135623730950488
	/// can replace 1.6487212707001281468486507878 with math.SqrtE
	_ = 1.6487212707001281468486507878
	/// can replace 1.77245385090 with math.SqrtPi
	_ = 1.77245385090
	/// can replace 1.2720196 with math.SqrtPhi
	_ = 1.2720196

	/// can replace 0.693147180559945309 with math.Ln2
	_ = 0.693147180559945309
	/// can replace 2.30258509299404568 with math.Ln10
	_ = 2.30258509299404568
}

func httpMethod() {
	/// can replace "GET" with net/http.MethodGet
	_ = "GET"
	/// can replace "HEAD" with net/http.MethodHead
	_ = "HEAD"
	/// can replace "POST" with net/http.MethodPost
	_ = "POST"
	/// can replace "PUT" with net/http.MethodPut
	_ = "PUT"
	/// can replace "DELETE" with net/http.MethodDelete
	_ = "DELETE"
}

func timeStrings() {
	/// can replace "Mon Jan _2 15:04:05 MST 2006" with time.UnixDate
	_ = "Mon Jan _2 15:04:05 MST 2006"
	/// can replace "3:04PM" with time.Kitchen
	_ = "3:04PM"
}
