package checker_test

import (
	mytime "time"
)

type color int

const (
	colDefault color = iota
	colRed
	colGreen
	colBlue
)

type typename string

const (
	tnInt    typename = "int"
	tnString typename = "string"
)

func rawReturn() {
	_ = func() color {
		/// use colDefault instead of 0
		return 0
	}

	_ = func() mytime.Duration {
		/// use mytime.Nanosecond instead of 1
		return 1
	}

	_ = func() color {
		/// use colRed instead of 1
		return 1
	}

	_ = func() typename {
		/// use tnInt instead of "int"
		return "int"
	}

	_ = func() typename {
		/// use tnInt instead of "int"
		/// use tnString instead of "string"
		return "int" + typename("/") + "string"
	}
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
