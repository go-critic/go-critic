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

//go:noinline
func f1() {
	//	code
	//	example
	//	leading tabs

	// comment with normal style

	// this comment has empty lines
	//
	//
	// inside it.
}
