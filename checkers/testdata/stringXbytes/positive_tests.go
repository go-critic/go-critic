package checker_test

import "regexp"

func warnings() {
	var b []byte
	var b2 []byte
	var s string

	/*! can simplify `[]byte(s)` to `s` */
	copy(b, []byte(s))

	re := regexp.MustCompile(`\w+`)

	/*! suggestion: len(b) */
	_ = len(string(b))

	/*! suggestion: len(b) */
	_ = len(string(b)) == 0

	/*! suggestion: len(b) == 0 */
	_ = string(b) == ""

	/*! suggestion: len(b) != 0 */
	_ = string(b) != ""

	/*! suggestion: re.MatchString(s) */
	_ = re.Match([]byte(s))

	/*! suggestion: re.FindStringIndex(s) */
	_ = re.FindIndex([]byte(s))

	/*! suggestion: re.FindAllStringIndex(s, -1) */
	_ = re.FindAllIndex([]byte(s), -1)

	/*! suggestion: bytes.Equal(b, b2) */
	_ = string(b) == string(b2)

	/*! suggestion: !bytes.Equal(b, b2) */
	_ = string(b) != string(b2)
}
