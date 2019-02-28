// +build tools

package critic

import (
	// Used for CI linting.
	_ "github.com/Quasilyte/go-consistent"
	// Used to generate code coverage.
	_ "github.com/mattn/goveralls"
)
