package checker_tests

import (
	"time"
)

func returnExpr() {
	// The value is 2 (colorGreen), but it's an expression,
	// not a raw value, so the programmer intention can be
	// more intricate. To avoid false positives, leave these alone.
	var _ color = 1 + 1
	var _ color = colRed + colRed
}

func nonExistingValues() {
	var _ color = 423
	var _ color = 10 + 132
}

func constexpr() {
	// This kind of usage is idiomatic.
	// We should not suggest writing `time.Nanosecond * time.Second` here.
	var x time.Duration = 1 * time.Second
	var _ time.Duration = x / 100
}
