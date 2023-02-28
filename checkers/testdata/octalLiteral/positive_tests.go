package checker_test

import (
	"io/fs"
	"log"
	"math"
	"os"
)

func calculateIntPair(x, y int) (int, int) {
	return x, y
}

func calculateManyArgs(x int, s string, y int) (int, string, int) {
	return x, s, y
}

func warningsCalc() {
	/*! use new octal literal style, 0o3 */
	var x = 03
	_ = calculateInt(x)

	/*! use new octal literal style, 0o0 */
	_ = calculateInt(00)

	/*! use new octal literal style, 0o1 */
	_ = calculateInt(+01)

	/*! use new octal literal style, 0o1 */
	_ = calculateInt(-01)

	/*! use new octal literal style, 0o12 */
	_ = calculateInt(calculateInt(012))

	/*! use new octal literal style, 0o1 */
	_, _ = calculateIntPair(01, 2)

	/*! use new octal literal style, 0o12 */
	_, _ = calculateIntPair(-1, -012)

	/*! use new octal literal style, 0o1 */
	/*! use new octal literal style, 0o2 */
	_, _ = calculateIntPair(01, 02)

	/*! use new octal literal style, 0o1 */
	/*! use new octal literal style, 0o2 */
	_, _ = calculateIntPair(calculateInt(01), calculateInt(02))

	/*! use new octal literal style, 0o1 */
	/*! use new octal literal style, 0o2 */
	_, _ = calculateIntPair(01, calculateInt(02))

	/*! use new octal literal style, 0o13 */
	_, _, _ = calculateManyArgs(11, "12", 013)

	/*! use new octal literal style, 0o2 */
	/*! use new octal literal style, 0o4 */
	_, _, _ = calculateManyArgs(-02, "3", -04)

	/*! use new octal literal style, 0o12 */
	_ = math.Exp(012)

	/*! use new octal literal style, 0o1 */
	_ = math.Max(12, 01)

	/*! use new octal literal style, 0o1 */
	_ = math.Max(1, math.Max(1, 01))
}

type OpenServer struct {
	x int
}

func (os *OpenServer) Init(x int) {
	os.x = x
}

func warningsFs() {
	/*! use new octal literal style, 0o555 */
	_ = fs.FileMode(0555)
}

func warningsOsOpenFile() {
	/*! use new octal literal style, 0o755 */
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func warningsOsWriteFile() {
	/*! use new octal literal style, 0o666 */
	_ = os.WriteFile("notes.txt", nil, 0666)
}
