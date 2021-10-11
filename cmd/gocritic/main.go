package main

import (
	_ "github.com/go-critic/go-critic/checkers" // Register go-critic checkers
	"github.com/go-critic/go-critic/framework/lintmain"
)

var Version string

func main() {
	lintmain.Run(lintmain.Config{
		Name:    "gocritic",
		Version: Version,
	})
}
