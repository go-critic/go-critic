package checker_test

import (
	"strings"
)

func foo() {
	f, b := "aaa", "bbb"

	/*! suggestion: f == b */
	if strings.Compare(f, b) == 0 {
	}

	/*! suggestion: f == b */
	if 0 == strings.Compare(f, b) {
	}

	/*! suggestion: f > b */
	switch dd := strings.Compare(f, b) > 100; dd {
	case true:
		print(0)
	case false:
		print(1)
	}

	_ = strings.Compare("s", "ww") < 10
	_ = 10 > strings.Compare("s", "ww")

	_ = strings.Compare(f, b) < 10
	_ = 10 > strings.Compare(f, b)
}
