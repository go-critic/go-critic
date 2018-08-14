package flagparser

import (
	"flag"
	"fmt"
	"strings"
)

const EnableAll = "all"
const DisableAll = "all"

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

func (fp *FlagParser) Error() error {
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

func (fp *FlagParser) EnabledCheckers() []string {
	return strings.Split(fp.Enable, ",")
}

func (fp *FlagParser) DisabledCheckers() []string {
	return strings.Split(fp.Disable, ",")
}

func (fp *FlagParser) Parse() {
	flag.Parse()
}

func (fp *FlagParser) ParsedArgs() []string {
	return flag.Args()
}

func (fp *FlagParser) Args() []string {
	return []string{
		"-enable", fp.Enable,
		"-disable", fp.Disable,
		"-config", fp.ConfigFile,
		"-withExperimental=" + fmt.Sprint(fp.WithExperimental),
		"-withOpinionated=" + fmt.Sprint(fp.WithOpinionated),
		"-failcode", fmt.Sprint(fp.FailureExitCode),
		"-checkGenerated=" + fmt.Sprint(fp.CheckGenerated),
		"-shorterErrLocation=" + fmt.Sprint(fp.ShorterErrLocation),
	}
}
