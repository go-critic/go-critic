package checker_test

import (
	"flag"
)

func dynamicFlagName(flagName string) {
	_ = flag.String(flagName, "", "") // can't analyze
}

func methodCall() {
	// See #784.
	var getter flag.Getter
	_ = getter.String()
}

func noWhitespace() {
	_ = flag.Bool("name", false, "")
	_ = flag.Duration("name", 0, "")
	_ = flag.Float64("name", 0, "")
	_ = flag.String("name", "", "")
	_ = flag.Int("name", 0, "")
	_ = flag.Int64("name", 0, "")
	_ = flag.Uint("name", 0, "")
	_ = flag.Uint64("name", 0, "")

	flag.BoolVar(nil, "name", false, "")
	flag.DurationVar(nil, "name", 0, "")
	flag.Float64Var(nil, "name", 0, "")
	flag.StringVar(nil, "name", "", "")
	flag.IntVar(nil, "name", 0, "")
	flag.Int64Var(nil, "name", 0, "")
	flag.UintVar(nil, "name", 0, "")
	flag.Uint64Var(nil, "name", 0, "")
}

func namesWithNumbers() {
	_ = flag.Bool("name13", false, "")
	_ = flag.Duration("na34me", 0, "")
	_ = flag.Float64("nam2e", 0, "")
	_ = flag.String("nam0e", "", "")

	flag.IntVar(nil, "1name", 0, "")
	flag.Int64Var(nil, "name1", 0, "")
	flag.UintVar(nil, "name0", 0, "")
	flag.Uint64Var(nil, "name1000", 0, "")
}

func namesWithDashes() {
	_ = flag.Bool("name-with-a-dash", false, "")
	_ = flag.Duration("name-with-a-dash", 0, "")
	_ = flag.Float64("name-with-a-dash", 0, "")
	_ = flag.String("name-with-a-dash", "", "")
	_ = flag.Int("name-with-a-dash", 0, "")
	_ = flag.Int64("name-with-a-dash", 0, "")
	_ = flag.Uint("name-with-a-dash", 0, "")
	_ = flag.Uint64("name-with-a-dash", 0, "")

	flag.BoolVar(nil, "name-with-a-dash", false, "")
	flag.DurationVar(nil, "name-with-a-dash", 0, "")
	flag.Float64Var(nil, "name-with-a-dash", 0, "")
	flag.StringVar(nil, "name-with-a-dash", "", "")
	flag.IntVar(nil, "name-with-a-dash", 0, "")
	flag.Int64Var(nil, "name-with-a-dash", 0, "")
	flag.UintVar(nil, "name-with-a-dash", 0, "")
	flag.Uint64Var(nil, "name-with-a-dash", 0, "")
}
