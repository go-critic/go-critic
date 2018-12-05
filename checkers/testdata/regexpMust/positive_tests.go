package checker_test

import "regexp"

/*! for const patterns like `pat123`, use regexp.MustCompile */
var myRegexp, _ = regexp.Compile(`pat123`)

func warnings() {
	/*! for const patterns like `[0-9]+`, use regexp.MustCompile */
	re, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		panic(err)
	}
	_ = re

	/*! for const patterns like (`go-critic linter`), use regexp.MustCompile */
	_, _ = regexp.Compile((`go-critic linter`))
	/*! for const patterns like `go-critic linter`, use regexp.MustCompilePOSIX */
	_, _ = regexp.CompilePOSIX(`go-critic linter`)
}
