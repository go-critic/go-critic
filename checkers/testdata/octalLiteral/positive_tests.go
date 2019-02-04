package checker_test

import (
	"math"
)

func calculateInt(x int) int {
	return x
}

func calculateIntPair(x, y int) (int, int) {
	return x, y
}

func calculateManyArgs(x int, s string, y int) (int, string, int) {
	return x, s, y
}

func warningsCalc() {
	/*! suspicious octal args in `calculateInt(00)` */
	_ = calculateInt(00)

	/*! suspicious octal args in `calculateInt(+01)` */
	_ = calculateInt(+01)

	/*! suspicious octal args in `calculateInt(-01)` */
	_ = calculateInt(-01)

	/*! suspicious octal args in `calculateInt(012)` */
	_ = calculateInt(calculateInt(012))

	/*! suspicious octal args in `calculateIntPair(01, 2)` */
	_ = calculateIntPair(01, 2)

	/*! suspicious octal args in `calculateIntPair(-1, -012)` */
	_ = calculateIntPair(-1, -012)

	/*! suspicious octal args in `calculateIntPair(01, 02)` */
	_ = calculateIntPair(01, 02)

	/*! suspicious octal args in `calculateInt(01)` */
	/*! suspicious octal args in `calculateInt(02)` */
	_ = calculateIntPair(calculateInt(01), calculateInt(02))

	/*! suspicious octal args in `calculateIntPair(01, calculateInt(02))` */
	/*! suspicious octal args in `calculateInt(02)` */
	_ = calculateIntPair(01, calculateInt(02))

	/*! suspicious octal args in `calculateManyArgs(11, "12", 013)` */
	_ = calculateManyArgs(11, "12", 013)

	/*! suspicious octal args in `calculateManyArgs(-02, "3", -04)` */
	_ = calculateManyArgs(-02, "3", -04)

	/*! suspicious octal args in `math.Exp(012)` */
	_ = math.Exp(012)

	/*! suspicious octal args in `math.Max(12, 01)` */
	_ = math.Max(12, 01)

	/*! suspicious octal args in `math.Max(1, 01)` */
	_ = math.Max(1, math.Max(1, 01))
}

type OpenServer struct {
	x int
}

func (os *OpenServer) Init(x int) {
	os.x = x
}

func warningsOs() {
	var os OpenServer
	/*! suspicious octal args in `os.Init(02)` */
	os.Init(02)
}
