package checker_test

import (
	"bytes"
	"strings"
)

func makeSlice() []int {
	return []int{}
}

func goodLenIndex(xs []int, ys []string) {
	_ = xs[len(xs)-1]
	_ = ys[len(ys)-1]

	// Conservative with function call.
	// Might return different lengths for both calls.
	_ = makeSlice()[len(makeSlice())]

	var m map[int]int

	// Not an error. Doesn't panic.
	_ = m[len(m)]
}

func goodIndexSlicing() {
	var s string
	var b []byte

	{
		start := strings.Index(s, "/") + 1
		_ = s[start:]
	}
	{
		start := bytes.Index(b, []byte("/")) + 1
		_ = b[start:]
	}
	{
		start := strings.Index(s, "/")
		_ = s[start+1:]
	}
	{
		start := bytes.Index(b, []byte("/"))
		_ = b[start+1:]
	}
	{
		_ = s[strings.Index(s, "/")+1:]
	}
	{
		_ = b[bytes.Index(b, []byte("/"))+1:]
	}
	{
		start := strings.Index(s, "/") + 1
		res := s[start:]
		sink(res)
	}
	{
		start := bytes.Index(b, []byte("/")) + 1
		res := b[start:]
		sink(res)
	}
	{
		start := strings.Index(s, "/")
		res := s[start+1:]
		sink(res)
	}
	{
		start := bytes.Index(b, []byte("/"))
		res := b[start+1:]
		sink(res)
	}

	{
		sink(s[:strings.Index(s, "/")+1])
		sink(s[:1+strings.Index(s, "/")])
	}

	{
		sink(b[:bytes.Index(b, []byte("/"))+1])
		sink(b[:1+bytes.Index(b, []byte("/"))])
	}
}
