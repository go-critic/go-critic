package checker_test

import "regexp"

func noWarnings() {
	var b []byte
	var s string

	copy(b, s)
}

func _() {
	copy := func(int) {}

	copy(1)

	var s string
	re := regexp.MustCompile(`\w+`)

	_ = re.MatchString(s)

	_ = re.FindStringIndex(s)

	_ = re.FindAllStringIndex(s, -1)
}
