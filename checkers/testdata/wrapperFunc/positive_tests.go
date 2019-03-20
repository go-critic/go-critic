package checker_test

import (
	"bytes"
	"net/http"
	"strings"
	"sync"
	"unicode"
)

func f(s string, b []byte) {
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

	/*! use http.NotFoundHandler method in `http.HandlerFunc(http.NotFound)` */
	_ = http.HandlerFunc(http.NotFound)
}
