package flagparser

import (
	"flag"
	"fmt"
	"strings"
)

// EnableAll represent all checkers value for "enable" option
const EnableAll = "all"

// DisableAll represent all checkers value for "disable" option
const DisableAll = "all"

// NewFlagParser create new FlagParser
func NewFlagParser() *FlagParser {
	fp := &FlagParser{}

	flag.StringVar(&fp.Enable, "enable", EnableAll,
		`comma-separated list of enabled checkers`)
	flag.StringVar(&fp.ConfigFile, "config", "",
		`name of JSON file containing checkers configurations`)
	flag.StringVar(&fp.Disable, "disable", "",
		`comma-separated list of disabled checkers`)
	flag.BoolVar(&fp.WithExperimental, `withExperimental`, false,
		`only for -enable=all, include experimental checks`)
	flag.BoolVar(&fp.WithOpinionated, `withOpinionated`, false,
		`only for -enable=all, include very opinionated checks`)
	flag.IntVar(&fp.FailureExitCode, "failcode", 1,
		`exit code to be used when lint issues are found`)
	flag.BoolVar(&fp.CheckGenerated, "checkGenerated", false,
		`whether to check machine-generated files`)
	flag.BoolVar(&fp.ShorterErrLocation, "shorterErrLocation", true,
		`whether to replace error location prefix with $GOROOT and $GOPATH`)

	return fp
}

// FlagParser help to parse and operate with command flags, share them between commands, etc.
type FlagParser struct {
	Enable             string
	Disable            string
	ConfigFile         string
	WithExperimental   bool
	WithOpinionated    bool
	FailureExitCode    int
	CheckGenerated     bool
	ShorterErrLocation bool
}

// EnabledCheckers return checkers, provided by enable argument
func (fp *FlagParser) EnabledCheckers() []string {
	return strings.Split(fp.Enable, ",")
}

// DisabledCheckers return checkers, provided by disable argument
func (fp *FlagParser) DisabledCheckers() []string {
	return strings.Split(fp.Disable, ",")
}

// Parse and validate command arguments. Return error in case of validation failed.
func (fp *FlagParser) Parse() error {
	flag.Parse()

	if fp.Enable != EnableAll {
		if fp.WithExperimental {
			return fmt.Errorf("-withExperimental used with -enable=%q", fp.Enable)
		}
		if fp.WithOpinionated {
			return fmt.Errorf("-withOpinionated used with -enable=%q", fp.Enable)
		}
	}

	return nil
}

// ParsedArgs return parsed command line arguments
func (fp *FlagParser) ParsedArgs() []string {
	return flag.Args()
}

// Args return slice of arguments, that can be used as arguments for exec command
func (fp *FlagParser) Args() []string {
	return []string{
		"-enable=" + fp.Enable,
		"-disable=" + fp.Disable,
		"-config=" + fp.ConfigFile,
		"-withExperimental=" + fmt.Sprint(fp.WithExperimental),
		"-withOpinionated=" + fmt.Sprint(fp.WithOpinionated),
		"-failcode=" + fmt.Sprint(fp.FailureExitCode),
		"-checkGenerated=" + fmt.Sprint(fp.CheckGenerated),
		"-shorterErrLocation=" + fmt.Sprint(fp.ShorterErrLocation),
	}
}
