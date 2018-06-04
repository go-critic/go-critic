package checker_test

import (
	"math"
	"math/bits"
	"unsafe"
)

func returnZero() int { return 0 }

func sizeofCallExprN() {
	// Arbitrary int expressions should not trigger warnings,
	_ = unsafe.Sizeof(returnZero())
	_ = unsafe.Sizeof(myInt(0))
	_ = unsafe.Sizeof(intVar)
}

func maxValueN() {
	// Literals should not trigger warnings because they're
	// frequently used to specify bit masks.
	_ = 0xff
	_ = 0xffff
}

func mathPiN() {
	// Not precise enough to trigger a warning:
	_ = 3.1415
}

func mathEN() {
	// Not precise enough to trigger a warning:
	_ = 2.7182
}

func baitFalsePositives() {
	// No warnings should occur if constatns are used.

	// Note: avoid importing packages that have many dependencies
	// because it will slow testing down significantly.
	// For example, importing net/http can make test run a few seconds
	// longer than without it.

	_ = math.MaxInt32
	_ = math.MaxInt16
	_ = math.Pi
	_ = math.E
	_ = bits.UintSize
}
