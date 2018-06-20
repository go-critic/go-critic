package sanity_test

import (
	. "errors"
	_ "errors"
	errorspkg1 "errors"
	errorspkg2 "errors"
)

import "errors"
import "unsafe"

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
}

type emptyIface interface{}

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
