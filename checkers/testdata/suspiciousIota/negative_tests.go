package checker_test

const (
	Good1 = iota
	Good2
	Good3
	Good4
)

const Single = iota

const (
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	SkipFirst = iota
	_
	SkipThird
)

const (
	Red = iota + 10
	Green
	Blue
)

const (
	StatusOK = 200
	StatusError = 500
	StatusBad = 400
)

const (
	MultiA, MultiB = iota, iota + 10
	MultiC, MultiD
)

const (
	FlagNone = 0
	FlagRead = 1 << iota
	FlagWrite
	FlagExecute
)

const (
	StartGood = iota
	EndGood = 100
)

const (
	Pi = 3.14
	Title = "Hello" 
	Count = 42
)

type Color int

const (
	ColorRed Color = iota
	ColorGreen
	ColorBlue
)
