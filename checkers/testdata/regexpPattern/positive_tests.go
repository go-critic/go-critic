package checker_test

import (
	"regexp"
)

func domainDots() {
	/*! '.com' should probably be '\.com' */
	regexp.MustCompile(`google.com`)

	/*! '.ru' should probably be '\.ru' */
	regexp.CompilePOSIX(`yandex.ru|radio.yandex.ru`)
}

func unescapedAnchors() {
	/*! unescaped ^ in the middle of the regexp */
	regexp.Compile(`2^10`)

	/*! unescaped $ in the middle of the regexp */
	regexp.Compile(`100$\+1`)
}
