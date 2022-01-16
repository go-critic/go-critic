package checker_test

import (
	"strings"
)

type aliasInt struct {
	int
}

func foo() {
	f, b := "aaa", "bbb"

	/*! don't use strings.Compare on hot path, change it to built-in operators */
	if strings.Compare(f, b) == 0 {
	}

	/*! don't use strings.Compare on hot path, change it to built-in operators */
	switch strings.Compare(f, b) {
	case 0:
		print(0)
	case 1:
		print(1)
	case -1:
		print(-1)
	}

	/*! don't use strings.Compare on hot path, change it to built-in operators */
	switch dd := strings.Compare(f, b); dd {
	case 0:
		print(0)
	case 1:
		print(1)
	case -1:
		print(-1)
	}
}

func kk(b ...int) { print(b) }

func bar() {
	/*! don't use strings.Compare on hot path, change it to built-in operators */
	kk := strings.Compare("s", "ww")
	print(kk)
}
