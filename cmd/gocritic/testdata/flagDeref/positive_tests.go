package checker_test

import "flag"

func shouldWarn() {
	/// immediate deref in *flag.Bool("b", false, "") is most likely an error; consider using flag.BoolVar
	_ = *flag.Bool("b", false, "")

	/// immediate deref in *flag.Duration("d", 0, "") is most likely an error; consider using flag.DurationVar
	_ = *flag.Duration("d", 0, "")

	/// immediate deref in *flag.Float64("f64", 0, "") is most likely an error; consider using flag.Float64Var
	_ = *flag.Float64("f64", 0, "")

	/// immediate deref in *flag.Int("i", 0, "") is most likely an error; consider using flag.IntVar
	_ = *flag.Int("i", 0, "")

	/// immediate deref in *flag.Int64("i64", 0, "") is most likely an error; consider using flag.Int64Var
	_ = *flag.Int64("i64", 0, "")

	/// immediate deref in *flag.String("s", "", "") is most likely an error; consider using flag.StringVar
	_ = *flag.String("s", "", "")

	/// immediate deref in *flag.Uint("u", 0, "") is most likely an error; consider using flag.UintVar
	_ = *flag.Uint("u", 0, "")

	/// immediate deref in *flag.Uint64("u64", 0, "") is most likely an error; consider using flag.Uint64Var
	_ = *flag.Uint64("u64", 0, "")
}
