// +build tools

package critic

// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module.

import (
	// Used for CI linting.
	_ "github.com/quasilyte/go-consistent"
	// Used to generate code coverage.
	_ "github.com/mattn/goveralls"
)
