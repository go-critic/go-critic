// Package analyzer implements `go/analysis` compatible interfaces.
package analyzer

import (
	"github.com/go-critic/go-critic/linter"

	"golang.org/x/tools/go/analysis"
)

// Analyzer exports go-critic checkers as analysis-compatible object.
// The set of enabled checkers is controlled via the flags.
// Per-checker params are also passed via the flags.
var Analyzer = &analysis.Analyzer{
	Name: "ruleguard",
	Doc:  "The most opinionated Go source code linter",
	Run:  runAnalyzer,
}

// DisableCache disables initialization optimization.
// This should only be useful for analyzer testing.
var DisableCache = false

var (
	flagGoVersion string
	flagEnable    string
	flagDisable   string
	flagEnableAll bool
	flagDebugInit bool
)

var (
	intParams    = make(map[string]*int)
	boolParams   = make(map[string]*bool)
	stringParams = make(map[string]*string)
)

var registeredCheckers = linter.GetCheckersInfo()

func init() {
	Analyzer.Flags.BoolVar(&flagDebugInit, "debug-init", false,
		`print gocritic initialization related debug info`)
	Analyzer.Flags.BoolVar(&flagEnableAll, "enable-all", false,
		`identical to -enable with all checkers listed. If true, -enable is ignored`)
	Analyzer.Flags.StringVar(&flagEnable, "enable", "#diagnostic,#style,#security",
		`comma-separated list of enabled checkers. Can include #tags`)
	Analyzer.Flags.StringVar(&flagDisable, "disable", "<default>",
		`comma-separated list of checkers to be disabled. Can include #tags`)
	Analyzer.Flags.StringVar(&flagGoVersion, "go", "",
		`select the Go version to target. Leave as string for the latest`)

	for _, info := range registeredCheckers {
		for pname, param := range info.Params {
			key := checkerParamName(info, pname)
			switch v := param.Value.(type) {
			case int:
				intParams[key] = Analyzer.Flags.Int(key, v, param.Usage)
			case bool:
				boolParams[key] = Analyzer.Flags.Bool(key, v, param.Usage)
			case string:
				stringParams[key] = Analyzer.Flags.String(key, v, param.Usage)
			default:
				panic("unreachable") // Checked in AddChecker
			}
		}
	}
}

func checkerParamName(info *linter.CheckerInfo, pname string) string {
	return "@" + info.Name + "." + pname
}
