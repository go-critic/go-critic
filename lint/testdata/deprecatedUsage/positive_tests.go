package checker_test

// Foo disables generics in Go.
//
// Deprecated: Use Bar from pkg/baz instead.
func Foo() {
}

func baz() {
	/// Foo() is deprected, see doc: Use Bar from pkg/baz instead.
	Foo()
}
