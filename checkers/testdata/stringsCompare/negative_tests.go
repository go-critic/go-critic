package checker_test

import (
	"bytes"
)

type s []string

func (s) Compare(x, y string) bool { return x > y }

func negative() {
	bytes.Compare([]byte{}, []byte{})

	strings := s{}
	print(strings.Compare("1", "3"))
}
