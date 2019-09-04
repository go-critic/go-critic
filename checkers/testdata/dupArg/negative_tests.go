package checker_test

import (
	"bytes"
	"go/types"
	"image"
	"image/draw"
	"reflect"
	"strings"
)

func differentArgs() {
	var dstSlice, srcSlice []int
	var s, s2 string
	var b, b2 []byte
	var dstRV, srcRV reflect.Value
	var typ, typ2 types.Type
	var dstImg, srcImg draw.Image

	copy(dstSlice, srcSlice)

	_ = reflect.Copy(dstRV, srcRV)
	_ = reflect.DeepEqual(s, s2)

	_ = strings.Contains(s, s2)
	_ = strings.Compare(s, s2)
	_ = strings.EqualFold(s, s2)
	_ = strings.HasPrefix(s, s2)
	_ = strings.HasSuffix(s, s2)
	_ = strings.Index(s, s2)
	_ = strings.LastIndex(s, s2)
	_ = strings.Split(s, s2)
	_ = strings.SplitAfter(s, s2)
	_ = strings.SplitAfterN(s, s2, 2)
	_ = strings.SplitN(s, s2, 2)
	_ = strings.Replace("", s, s2, 1)
	_ = strings.Replace("", "a", "b", 1)

	_ = bytes.Contains(b, b2)
	_ = bytes.Compare(b, b2)
	_ = bytes.Equal(b, b2)
	_ = bytes.EqualFold(b, b2)
	_ = bytes.HasPrefix(b, b2)
	_ = bytes.HasSuffix(b, b2)
	_ = bytes.LastIndex(b, b2)
	_ = bytes.Split(b, b2)
	_ = bytes.SplitAfter(b, b2)
	_ = bytes.SplitAfterN(b, b2, 2)
	_ = bytes.SplitN(b, b2, 2)
	_ = bytes.Replace(nil, b, b2, 1)
	_ = bytes.Replace(nil, []byte("a"), []byte("b"), 1)

	_ = types.Identical(typ, typ2)
	_ = types.IdenticalIgnoreTags(typ, typ2)

	// shadowing builtin "copy"
	copy := func(x int) int { return x }
	_ = copy(1)

	{
		var area image.Rectangle
		var point image.Point
		var op draw.Op
		draw.Draw(dstImg, area, srcImg, point, op)
	}
}
