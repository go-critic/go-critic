package checker_test

import (
	"flag"
)

func flagsWithWhitespace() {
	/*! flag name " name" contains whitespace */
	_ = flag.Bool(" name", false, "")
	/*! flag name "  name  " contains whitespace */
	_ = flag.Duration("  name  ", 0, "")
	/*! flag name "  name " contains whitespace */
	_ = flag.Float64("  name ", 0, "")
	/*! flag name "name " contains whitespace */
	_ = flag.String("name ", "", "")
	/*! flag name "name      " contains whitespace */
	_ = flag.Int("name      ", 0, "")
	/*! flag name "name    " contains whitespace */
	_ = flag.Int64("name    ", 0, "")
	/*! flag name "name  " contains whitespace */
	_ = flag.Uint("name  ", 0, "")
	/*! flag name "   name   " contains whitespace */
	_ = flag.Uint64("   name   ", 0, "")

	/*! flag name " name" contains whitespace */
	flag.BoolVar(nil, " name", false, "")
	/*! flag name "  name  " contains whitespace */
	flag.DurationVar(nil, "  name  ", 0, "")
	/*! flag name "  name " contains whitespace */
	flag.Float64Var(nil, "  name ", 0, "")
	/*! flag name "name " contains whitespace */
	flag.StringVar(nil, "name ", "", "")
	/*! flag name "name      " contains whitespace */
	flag.IntVar(nil, "name      ", 0, "")
	/*! flag name "name    " contains whitespace */
	flag.Int64Var(nil, "name    ", 0, "")
	/*! flag name "name  " contains whitespace */
	flag.UintVar(nil, "name  ", 0, "")
	/*! flag name "   name   " contains whitespace */
	flag.Uint64Var(nil, "   name   ", 0, "")
}
