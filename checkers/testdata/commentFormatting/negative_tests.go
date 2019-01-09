package checker_test

/*
multi-line comments
are ignored
*/

// Special kinds of comments are permitted:
//+build
//-foo

//-style comments

//directive: abc

//#ifdef foo
//#endif

//!something

//nolint

//line /foo/bar.go:10

//line /foo/bar/f-ad/a_d.go:13
//line /bar.go:14

//go:noinline
func f1() {
	//nolint

	//	code
	//	example
	//	leading tabs

	// comment with normal style

	// this comment has empty lines
	//
	//
	// inside it.
}
