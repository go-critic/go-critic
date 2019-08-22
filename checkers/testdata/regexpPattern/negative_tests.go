package checker_test

import (
	"regexp"
)

func escapedDots() {
	regexp.MustCompile(`google\.com`)

	regexp.CompilePOSIX(`yandex\.ru|radio.yandex\.ru`)
}

func goodAnchors() {
	regexp.Compile(`^100`)
	regexp.Compile(`2\^10`)
	regexp.Compile(`(?im)2\^10`)
	regexp.Compile(`(?m)77^10$`)
	regexp.Compile(`100\$\+1`)
	regexp.Compile(`100$`)
}
