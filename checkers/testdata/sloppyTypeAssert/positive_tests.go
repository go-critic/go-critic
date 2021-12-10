package checker_test

import (
	"io"
)

type underlyingReader io.Reader

func redundantTypeAsserts(eface interface{}, r io.Reader, rc io.ReadCloser) {
	/*! type assertion from/to types are identical */
	_ = rc.(io.ReadCloser)
}
