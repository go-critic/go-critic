package lintmain

import (
	"fmt"
	"log"

	"github.com/go-critic/go-critic/framework/cmdutil"
	"github.com/go-critic/go-critic/framework/lintmain/internal/check"
	"github.com/go-critic/go-critic/framework/lintmain/internal/lintdoc"
)

// Config is used to parametrize the linter.
type Config struct {
	Version string
	Name    string
}

var config *Config

// Run executes corresponding main after sub-command resolving.
// Does not return.
func Run(cfg Config) {
	config = &cfg // TODO(quasilyte): don't use global var for this
	log.SetFlags(0)

	// makeExample replaces all ${linter} placeholders to a bound linter name.
	makeExamples := func(examples ...string) []string {
		for i := range examples {
			examples[i] = fmt.Sprintf(examples[i], cfg.Name)
		}
		return examples
	}

	subCommands := []*cmdutil.SubCommand{
		{
			Main:  check.Main,
			Name:  "check",
			Short: "run linter over specified targets",
			Examples: makeExamples(
				"%s check -help",
				"%s check -enable='paramTypeCombine,unslice' strings bytes",
				"%s check -v -enable='#diagnostic' -disable='#experimental,#opinionated' ./...",
			),
		},
		{
			Main:     printVersion,
			Name:     "version",
			Short:    "print linter version",
			Examples: makeExamples("%s version"),
		},
		{
			Main:  lintdoc.Main,
			Name:  "doc",
			Short: "get installed checkers documentation",
			Examples: makeExamples(
				"%s doc -help",
				"%s doc",
				"%s doc checkerName",
			),
		},
	}

	cmdutil.DispatchCommand(subCommands)
}

func printVersion() {
	log.Println(config.Version)
}
