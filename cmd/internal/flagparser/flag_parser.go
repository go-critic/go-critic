package flagparser

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// EnableAll represent all checkers value for "enable" option
const EnableAll = "all"

// DisableAll represent all checkers value for "disable" option
const DisableAll = "all"

// NewFlagParser create new FlagParser
func NewFlagParser(flagSet *flag.FlagSet) *FlagParser {
	fp := &FlagParser{}

	fp.flagSet = flagSet

	fp.flagSet.StringVar(&fp.Enable, "enable", EnableAll,
		`comma-separated list of enabled checkers`)
	fp.flagSet.StringVar(&fp.ConfigFile, "config", "",
		`name of JSON file containing checkers configurations`)
	fp.flagSet.StringVar(&fp.Disable, "disable", "",
		`comma-separated list of disabled checkers`)
	fp.flagSet.BoolVar(&fp.WithExperimental, `withExperimental`, false,
		`only for -enable=all, include experimental checks`)
	fp.flagSet.BoolVar(&fp.WithOpinionated, `withOpinionated`, false,
		`only for -enable=all, include very opinionated checks`)
	fp.flagSet.IntVar(&fp.FailureExitCode, "failcode", 1,
		`exit code to be used when lint issues are found`)
	fp.flagSet.BoolVar(&fp.CheckGenerated, "checkGenerated", false,
		`whether to check machine-generated files`)
	fp.flagSet.BoolVar(&fp.ShorterErrLocation, "shorterErrLocation", true,
		`whether to replace error location prefix with $GOROOT and $GOPATH`)

	return fp
}

func newDefaultInvertedFlagParser() *FlagParser {
	fp := NewFlagParser(flag.NewFlagSet("defaultChecker", flag.ContinueOnError))

	fp.flagSet.VisitAll(func(f *flag.Flag) {
		switch f.DefValue {
		case "true":
			f.DefValue = "false"
		case "false":
			f.DefValue = "true"
		default:
			f.DefValue += "_inverted"
		}

		f.Value.Set(f.DefValue)
	})

	return fp
}

// FlagParser help to parse and operate with command flags, share them between commands, etc.
type FlagParser struct {
	flagSet *flag.FlagSet

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
	err := fp.flagSet.Parse(os.Args[1:])

	if err != nil {
		return err
	}

	if fp.Enable != EnableAll {
		if fp.WithExperimental {
			return fmt.Errorf("-withExperimental used with -enable=%q", fp.Enable)
		}
		if fp.WithOpinionated {
			return fmt.Errorf("-withOpinionated used with -enable=%q", fp.Enable)
		}
	}

	if fp.ConfigFile != "" {
		// If config file used, we restrict to use options that control enabledChecker like "enable", "disable", etc.
		// Purpose of this - to avoid complexity with definition which one checkers are enabled where and so on.
		// It will be easier to use, and easier to support.

		// defaultInvertedFlagParser, with default values different from real ones, used in trick for define that
		// arguments was provided by user, or just was set up by default
		// See https://stackoverflow.com/a/51903637/4143494
		defaultInvertedFlagParser := newDefaultInvertedFlagParser()
		err := defaultInvertedFlagParser.flagSet.Parse(os.Args[1:])

		if err != nil {
			return err
		}

		if fp.Enable == defaultInvertedFlagParser.Enable {
			return errors.New("-enable cannot be used with -config option")
		}
		if fp.Disable == defaultInvertedFlagParser.Disable {
			return errors.New("-disable cannot be used with -config option")
		}
		if fp.WithExperimental == defaultInvertedFlagParser.WithExperimental {
			return errors.New("-withExperimental cannot be used with -config option")
		}
		if fp.WithOpinionated == defaultInvertedFlagParser.WithOpinionated {
			return errors.New("-withOpionated cannot be used with -config option")
		}
	}

	return nil
}

// ParsedArgs return parsed command line arguments
func (fp *FlagParser) ParsedArgs() []string {
	return fp.flagSet.Args()
}

// Args return slice of arguments, that can be used as arguments for exec command
func (fp *FlagParser) Args() []string {
	args := []string{
		"-config=" + fp.ConfigFile,
		"-failcode=" + fmt.Sprint(fp.FailureExitCode),
		"-checkGenerated=" + fmt.Sprint(fp.CheckGenerated),
		"-shorterErrLocation=" + fmt.Sprint(fp.ShorterErrLocation),
	}

	if fp.ConfigFile == "" {
		args = append(
			args,
			[]string{
				"-enable=" + fp.Enable,
				"-disable=" + fp.Disable,
				"-withExperimental=" + fmt.Sprint(fp.WithExperimental),
				"-withOpinionated=" + fmt.Sprint(fp.WithOpinionated),
			}...,
		)
	}

	return args
}
