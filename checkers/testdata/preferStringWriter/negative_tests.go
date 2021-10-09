package checker_test

import "io"

func _(wr io.Writer) {
	wr.Write([]byte("foo"))
}
