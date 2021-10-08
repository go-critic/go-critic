package checker_test

import (
	"io"
)

type f struct{}

func (f) Sprintf() string {
	return "abc"
}

func _(w io.Writer) {
	var fmt f
	w.Write([]byte(fmt.Sprintf()))
}
