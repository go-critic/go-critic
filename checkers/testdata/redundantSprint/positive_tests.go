package checker_test

import "fmt"

type withStringer struct{}

func (withStringer) String() string { return "" }

func _() {
	{
		var foo withStringer
		/*! use foo.String() instead */
		_ = fmt.Sprint(foo)
		/*! use foo.String() instead */
		_ = fmt.Sprintf("%s", foo)
		/*! use foo.String() instead */
		_ = fmt.Sprintf("%v", foo)
	}

	{
		var s string
		/*! s is already string */
		_ = fmt.Sprint(s)
		/*! s is already string */
		_ = fmt.Sprintf("%s", s)
		/*! s is already string */
		_ = fmt.Sprintf("%v", s)

		/*! "x" is already string */
		_ = fmt.Sprint("x")
		/*! "x" is already string */
		_ = fmt.Sprintf("%s", "x")
		/*! "x" is already string */
		_ = fmt.Sprintf("%v", "x")
	}
}
