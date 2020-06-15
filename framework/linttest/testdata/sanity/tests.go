package sanity_test

import (
	"errors"
	. "errors"
	_ "errors"
	errorspkg1 "errors"
	errorspkg2 "errors"
	"unsafe"
)

var (
	_    = New
	_, _ = errorspkg1.New, errorspkg2.New
	_    = errors.New
	_    = unsafe.Sizeof(int(0))
)

func empty1() {}

func empty2() {
	{
	}
	{
	}
}

func external()

const ()

var ()

type ()

func emptySpecs() {
	const ()
	var ()
	type ()
}

func forRange() {
	var xs []int

	for range xs {
	}
	for _ = range xs {
	}
	for _, _ = range xs {
	}
}

func emptyFor1() {
	for {
	}
}

func emptyFor2() {
	for {
		continue
	}
}

func emptySelect() {
	select {}
}

func emptyStatements() {

	for {
		break
	}

	switch {
	}
	switch {
	case false:
	case true:
	default:
	}
	switch {
	case false:
		fallthrough
	case true:
		break
	default:
	}

	if true {
	}
	if false {
	} else {
	}
	if false {
	} else if false {
	} else {
	}

	goto L0
L0:
}

func initStatements() {
	for _ = 0; ; {
		break
	}
	for _, _ = 0, 0; ; {
		break
	}

	switch _ = 0; {
	}

	if _ = 0; true {
	}

	if variadicArg(1, 2, 3); false {
	}

	switch variadicArg(1, 2, 3); true {
	}

	for variadicArg(1, 2, 3); false; {
	}

	// select can't have init statement.
}

type intAlias = int

type structType struct {
	x int `tag1:"1"`
	y int `tag2:"1,2"`
	z struct{ value int }
}

type typeWithEmbedding struct {
	structType
	intAlias
}

type emptyIface interface{}

type embeddingIface interface {
	emptyIface
	ifaceType
}

type ifaceType interface {
	A() func()
	B() func() func()
	c(func(func()) func()) func(func() func() func()) func()
}

type (
	_ chan [2]int
	_ chan chan int
	_ chan<- int
	_ <-chan int
	_ <-chan <-chan bool
	_ chan []chan []chan<- int
)

type myString string

func convertPtr(x string) *myString {
	return (*myString)(&x)
}

func (myString) noReceiverName1(a, b string) {}

func (*myString) noReceiverName2() (a, b string) { return "", "" }

var noInit1, noInit2 int

func variadicArg(xs ...interface{}) {}

func funcCalls() {
	f0 := func() {}
	f1 := func(x int) {}
	f2 := func(x, y int) {}
	f3 := func(x, y, z int) {}
	fVariadic := func(xs ...int) {}

	f0()
	f1(1)
	f2(1, 2)
	f3(1, 2, 3)
	fVariadic()
	fVariadic(1, 2)
	fVariadic([]int{1, 2}...)
}

func sliceExpressions(xs []int) {
	_ = xs[:]
	_ = xs[0:]
	_ = xs[:0]
	_ = xs[0:0]
	_ = xs[:0:0]
	_ = xs[0:0:0]
}

type myStruct struct {
	field string
}

func (m myStruct) method(a string) {}

func methodExprCall() {
	m := myStruct{}

	myStruct.method(m, "field")
}
