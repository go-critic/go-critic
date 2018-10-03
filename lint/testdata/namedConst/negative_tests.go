package checker_test

import (
	"time"
)

const (
	untyped1 = "one"
	untyped2 = "two"
)

type globalType int32

func localConsts() {
	// May consider checking for local const defs,
	// but without proper index, it's going to be notoriously slow.

	type localType int
	const (
		l1 localType  = 10
		l2 globalType = 20
	)
	var _ localType = 10
	var _ globalType = 20
}

func untypedConsts() {
	_ = "one"
	_ = "two"
	_ = 0
}

func expression() {
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
