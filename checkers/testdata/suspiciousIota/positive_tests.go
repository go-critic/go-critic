package checker_test

const (
	A = iota
	/*! redundant iota usage for B; iota auto-increments without explicit assignment */
	B = iota
	/*! redundant iota usage for C; iota auto-increments without explicit assignment */
	C = iota
)

const (
	Zero = iota
	/*! redundant iota usage for One; iota auto-increments without explicit assignment */
	One = iota
	/*! redundant iota usage for Two; iota auto-increments without explicit assignment */
	Two = iota
	/*! redundant iota usage for Three; iota auto-increments without explicit assignment */
	Three = iota
)

const (
	First = iota
	/*! mixing explicit values with iota in const block may be confusing for Second */
	Second = 10
	Third
)

const (
	Alpha = iota
	/*! mixing explicit values with iota in const block may be confusing for Beta */
	Beta = 42
	/*! mixing explicit values with iota in const block may be confusing for Gamma */
	Gamma = 100
	Delta
)

const (
	/*! const maxCommandID appears before iota usage; this affects iota values and may cause bugs */
	maxCommandID  = 9999
	ValueGrayFlag = iota + (maxCommandID + 1)
	ValueBlueFlag
	ValueTealFlag
)
