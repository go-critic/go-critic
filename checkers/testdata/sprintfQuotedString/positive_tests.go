package checker_test

import "fmt"

func _(s string) {
	/*! use %q instead of "%s" for quoted strings */
	_ = fmt.Sprintf(`"%s"`, s)
	/*! use %q instead of "%s" for quoted strings */
	_ = fmt.Sprintf(`foo "%s" bar`, s)

	/*! use %q instead of "%s" for quoted strings */
	_ = fmt.Sprintf("\"%s\"", s)
	/*! use %q instead of "%s" for quoted strings */
	_ = fmt.Sprintf("foo \"%s\" bar", s)
}
