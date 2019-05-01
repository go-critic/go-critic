package checker_test

func ExampleFoo() {
	// Output:
	// blah
	// blah
}

func permittedComments() {
	// TODO: foo(1+2)

	// reflect.DeepEqual

	// a+b

	// x

	// 1

	// https://foo.org/a/b/c/resouce#anchor

	// http://bar.baz/a/b/c/resouce#anchor

	// quote: "This is quote from some source"

	// method (e.g. foo.String())

	// type foo bar

	// type is int

	// not code at all

	// multi line commentary text
	// example that should not trigger the warning.

	// type myInt int

	// <-ch

	// functionCall (1, 2)

	// fooooooooooo (())

	// funcfuncfunc0_1 ()

	// FcallIntString  (1, "123")

	// notAfunccall (this is not a function call)

	// funccall () with period .

	// pretty-print

	// Output: "64 bytes or fewer"

	// Result: int *int

	// golang.org/issue/5290

	// testRead | testWriteTo | testRemaining

	// WriteHeader(hdr) == wantErr

	// Size: SizeOf(uint8) + SizeOf(uint32)

	// Flags: ModTime

	// 0 <= n <= len(b.buf)

	// b.r == 0 && b.w == 0

	// ignored //line directives

	// invalid //line directives with one colon

	// //line directives with omitted filenames lead to empty filenames

	/* CINC/CINV/CNEG */
}
