package checker_tests

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
}
