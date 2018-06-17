package lint

type attributeSet struct {
	// SyntaxOnly marks rules that can be checker by using AST only.
	// Such rule implementations should not use types info.
	//
	// It's OK not to mark rule as SyntaxOnly even if current
	// implementation does not use types info.
	SyntaxOnly bool

	// Experimental marks rule implementation as experimental.
	// Reasons to mark anything as experimental:
	//	- rule name can change in near time
	//	- checker gives more false positives than anticipated
	//	- rule requires more testing or alternative design
	Experimental bool

	// VeryOpinionated marks rule as controversial for some audience and
	// that it might be not suitable for everyone.
	VeryOpinionated bool
}

type checkerAttribute int

const (
	attrExperimental checkerAttribute = iota
	attrSyntaxOnly
	attrVeryOpinionated
)
