package checker_test

import (
	"io/fs"
	"log"
	"math"
	"os"
)

func calculateInt(x int) int {
	return x
}

func calculateHex(x int) int {
	return x
}

func calculateFloat(x float64) float64 {
	return x
}

func calculateString(x string) string {
	return x
}

func NoWarningsCalc() {
	_ = calculateInt(0)
	_ = calculateInt(1)
	_ = calculateInt(+1)
	_ = calculateInt(-1)
	_ = calculateInt(12)
	_ = calculateInt(1 + 2)

	var x = 0x3
	_ = calculateInt(x)

	_ = calculateHex(0x0)
	_ = calculateHex(0X42)
	_ = calculateHex(0xAA1)
	_ = calculateHex(-0xaa1)

	_ = calculateFloat(0.2)
	_ = calculateFloat(+0.2)
	_ = calculateFloat(-0.2)

	_ = calculateString("1")
	_ = calculateString("01")
	_ = calculateString("0.1")

	_, _ = calculateIntPair(1, 2)
	_, _ = calculateIntPair(-1, 2)
	_, _ = calculateIntPair(0, 2)

	_ = 0b00
	_ = 0b1
	_ = 0b_0000_0101

	_ = math.Exp(12)
	_ = math.Exp(0x12)
	_ = math.Max(12, 0xd)
	_ = math.Min(0, 1)
}

func NoWarningsFs() {
	_ = fs.FileMode(0o555)
}

func NoWarningsOsOpenFile() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0o755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func NoWarningsOsWriteFile() {
	_ = os.WriteFile("notes.txt", nil, 0o666)
}
