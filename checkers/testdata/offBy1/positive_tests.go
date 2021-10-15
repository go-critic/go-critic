package checker_test

import (
	"bytes"
	"strings"
)

func sink(args ...interface{}) {}

func lenIndex(xs []int, ys []string) {
	/*! index expr always panics; maybe you wanted xs[len(xs)-1]? */
	_ = xs[len(xs)]
	/*! index expr always panics; maybe you wanted ys[len(ys)-1]? */
	_ = ys[len(ys)]
}

func indexSlicing() {
	var s string
	var b []byte

	{
		start := strings.Index(s, "/")
		/*! Index() can return -1; maybe you wanted to do s[start+1:] */
		_ = s[start:]
	}
	{
		start := bytes.Index(b, []byte("/"))
		/*! Index() can return -1; maybe you wanted to do b[start+1:] */
		_ = b[start:]
	}
	{
		/*! Index() can return -1; maybe you wanted to do Index()+1 */
		_ = s[strings.Index(s, "/"):]
	}
	{
		/*! Index() can return -1; maybe you wanted to do Index()+1 */
		_ = b[bytes.Index(b, []byte("/")):]
	}
	{
		start := strings.Index(s, "/")
		/*! Index() can return -1; maybe you wanted to do s[start+1:] */
		res := s[start:]
		sink(res)
	}
	{
		start := bytes.Index(b, []byte("/"))
		/*! Index() can return -1; maybe you wanted to do b[start+1:] */
		res := b[start:]
		sink(res)
	}

	{
		/*! Index() can return -1; maybe you wanted to do Index()+1 */
		sink(s[:strings.Index(s, "/")])
	}

	{
		/*! Index() can return -1; maybe you wanted to do Index()+1 */
		sink(b[:bytes.Index(b, []byte("/"))])
	}

	{
		end := strings.Index(s, "/")
		/*! Index() can return -1; maybe you wanted to do s[:end+1] */
		res := s[:end]
		sink(res)
	}

	{
		end := bytes.Index(b, []byte("/"))
		/*! Index() can return -1; maybe you wanted to do b[:end+1] */
		res := b[:end]
		sink(res)
	}

	{
		end := strings.Index(s, "/")
		/*! Index() can return -1; maybe you wanted to do s[:end+1] */
		_ = s[:end]
	}

	{
		end := bytes.Index(b, []byte("/"))
		/*! Index() can return -1; maybe you wanted to do b[:end+1] */
		_ = b[:end]
	}
}
