package checker_test

import (
	"bytes"
	"io"
)

func _(w1, w2 *bytes.Buffer) {
	/*! w1.WriteString("foo") should be preferred to the w1.Write([]byte("foo")) */
	w1.Write([]byte("foo"))

	/*! w2.WriteString("bar") should be preferred to the io.WriteString(w2, "bar") */
	io.WriteString(w2, "bar")
}
