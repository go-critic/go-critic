package checker_test

import (
	"bytes"
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
