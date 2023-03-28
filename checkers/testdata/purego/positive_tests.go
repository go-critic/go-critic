/*! Importing `unsafe` with `go:build purego` is not allowed */
//go:build purego

package main

import (
	"unsafe"
)

var _ unsafe.Pointer
