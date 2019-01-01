package checker_test

// Z .
/*! silencing go lint doc-comment warnings is unadvised */
func Z() {
}

type (
	// An XX .
	/*! silencing go lint doc-comment warnings is unadvised */
	XX struct{}
)

// An YY whatever.
/*! silencing go lint doc-comment warnings is unadvised */
type YY struct{}

// The ZZ whatever
/*! silencing go lint doc-comment warnings is unadvised */
type ZZ struct{}

// Foo ...
/*! silencing go lint doc-comment warnings is unadvised */
func Foo() {}

// Baz XXX.
/*! silencing go lint doc-comment warnings is unadvised */
func Baz() {}

// Barr XXX
/*! silencing go lint doc-comment warnings is unadvised */
func Barr() {}
