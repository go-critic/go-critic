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

	r := acmd.RunnerOf(nil, acmd.Config{
		AppName: cfg.Name,
		Version: cfg.Version,
	})
	_ = r.Run()
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
}
