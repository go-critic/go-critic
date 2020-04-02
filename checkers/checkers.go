// Package checkers is a gocritic linter main checkers collection.
package checkers

import (
	"os"

	"github.com/go-lintpack/lintpack"
)

var collection = &lintpack.CheckerCollection{
	URL: "https://github.com/go-critic/go-critic/checkers",
}

var debug = func() func() bool {
	v := os.Getenv("DEBUG") != ""
	return func() bool {
		return v
	}
}()
