package checker_test

import "strings"

func _(x, y, z string) {
	/*! suggestion: x + "_" + y */
	_ = strings.Join([]string{x, y}, "_")

	/*! suggestion: y + x */
	_ = strings.Join([]string{y, x}, "")

	/*! suggestion: x + y + z */
	_ = strings.Join([]string{x, y, z}, "")
}
