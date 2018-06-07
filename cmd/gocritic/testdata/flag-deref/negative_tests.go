package checker_test

import "flag"

func noDerefUsage() {
	_ = flag.Bool("b", false, "")
	_ = flag.Duration("d", 0, "")
	_ = flag.Float64("f64", 0, "")
	_ = flag.Int("i", 0, "")
	_ = flag.Int64("i64", 0, "")
	_ = flag.String("s", "", "")
	_ = flag.Uint("u", 0, "")
	_ = flag.Uint64("u64", 0, "")
}

func varFunctionsUsage() {
	// This test avoids time.Duration value to exclude "time"
	// from imports list. (Makes test faster.)

	var b bool
	flag.BoolVar(&b, "b", false, "")
	var f64 float64
	flag.Float64Var(&f64, "f64", 0, "")
	var i int
	flag.IntVar(&i, "i", 0, "")
	var i64 int64
	flag.Int64Var(&i64, "i64", 0, "")
	var s string
	flag.StringVar(&s, "s", "", "")
	var u uint
	flag.UintVar(&u, "u", 0, "")
	var u64 uint64
	flag.Uint64Var(&u64, "u64", 0, "")
}
