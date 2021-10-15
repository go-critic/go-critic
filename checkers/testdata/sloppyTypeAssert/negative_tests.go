package checker_test

import (
	"io"
)

func noWarnings(eface interface{}, r io.Reader, rc io.ReadCloser) {
	// interface{} -> other non-empty interface assertion.
	_ = eface.(io.Reader)

	// assertion to a wider interface.
	_ = r.(io.ReadCloser)

	_ = r.(interface{})

	_ = rc.(io.Reader)

	var ur underlyingReader

	_ = ur.(io.Reader)

	_ = r.(underlyingReader)
}
