package sanity_test

import (
	. "fmt"
	_ "fmt"
	fmtpkg1 "fmt"
	fmtpkg2 "fmt"
)

import "fmt"
import "unsafe"

var (
	_    = Printf
	_, _ = fmtpkg1.Printf, fmtpkg2.Printf
	_    = fmt.Printf
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

type intAlias = int

type structType struct {
	x int `tag1:"1"`
	y int `tag2:"1,2"`
	z struct{ value int }
}

type typeWithEmbedding struct {
	structType
}

type ifaceType interface {
	A() func()
	B() func() func()
	c(func(func()) func()) func(func() func() func()) func()
}
