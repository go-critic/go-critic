package linter_test

/// use `Deprecated: ` (note the casing) instead of `deprecated: `
// deprecated: part of the old API; use API v2
func LowerCasePrefix() {}

/// use `Deprecated: ` (note the casing) instead of `DEPRECATED: `
// DEPRECATED: part of the old API; use API v2
func UpperCasePrefix() {}

/// use `:` instead of `,` in `Deprecated, `
// Deprecated, use XYZ instead.
func CommaInsteadOfColon() {}

/// the proper format is `Deprecated: <text>`
// BadFormat1 is an example.
// This function is deprecated, use XYZ instead.
func BadFormat1() {}

/// the proper format is `Deprecated: <text>`
// BadFormat2 is an example, too.
//
// this function is deprecated, use XYZ instead.
func BadFormat2() {}

/// the proper format is `Deprecated: <text>`
// BadFormat3 is an example, too.
//
// This type is deprecated, use XYZ instead.
type BadFormat3 int

/// the proper format is `Deprecated: <text>`
// this type is deprecated, use XYZ instead.
type badFormat4 int

/// the proper format is `Deprecated: <text>`
// deprecated! use something-else/a.f() instead
const BadFormat5 int = 10

/// the proper format is `Deprecated: <text>`
//
//
// deprecated use XYZ instead
const BadFormat6 int = 10

/// the proper format is `Deprecated: <text>`
//
// DEPRECATED. use XYZ instead
const BadFormat7 int = 10

/// the proper format is `Deprecated: <text>`
//
// (This is why we're using case-insensitive patterns.)
//
// Deprecated! USE ANYTHING INSTEAD!
const BadFormat8 = 10

type badNestedDoc struct {
	/// use `Deprecated: ` (note the casing) instead of `deprecated: `
	// deprecated: ha-ha
	foo struct {
		/// use `:` instead of `,` in `Deprecated, `
		// Deprecated, first deprecated field
		field int

		/// use `Deprecated: ` (note the casing) instead of `deprecated: `
		// deprecated: another one
		bar struct {
			/// use `Deprecated: ` (note the casing) instead of `deprecated: `
			// deprecated: deprecated field
			field int
		}
	}
}
