package lint

// TODO(quasilyte): move attribute-related things here
// when this refactoring is complete.

type checkerAttribute int

const (
	attrExperimental checkerAttribute = iota
	attrSyntaxOnly
	attrVeryOpinionated
)
