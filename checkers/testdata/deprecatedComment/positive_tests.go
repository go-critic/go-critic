package checker_test

/*! use `Deprecated: ` (note the casing) instead of `deprecated: ` */
// deprecated: part of the old API; use API v2
func LowerCasePrefix() {}

/*! use `Deprecated: ` (note the casing) instead of `DEPRECATED: ` */
// DEPRECATED: part of the old API; use API v2
func UpperCasePrefix() {}

/*! use `:` instead of `,` in `Deprecated, ` */
// Deprecated, use XYZ instead.
func CommaInsteadOfColon() {}

// BadFormat1 is an example.
/*! the proper format is `Deprecated: <text>` */
// This function is deprecated, use XYZ instead.
func BadFormat1() {}

// BadFormat2 is an example, too.
//
/*! the proper format is `Deprecated: <text>` */
// this function is deprecated, use XYZ instead.
func BadFormat2() {}

// BadFormat3 is an example, too.
//
/*! the proper format is `Deprecated: <text>` */
// This type is deprecated, use XYZ instead.
type BadFormat3 int

/*! the proper format is `Deprecated: <text>` */
// this type is deprecated, use XYZ instead.
type badFormat4 int

/*! the proper format is `Deprecated: <text>` */
// deprecated! use something-else/a.f() instead
const BadFormat5 int = 10

//
//
/*! the proper format is `Deprecated: <text>` */
// deprecated use XYZ instead
const BadFormat6 int = 10

//
/*! the proper format is `Deprecated: <text>` */
// DEPRECATED. use XYZ instead
const BadFormat7 int = 10

//
// (This is why we're using case-insensitive patterns.)
//
/*! the proper format is `Deprecated: <text>` */
// Deprecated! USE ANYTHING INSTEAD!
const BadFormat8 = 10

//
// (This is why we're using case-insensitive patterns.)
//
/*! the proper format is `Deprecated: <text>` */
// [[deprecated]]
const BadFormat9 = 10

// BadFormat10 is an example, too.
//
/*! the proper format is `Deprecated: <text>` */
// ThIs TyPe iS DepRecateD, use foo instead.
type BadFormat10 int

/*! use `Deprecated: ` (note the casing) instead of `DEPRECATED: ` */
// DEPRECATED: part of the old API; use API v2
func BadFormat11() {}

type badNestedDoc struct {
	/*! use `Deprecated: ` (note the casing) instead of `deprecated: ` */
	// deprecated: ha-ha
	foo struct {
		/*! use `:` instead of `,` in `Deprecated, ` */
		// Deprecated, first deprecated field
		field int

		/*! use `Deprecated: ` (note the casing) instead of `deprecated: ` */
		// deprecated: another one
		bar struct {
			/*! use `Deprecated: ` (note the casing) instead of `deprecated: ` */
			// deprecated: deprecated field
			field int
		}
	}
}

/*! typo in `Dprecated`; should be `Deprecated` */
// Dprecated: ...
func withTypo1() {}

var (
	/*! typo in `Dprecated`; should be `Deprecated` */
	// Dprecated: ...
	_ = 0

	/*! typo in `Derecated`; should be `Deprecated` */
	// Derecated: ...
	_ = 0

	/*! typo in `Depecated`; should be `Deprecated` */
	// Depecated: ...
	_ = 0

	/*! typo in `Deprcated`; should be `Deprecated` */
	// Deprcated: ...
	_ = 0

	/*! typo in `Depreated`; should be `Deprecated` */
	// Depreated: ...
	_ = 0

	/*! typo in `Deprected`; should be `Deprecated` */
	// Deprected: ...
	_ = 0

	/*! typo in `Deprecaed`; should be `Deprecated` */
	// Deprecaed: ...
	_ = 0

	/*! typo in `Deprecatd`; should be `Deprecated` */
	// Deprecatd: ...
	_ = 0

	/*! typo in `Deprecate`; should be `Deprecated` */
	// Deprecate: ...
	_ = 0

	/*! typo in `DeprEcate`; should be `Deprecated` */
	// DeprEcate: ...
	_ = 0

	/*! typo in `deprecate`; should be `Deprecated` */
	// deprecate: ...
	_ = 0

	/*! typo in `dePrecate`; should be `Deprecated` */
	// dePrecate: ...
	_ = 0

	/*! typo in `Depekated`; should be `Deprecated` */
	// Depekated: ...
	_ = 0

	/*! typo in `DepeKated`; should be `Deprecated` */
	// DepeKated: ...
	_ = 0

	/*! typo in `Derpecate`; should be `Deprecated` */
	// Derpecate: ...
	_ = 0

	/*! typo in `DERPecate`; should be `Deprecated` */
	// DERPecate: ...
	_ = 0

	/*! typo in `Depreacted`; should be `Deprecated` */
	// Depreacted: ...
	_ = 0
)

/*! the proper format is `Deprecated: <text>` */
// NOTE: Deprecated. Use bar instead.
func foo1() {
}

/*! the proper format is `Deprecated: <text>` */
// NOTE: Deprecated.
func foo2() {
}

/*! the proper format is `Deprecated: <text>` */
// deprecated in 1.8: use bar instead.
type foo3 string

/*! the proper format is `Deprecated: <text>` */
// deprecated in 1.11: use f instead.
type foo4 string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComponentStatusBad (and ComponentStatusList) holds the cluster validation info.
/*! use `Deprecated: ` (note the casing) instead of `DeprecaTEd: ` */
// DeprecaTEd: This API is deprecated in v1.19+
type ComponentStatusBad struct {
	foo string
	// +optional
	bar string

	// +optional
	fooBard []string
}

// ComponentStatusList represents the list of component statuses
/*! use `Deprecated: ` (note the casing) instead of `DeprecaTed: ` */
// DeprecaTed: This API is deprecated in v1.19+
type ComponentStatusListBad struct {
	string
	// +optional
	int

	Items []int
}
