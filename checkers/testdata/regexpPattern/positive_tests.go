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
