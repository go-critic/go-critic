package checker_test

import (
	"testing"
)

/*! func(a int, b int, c int) could be replaced with func(a, b, c int) */
func foo(a int, t *testing.T) {}
