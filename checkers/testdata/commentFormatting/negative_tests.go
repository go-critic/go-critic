//nolint // reason
package checker_test

//nolint reason

//noinspection this is ignored
//noinspection ALL

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

//go:noinline
func example() {
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

//region GoLand custom folding region
func goland() {
}
//endregion

//<editor-fold desc="VSCode custom folding region">
func vscode() {
}
//</editor-fold>
