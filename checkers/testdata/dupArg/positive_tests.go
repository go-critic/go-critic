package checker_test

import (
	"bytes"
	"go/types"
	"image"
	"image/draw"
	"math"
	"reflect"
	"strings"
)

type MyInt struct{ value int }

func (x MyInt) Equal(other MyInt) bool  { return x.value == other.value }
func (x MyInt) Equals(other MyInt) bool { return x.value == other.value }
func (x MyInt) Compare(other MyInt) int { return 0 }
func (x MyInt) Cmp(other MyInt) int     { return 0 }

func duplicatedReceiver() {
	{
		var x MyInt
		/*! suspicious method call with the same argument and receiver */
		_ = x.Equal(x)
	}
	{
		var x MyInt
		/*! suspicious method call with the same argument and receiver */
		_ = x.Equals(x)
	}
	{
		var x MyInt
		/*! suspicious method call with the same argument and receiver */
		_ = x.Compare(x)
	}
	{
		var x MyInt
		/*! suspicious method call with the same argument and receiver */
		_ = x.Cmp(x)
	}
}

func duplicatedArgs() {
	var dstSlice []int
	var s string
	var b []byte
	var dstRV reflect.Value
	var typ types.Type
	var dstImg draw.Image

	/*! suspicious duplicated args in copy(dstSlice, dstSlice) */
	copy(dstSlice, dstSlice)

	/*! suspicious duplicated args in reflect.Copy(dstRV, dstRV) */
	_ = reflect.Copy(dstRV, dstRV)
	/*! suspicious duplicated args in reflect.DeepEqual(s, s) */
	_ = reflect.DeepEqual(s, s)

	/*! suspicious duplicated args in strings.Contains(s, s) */
	_ = strings.Contains(s, s)
	/*! suspicious duplicated args in strings.Compare(s, s) */
	_ = strings.Compare(s, s)
	/*! suspicious duplicated args in strings.EqualFold(s, s) */
	_ = strings.EqualFold(s, s)
	/*! suspicious duplicated args in strings.HasPrefix(s, s) */
	_ = strings.HasPrefix(s, s)
	/*! suspicious duplicated args in strings.HasSuffix(s, s) */
	_ = strings.HasSuffix(s, s)
	/*! suspicious duplicated args in strings.Index(s, s) */
	_ = strings.Index(s, s)
	/*! suspicious duplicated args in strings.LastIndex(s, s) */
	_ = strings.LastIndex(s, s)
	/*! suspicious duplicated args in strings.Split(s, s) */
	_ = strings.Split(s, s)
	/*! suspicious duplicated args in strings.SplitAfter(s, s) */
	_ = strings.SplitAfter(s, s)
	/*! suspicious duplicated args in strings.SplitAfterN(s, s, 2) */
	_ = strings.SplitAfterN(s, s, 2)
	/*! suspicious duplicated args in strings.SplitN(s, s, 2) */
	_ = strings.SplitN(s, s, 2)
	/*! suspicious duplicated args in strings.Replace("", s, s, 1) */
	_ = strings.Replace("", s, s, 1)
	/*! suspicious duplicated args in strings.Replace("", "a", "a", 1) */
	_ = strings.Replace("", "a", "a", 1)

	/*! suspicious duplicated args in bytes.Contains(b, b) */
	_ = bytes.Contains(b, b)
	/*! suspicious duplicated args in bytes.Compare(b, b) */
	_ = bytes.Compare(b, b)
	/*! suspicious duplicated args in bytes.Equal(b, b) */
	_ = bytes.Equal(b, b)
	/*! suspicious duplicated args in bytes.EqualFold(b, b) */
	_ = bytes.EqualFold(b, b)
	/*! suspicious duplicated args in bytes.HasPrefix(b, b) */
	_ = bytes.HasPrefix(b, b)
	/*! suspicious duplicated args in bytes.HasSuffix(b, b) */
	_ = bytes.HasSuffix(b, b)
	/*! suspicious duplicated args in bytes.LastIndex(b, b) */
	_ = bytes.LastIndex(b, b)
	/*! suspicious duplicated args in bytes.Split(b, b) */
	_ = bytes.Split(b, b)
	/*! suspicious duplicated args in bytes.SplitAfter(b, b) */
	_ = bytes.SplitAfter(b, b)
	/*! suspicious duplicated args in bytes.SplitAfterN(b, b, 2) */
	_ = bytes.SplitAfterN(b, b, 2)
	/*! suspicious duplicated args in bytes.SplitN(b, b, 2) */
	_ = bytes.SplitN(b, b, 2)
	/*! suspicious duplicated args in bytes.Replace(nil, b, b, 1) */
	_ = bytes.Replace(nil, b, b, 1)
	/*! suspicious duplicated args in bytes.Replace(nil, []byte("a"), []byte("a"), 1) */
	_ = bytes.Replace(nil, []byte("a"), []byte("a"), 1)

	/*! suspicious duplicated args in types.Identical(typ, typ) */
	_ = types.Identical(typ, typ)
	/*! suspicious duplicated args in types.IdenticalIgnoreTags(typ, typ) */
	_ = types.IdenticalIgnoreTags(typ, typ)

	{
		var area image.Rectangle
		var point image.Point
		var op draw.Op
		/*! suspicious duplicated args in draw.Draw(dstImg, area, dstImg, point, op) */
		draw.Draw(dstImg, area, dstImg, point, op)
	}

	var x float64
	/*! suspicious duplicated args in math.Max(x, x) */
	_ = math.Max(x, x)
	/*! suspicious duplicated args in math.Min(x, x) */
	_ = math.Min(x, x)
}
