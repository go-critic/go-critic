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
