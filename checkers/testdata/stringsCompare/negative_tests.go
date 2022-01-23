package checker_test

import (
	"bytes"
	"strings"
)

type s []string

func (s) Compare(x, y string) int {
	if x < y {
		return 1
	}
	return 0
}

func negative() {
	bytes.Compare([]byte{}, []byte{})

	strings := s{}
	_ = strings.Compare("1", "3") == 0
}

func negative2() {
	var a, b = "aaa", "bbb"
	if a > b {
		print(1)
	}

	if a == b {
		print(0)
	}

	if a < b {
		print(-1)
	}

	switch a > b {
	case true:
		print(1)
	case false:
		print(0, -1)
	}
}

func negative3() {
	if "aaa" > "bbb" {
		print(1)
	}

	if "aaa" == "bbb" {
		print(0)
	}

	if "aaa" < "bbb" {
		print(-1)
	}

	switch "aaa" > "bbb" {
	case true:
		print(1)
	case false:
		print(0, -1)
	}
}

func negative4() {
	f, b := "aaa", "bbb"

	_ = strings.Compare(f, b) > -100
	_ = strings.Compare(f, b) < 100
	_ = strings.Compare(f, b) >= -1
	_ = strings.Compare(f, b) <= -1
	_ = strings.Compare(f, b) >= 1
	_ = strings.Compare(f, b) <= 1
	_ = strings.Compare(f, b) >= 0
	_ = strings.Compare(f, b) <= 0
}
