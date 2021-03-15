package checker_test

import (
	"regexp"
)

func domainDotsEscaped() {
	regexp.MustCompile(`google\.com`)

	regexp.CompilePOSIX(`yandex\.ru|radio.yandex\.ru`)
}
