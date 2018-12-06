package checker_test

/*! shadowing of predeclared identifier: len */
/*! shadowing of predeclared identifier: cap */
/*! shadowing of predeclared identifier: complex64 */
func warnParams(len int, cap string) (complex64 complex128) {
	return 0
}

type shadower struct{}

/*! shadowing of predeclared identifier: bool */
func (bool shadower) shadowThem() {}

func genDeclShadow() {
	/*! shadowing of predeclared identifier: byte */
	var byte = 1
	var (
		/*! shadowing of predeclared identifier: true */
		true = false
		/*! shadowing of predeclared identifier: false */
		false = 10
	)
	/*! shadowing of predeclared identifier: error */
	const error = "error"
	const (
		/*! shadowing of predeclared identifier: float32 */
		float32 = 32
		/*! shadowing of predeclared identifier: float64 */
		float64 = 64
	)

	_ = byte
	_ = true
	_ = false
}

func assigningToBuiltinIdentifiers() {
	// Types:
	/*! shadowing of predeclared identifier: bool */
	bool := "1"
	_ = bool

	/*! shadowing of predeclared identifier: byte */
	byte := "1"
	_ = byte

	/*! shadowing of predeclared identifier: complex64 */
	complex64 := "1"
	_ = complex64

	/*! shadowing of predeclared identifier: complex128 */
	complex128 := "1"
	_ = complex128

	/*! shadowing of predeclared identifier: error */
	error := "1"
	_ = error

	/*! shadowing of predeclared identifier: float32 */
	float32 := "1"
	_ = float32

	/*! shadowing of predeclared identifier: float64 */
	float64 := "1"
	_ = float64

	/*! shadowing of predeclared identifier: int */
	int := "1"
	_ = int

	/*! shadowing of predeclared identifier: int8 */
	int8 := "1"
	_ = int8

	/*! shadowing of predeclared identifier: int16 */
	int16 := "1"
	_ = int16

	/*! shadowing of predeclared identifier: int32 */
	int32 := "1"
	_ = int32

	/*! shadowing of predeclared identifier: int64 */
	int64 := "1"
	_ = int64

	/*! shadowing of predeclared identifier: rune */
	rune := 10
	_ = rune

	/*! shadowing of predeclared identifier: string */
	string := 10
	_ = string

	/*! shadowing of predeclared identifier: uint */
	uint := "1"
	_ = uint

	/*! shadowing of predeclared identifier: uint8 */
	uint8 := "1"
	_ = uint8

	/*! shadowing of predeclared identifier: uint16 */
	uint16 := "1"
	_ = uint16

	/*! shadowing of predeclared identifier: uint32 */
	uint32 := "1"
	_ = uint32

	/*! shadowing of predeclared identifier: uint64 */
	uint64 := "1"
	_ = uint64

	/*! shadowing of predeclared identifier: uintptr */
	uintptr := "1"
	_ = uintptr

	// Constants:
	/*! shadowing of predeclared identifier: true */
	true := "1"
	_ = true

	/*! shadowing of predeclared identifier: false */
	false := "1"
	_ = false

	/*! shadowing of predeclared identifier: iota */
	iota := "1"
	_ = iota

	// Zero value:
	/*! shadowing of predeclared identifier: nil */
	nil := "1"
	_ = nil

	// Functions:
	/*! shadowing of predeclared identifier: append */
	append := 1
	_ = append

	/*! shadowing of predeclared identifier: cap */
	cap := 1
	_ = cap

	/*! shadowing of predeclared identifier: close */
	close := 1
	_ = close

	/*! shadowing of predeclared identifier: complex */
	complex := 1
	_ = complex

	/*! shadowing of predeclared identifier: copy */
	copy := 1
	_ = copy

	/*! shadowing of predeclared identifier: delete */
	delete := 1
	_ = delete

	/*! shadowing of predeclared identifier: imag */
	imag := 1
	_ = imag

	/*! shadowing of predeclared identifier: len */
	len := 1
	_ = len

	/*! shadowing of predeclared identifier: make */
	make := 1
	_ = make

	/*! shadowing of predeclared identifier: new */
	new := 1
	_ = new

	/*! shadowing of predeclared identifier: panic */
	panic := 1
	_ = panic

	/*! shadowing of predeclared identifier: print */
	print := 1
	_ = print

	/*! shadowing of predeclared identifier: println */
	println := 1
	_ = println

	/*! shadowing of predeclared identifier: real */
	real := 1
	_ = real

	/*! shadowing of predeclared identifier: recover */
	recover := 1
	_ = recover
}
