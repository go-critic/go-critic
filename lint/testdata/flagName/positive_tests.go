package checker_test

import (
	"flag"
)

func flagsWithWhitespace() {
	/// flag name " name" contains whitespace
	_ = flag.Bool(" name", false, "")
	/// flag name "\tname\t" contains whitespace
	_ = flag.Duration("\tname\t", 0, "")
	/// flag name "  name " contains whitespace
	_ = flag.Float64("  name ", 0, "")
	/// flag name "name " contains whitespace
	_ = flag.String("name ", "", "")
	/// flag name "name\t\t\t" contains whitespace
	_ = flag.Int("name\t\t\t", 0, "")
	/// flag name "name\t\t" contains whitespace
	_ = flag.Int64("name\t\t", 0, "")
	/// flag name "name\t" contains whitespace
	_ = flag.Uint("name\t", 0, "")
	/// flag name " \tname \t" contains whitespace
	_ = flag.Uint64(" \tname \t", 0, "")

	/// flag name " name" contains whitespace
	flag.BoolVar(nil, " name", false, "")
	/// flag name "\tname\t" contains whitespace
	flag.DurationVar(nil, "\tname\t", 0, "")
	/// flag name "  name " contains whitespace
	flag.Float64Var(nil, "  name ", 0, "")
	/// flag name "name " contains whitespace
	flag.StringVar(nil, "name ", "", "")
	/// flag name "name\t\t\t" contains whitespace
	flag.IntVar(nil, "name\t\t\t", 0, "")
	/// flag name "name\t\t" contains whitespace
	flag.Int64Var(nil, "name\t\t", 0, "")
	/// flag name "name\t" contains whitespace
	flag.UintVar(nil, "name\t", 0, "")
	/// flag name " \tname \t" contains whitespace
	flag.Uint64Var(nil, " \tname \t", 0, "")
}
