package main

import (
	"log"

	"github.com/go-critic/go-critic/checkers"

	"github.com/cristalhq/acmd"
)

var Version = "v0.0.0-SNAPSHOT"

func main() {
	err := checkers.InitEmbeddedRules()
	if err != nil {
		panic(err)
	}

	run(config{
		Name:    "gocritic",
		Version: Version,
	})
}

// config is used to parametrize the linter.
type config struct {
	Version string
	Name    string
}

// Run executes corresponding main after sub-command resolving.
// Does not return.
func run(cfg config) {
	log.SetFlags(0)

	cmds := []acmd.Command{
		{
			Name:        "check",
			Description: "run linter over specified targets",
			ExecFunc:    runCheck,
		},
		{
			Name:        "doc",
			Description: "get installed checkers documentation",
			ExecFunc:    runDocs,
		},
	}

	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName: cfg.Name,
		Version: cfg.Version,
	})
	if err := r.Run(); err != nil {
		log.Print(err.Error())
	}
}
