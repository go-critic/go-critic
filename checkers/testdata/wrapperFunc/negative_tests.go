package checker_test

import (
	"bytes"
	"image"
	"image/draw"
	"net/http"
	"strings"
	"sync"
)

func appliedSuggestions(s string, b []byte, i draw.Image, r image.Rectangle, p image.Point, o draw.Op) {
	var wg sync.WaitGroup
	wg.Done()
	wg.Add(1)

	var buf bytes.Buffer
	buf.Reset()
	buf.Truncate(1)

	strings.Split(s, ".")
	strings.ToTitle(s)
	strings.ReplaceAll(s, "a", "b")

	bytes.Split(b, []byte("."))
	bytes.ToUpper(b)
	bytes.ToLower(b)
	bytes.ToTitle(b)

	bytes.ReplaceAll(b, b, b)

	_ = http.NotFoundHandler()

	draw.Draw(i, r, i, p, o)
}

func nonMatchingArgs(s string, b []byte, i draw.Image, r image.Rectangle, p image.Point, o draw.Op) {
	var wg sync.WaitGroup
	wg.Add(1)

	strings.Map(nil, s)

	strings.Replace(s, "a", "b", 1)

	strings.Index(s, ":")
	strings.IndexAny(s, ":")
	strings.IndexRune(s, ':')

	bytes.Map(nil, b)
	bytes.Map(nil, b)
	bytes.Map(nil, b)

	bytes.Replace(b, b, b, 1)

	bytes.Index(b, []byte(":"))
	bytes.IndexAny(b, ":")
	bytes.IndexRune(b, ':')

	_ = http.HandlerFunc(nil)

	draw.DrawMask(i, r, i, p, i, image.Point{}, o)
}
