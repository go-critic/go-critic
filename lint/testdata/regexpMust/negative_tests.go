package checker_test

import "regexp"

var anotherRE = regexp.MustCompile(`pat123`)

func noWarnings() {
	_ = regexp.MustCompile(`[0-9]+`)
	_ = regexp.MustCompile(`go-critic linter`)
	_ = regexp.MustCompilePOSIX(`go-critic linter`)
}

func nonConstPatterns(pat string) {
	re, err := regexp.Compile(pat)
	if err != nil {
		panic(err)
	}
	_ = re

	_, _ = regexp.Compile(pat)
	_, _ = regexp.CompilePOSIX(pat)
}
