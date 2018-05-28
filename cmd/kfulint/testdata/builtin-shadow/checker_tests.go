package checker_test

func noWarnigs() {
	var foo struct {
		len int
		cap int
	}

	foo.len = 123
	foo.cap = 321

	foo.len, foo.cap = foo.cap, foo.len
}

func assigningToBuiltinIdentifiers() {
	// Types:
	/// assigning to predeclared identifier: bool
	bool := "1"
	_ = bool

	/// assigning to predeclared identifier: byte
	byte := "1"
	_ = byte

	/// assigning to predeclared identifier: complex64
	complex64 := "1"
	_ = complex64

	/// assigning to predeclared identifier: complex128
	complex128 := "1"
	_ = complex128

	/// assigning to predeclared identifier: error
	error := "1"
	_ = error

	/// assigning to predeclared identifier: float32
	float32 := "1"
	_ = float32

	/// assigning to predeclared identifier: float64
	float64 := "1"
	_ = float64

	/// assigning to predeclared identifier: int
	int := "1"
	_ = int

	/// assigning to predeclared identifier: int8
	int8 := "1"
	_ = int8

	/// assigning to predeclared identifier: int16
	int16 := "1"
	_ = int16

	/// assigning to predeclared identifier: int32
	int32 := "1"
	_ = int32

	/// assigning to predeclared identifier: int64
	int64 := "1"
	_ = int64

	/// assigning to predeclared identifier: rune
	rune := 10
	_ = rune

	/// assigning to predeclared identifier: string
	string := 10
	_ = string

	/// assigning to predeclared identifier: uint
	uint := "1"
	_ = uint

	/// assigning to predeclared identifier: uint8
	uint8 := "1"
	_ = uint8

	/// assigning to predeclared identifier: uint16
	uint16 := "1"
	_ = uint16

	/// assigning to predeclared identifier: uint32
	uint32 := "1"
	_ = uint32

	/// assigning to predeclared identifier: uint64
	uint64 := "1"
	_ = uint64

	/// assigning to predeclared identifier: uintptr
	uintptr := "1"
	_ = uintptr

	// Constants:
	/// assigning to predeclared identifier: true
	true := "1"
	_ = true

	/// assigning to predeclared identifier: false
	false := "1"
	_ = false

	/// assigning to predeclared identifier: iota
	iota := "1"
	_ = iota

	// Zero value:
	/// assigning to predeclared identifier: nil
	nil := "1"
	_ = nil

	// Functions:
	/// assigning to predeclared identifier: append
	append := 1
	_ = append

	/// assigning to predeclared identifier: cap
	cap := 1
	_ = cap

	/// assigning to predeclared identifier: close
	close := 1
	_ = close

	/// assigning to predeclared identifier: complex
	complex := 1
	_ = complex

	/// assigning to predeclared identifier: copy
	copy := 1
	_ = copy

	/// assigning to predeclared identifier: delete
	delete := 1
	_ = delete

	/// assigning to predeclared identifier: imag
	imag := 1
	_ = imag

	/// assigning to predeclared identifier: len
	len := 1
	_ = len

	/// assigning to predeclared identifier: make
	make := 1
	_ = make

	/// assigning to predeclared identifier: new
	new := 1
	_ = new

	/// assigning to predeclared identifier: panic
	panic := 1
	_ = panic

	/// assigning to predeclared identifier: print
	print := 1
	_ = print

	/// assigning to predeclared identifier: println
	println := 1
	_ = println

	/// assigning to predeclared identifier: real
	real := 1
	_ = real

	/// assigning to predeclared identifier: recover
	recover := 1
	_ = recover
}
