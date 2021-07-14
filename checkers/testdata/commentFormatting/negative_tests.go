//nolint // reason
package checker_test

//nolint reason

/*
multi-line comments
are ignored
*/

// Special kinds of comments are permitted:
//+build
//-foo

//-style comments

//directive: abc

//go:generate abc

//#ifdef foo
//#endif

//!something

//nolint

//export myfunc
func myfunc() {
}

func example() {
	//go:noinline
}

//go-sumtype:decl Data
type Data struct {
}

//go:noinline
func f2() {
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

//////////// Vertical ////////////////
//++++++++++++ bars ++++++++++++++++++
//############# are ##################
//---------- permitted ---------------

// Comment in a comment is //OK
// path */*//with asterisk
