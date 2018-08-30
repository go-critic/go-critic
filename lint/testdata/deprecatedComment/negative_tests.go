package linter_test

// Deprecated: part of the old API; use API v2
func ProperDeprecationComment() {}

// This is not a Deprecated: comment at all.
func FalsePositive1() {}

// Deprecated is a function name.
func Deprecated() {}

// deprecated is a type name.
type deprecated struct{}
