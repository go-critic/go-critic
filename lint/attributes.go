package lint

// TODO(quasilyte): move attribute-related things here
// when this refactoring is complete.

type checkerAttribute int

// All valid attributes.
var (
	attrExperimental    = new(checkerAttribute)
	attrSyntaxOnly      = new(checkerAttribute)
	attrVeryOpinionated = new(checkerAttribute)
)
