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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComponentStatus (and ComponentStatus) holds the cluster validation info.
//
// Deprecated: This API is deprecated in v1.19+
type ComponentStatus struct {
	foo string
	// +optional
	bar string

	// +optional
	fooBard []string
}

// ComponentStatusList represents the list of component statuses
//
// Deprecated: This API is deprecated in v1.19+
type ComponentStatusList struct {
	string
	// +optional
	int

	Items []ComponentStatus
}

// SomethingOlder is old
//
// Deprecated: This API is deprecated in v1.19+
//
// and these are information that are in another paragraph, but not part of the deprecation one
type SomethingOlder struct{}

// Deprecated: This API is deprecated in v1.19+
//
// SomethingLegacy is also old
type SomethingLegacy struct{}
