// +build tools

// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module.

package critic

import (
	// Used for CI linting.
	_ "github.com/Quasilyte/go-consistent"
	// Used to generate code coverage.
	_ "github.com/mattn/goveralls"
)
