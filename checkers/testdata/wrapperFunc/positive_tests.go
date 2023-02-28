package checker_test

import (
	"bytes"
	"image"
	"image/draw"
	"net/http"
	"strings"
	"sync"
	"unicode"
)

func f(s string, b []byte, i draw.Image, r image.Rectangle, p image.Point, o draw.Op) {
	var wg sync.WaitGroup
	/*! use WaitGroup.Done method in `wg.Add(-1)` */
	wg.Add(-1)

	var buf bytes.Buffer
	/*! use Buffer.Reset method in `buf.Truncate(0)` */
	buf.Truncate(0)

	/*! use strings.Split method in `strings.SplitN(s, ".", -1)` */
	strings.SplitN(s, ".", -1)

	/*! use strings.ToTitle method in `strings.Map(unicode.ToTitle, s)` */
	strings.Map(unicode.ToTitle, s)

	/*! use strings.ReplaceAll method in `strings.Replace(s, "a", "b", -1)` */
	strings.Replace(s, "a", "b", -1)

	/*! suggestion: strings.Contains(s, ":") */
	_ = strings.Index(s, ":") >= 0
	/*! suggestion: strings.Contains(s, ":") */
	_ = strings.Index(s, ":") != -1
	/*! suggestion: strings.ContainsAny(s, ":") */
	_ = strings.IndexAny(s, ":") >= 0
	/*! suggestion: strings.ContainsAny(s, ":") */
	_ = strings.IndexAny(s, ":") != -1
	/*! suggestion: strings.ContainsRune(s, ':') */
	_ = strings.IndexRune(s, ':') >= 0
	/*! suggestion: strings.ContainsRune(s, ':') */
	_ = strings.IndexRune(s, ':') != -1

	/*! use bytes.Split method in `bytes.SplitN(b, []byte("."), -1)` */
	bytes.SplitN(b, []byte("."), -1)
	/*! use bytes.ToUpper method in `bytes.Map(unicode.ToUpper, b)` */
	bytes.Map(unicode.ToUpper, b)
	/*! use bytes.ToLower method in `bytes.Map(unicode.ToLower, b)` */
	bytes.Map(unicode.ToLower, b)
	/*! use bytes.ToTitle method in `bytes.Map(unicode.ToTitle, b)` */
	bytes.Map(unicode.ToTitle, b)

	/*! use bytes.ReplaceAll method in `bytes.Replace(b, b, b, -1)` */
	bytes.Replace(b, b, b, -1)

	/*! suggestion: bytes.Contains(b, []byte(":")) */
	_ = bytes.Index(b, []byte(":")) >= 0
	/*! suggestion: bytes.Contains(b, []byte(":")) */
	_ = bytes.Index(b, []byte(":")) != -1
	/*! suggestion: bytes.ContainsAny(b, ":") */
	_ = bytes.IndexAny(b, ":") >= 0
	/*! suggestion: bytes.ContainsAny(b, ":") */
	_ = bytes.IndexAny(b, ":") != -1
	/*! suggestion: bytes.ContainsRune(b, ':') */
	_ = bytes.IndexRune(b, ':') >= 0
	/*! suggestion: bytes.ContainsRune(b, ':') */
	_ = bytes.IndexRune(b, ':') != -1

	/*! use http.NotFoundHandler method in `http.HandlerFunc(http.NotFound)` */
	_ = http.HandlerFunc(http.NotFound)

	/*! use draw.Draw method in `draw.DrawMask(i, r, i, p, nil, image.Point{}, o)` */
	draw.DrawMask(i, r, i, p, nil, image.Point{}, o)
}

func stringsCut(s, sep, host, port string) {
	{
		/*! suggestion: host, port, _ = strings.Cut(s, sep) */
		i := strings.Index(s, sep)
		host, port = s[:i], s[i+1:]
	}
	{
		/*! suggestion: host, port, _ = strings.Cut(s, sep) */
		i := strings.Index(s, sep)
		host = s[:i]
		port = s[i+1:]
	}
	{
		/*! suggestion: if host, port, ok = strings.Cut(s, sep); ok { ... } */
		if i := strings.Index(s, sep); i != -1 {
			host, port = s[:i], s[i+1:]
		}
		/*! suggestion: if host, port, ok = strings.Cut(s, sep); ok { ... } */
		if i := strings.Index(s, sep); i != -1 {
			host = s[:i]
			port = s[i+1:]
		}
		/*! suggestion: if host, port, ok = strings.Cut(s, sep); ok { ... } */
		if i := strings.Index(s, sep); i >= 0 {
			host, port = s[:i], s[i+1:]
		}
		/*! suggestion: if host, port, ok = strings.Cut(s, sep); ok { ... } */
		if i := strings.Index(s, sep); i >= 0 {
			host = s[:i]
			port = s[i+1:]
		}
	}
}
