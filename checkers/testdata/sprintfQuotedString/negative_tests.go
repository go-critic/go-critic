package checker_test

import "fmt"

func _(s string) {
	_ = fmt.Sprintf(`%q`, s)
	_ = fmt.Sprintf(`foo %q bar`, s)

	_ = fmt.Sprintf("%q", s)
	_ = fmt.Sprintf("foo %q bar", s)

	_ = fmt.Sprintf("%s", s)
}
