package checker_test

type kek struct{}

// Foo ...
/// silencing go lint doc-comment warnings is unadvised
func Foo() {}

// Wow
/// silencing go lint doc-comment warnings is unadvised
func (k *kek) Wow() {}

//Well
/// silencing go lint doc-comment warnings is unadvised
func Well() {}

// Goo go go.
/// silencing go lint doc-comment warnings is unadvised
func Go() {}
