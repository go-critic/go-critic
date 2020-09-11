package checker_test

type myint struct{}

type (
	myint8  = int
	myint16 = int
)

func mybool() {}

var (
	myfloat32 = 1
	myfloat64 = 2
)

const (
	mycomplex64  = 1
	mycomplex128 = 2
)

func (myint) close() {}
