package checker_tests

import "regexp"

var anotherRE = regexp.MustCompile(`pat123`)

func noWarnings() {
	_ = regexp.MustCompile(`[0-9]+`)
	_ = regexp.MustCompile(`go-critic linter`)
	_ = regexp.MustCompilePOSIX(`go-critic linter`)
}
