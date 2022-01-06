package checker_test

// Deprecated: part of the old API; use API v2
func ProperDeprecationComment() {}

// This is not a Deprecated: comment at all.
func FalsePositive1() {}

// This is not a Deprecated?
func FalsePositive2() {}

// Deprecated is a function name.
func Deprecated() {}

// deprecated is a type name.
type deprecated struct{}

// Derpecated is a type name.
type Derpecated struct{}

// Note that this one is not deprecated.
func f() {}

var (
	// Dep: ...
	_ = 0

	// deprec
	_ = 0

	// dePreca
	_ = 0
)
