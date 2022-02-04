package lintmain

import (
	"log"

	"github.com/go-critic/go-critic/framework/lintmain/internal/check"
	"github.com/go-critic/go-critic/framework/lintmain/internal/lintdoc"

	"github.com/cristalhq/acmd"
)

// Config is used to parametrize the linter.
type Config struct {
	Version string
	Name    string
}

// Run executes corresponding main after sub-command resolving.
// Does not return.
func Run(cfg Config) {
	log.SetFlags(0)

	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName: cfg.Name,
		Version: cfg.Version,
	})
	if err := r.Run(); err != nil {
		log.Print(err.Error())
	}
}

var cmds = []acmd.Command{
	{
		Name:        "check",
		Description: "run linter over specified targets",
		Do:          check.Main,
	},
	{
		Name:        "doc",
		Description: "get installed checkers documentation",
		Do:          lintdoc.Main,
	},
	{
		Name:        "__complete",
		Description: "get installed checkers documentation",
		Do:          acmd.AutocompleteFor,
	},
}
