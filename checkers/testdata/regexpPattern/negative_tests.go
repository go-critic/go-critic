package checker_test

import (
	"regexp"
)

func domainDots() {
	regexp.MustCompile(`google\.com`)

	regexp.CompilePOSIX(`yandex\.ru|radio.yandex\.ru`)
}
