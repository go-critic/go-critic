package checker_test

import (
	"io"
)

type underlyingReader io.Reader

func redundantTypeAsserts(eface interface{}, r io.Reader, rc io.ReadCloser) {
	/*! type assertion to interface{} may be redundant */
	_ = r.(interface{})

	/*! type assertion may be redundant as rc always implements selected interface */
	_ = rc.(io.Reader)

	/*! type assertion from/to types are identical */
	_ = rc.(io.ReadCloser)

	var ur underlyingReader

	/*! type assertion may be redundant as ur always implements selected interface */
	_ = ur.(io.Reader)

	/*! type assertion may be redundant as r always implements selected interface */
	_ = r.(underlyingReader)
}
