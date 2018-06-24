package checker_tests

import (
	"time"
)

type color int

const (
	colDefault color = iota
	colRed
	colGreen
	colBlue
)

func returnRaw1() color {
	/// use colDefault instead of 0
	return 0
}

func returnRaw2() time.Duration {
	/// use time.Nanosecond instead of 1
	return 1
}

func returnRaw4() color {
	/// use colRed instead of 1
	return 1
}

func conversions() {
	// TODO: may want to include whole conversion to warning message.
	// So, instead of 0, report color(0).

	/// use colDefault instead of 0
	_ = color(0)
	/// use colRed instead of 1
	_ = color(1)
	/// use colGreen instead of 2
	_ = color(2)
	/// use colBlue instead of 3
	_ = color(3)
}

func comparisons(x color) {
	/// use colDefault instead of 0
	if x == 0 {
		panic("default?")
	}

	/// use colDefault instead of 0
	if true || x == 0 {
	}

	switch x {
	/// use colRed instead of 1
	case 1:
		panic("red?")
	/// use colGreen instead of 2
	case 2:
		panic("green?")
	}
}
