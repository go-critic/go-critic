// Package checkers is a gocritic linter main checkers collection.
package checkers

import (
	"github.com/go-critic/go-critic/framework/linter"
	"os"
)

var collection = &linter.CheckerCollection{
	URL: "https://github.com/go-critic/go-critic/checkers",
}

var debug = func() func() bool {
	v := os.Getenv("DEBUG") != ""
	return func() bool {
		return v
	}
}()
