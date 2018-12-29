package checker_test

import "strings"

func foo(x, y string) {
	/*! consider replacing with strings.EqualFold(x, y) */
	_ = strings.ToLower(x) == y

	/*! consider replacing with strings.EqualFold(x, y) */
	_ = strings.ToLower(x) == strings.ToLower(y)

	/*! consider replacing with strings.EqualFold(x, y) */
	_ = x == strings.ToLower(y)

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = strings.ToLower(x) == "y"

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = strings.ToLower(x) == strings.ToLower("y")

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = x == strings.ToLower("y")
}

func bar(x, y string) {
	/*! consider replacing with strings.EqualFold(x, y) */
	_ = strings.ToUpper(x) == y

	/*! consider replacing with strings.EqualFold(x, y) */
	_ = strings.ToUpper(x) == strings.ToUpper(y)

	/*! consider replacing with strings.EqualFold(x, y) */
	_ = x == strings.ToUpper(y)

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = strings.ToUpper(x) == "y"

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = strings.ToUpper(x) == strings.ToUpper("y")

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = x == strings.ToUpper("y")
}
