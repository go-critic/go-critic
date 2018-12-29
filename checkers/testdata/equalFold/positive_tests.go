package checker_test

import (
	"bytes"
	"strings"
)

func concat(x, y string) string {
	return x + y
}

func stringsToLower(x, y string) {
	/*! consider replacing with strings.EqualFold(x, y) */
	_ = strings.ToLower(x) == y

	/*! consider replacing with strings.EqualFold(x, y) */
	_ = strings.ToLower(x) == strings.ToLower(y)

	/*! consider replacing with strings.EqualFold(x, y) */
	_ = x == strings.ToLower(y)

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = strings.ToLower(x) != "y"

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = strings.ToLower(x) == strings.ToLower("y")

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = x == strings.ToLower("y")

	/*! consider replacing with strings.EqualFold(x, concat(y, "123")) */
	_ = strings.ToLower(x) == concat(y, "123")
}

func stringsToUpper(x, y string) {
	/*! consider replacing with strings.EqualFold(x, y) */
	_ = strings.ToUpper(x) == y

	/*! consider replacing with strings.EqualFold(x, y) */
	_ = strings.ToUpper(x) != strings.ToUpper(y)

	/*! consider replacing with strings.EqualFold(x, y) */
	_ = x != strings.ToUpper(y)

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = strings.ToUpper(x) == "y"

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = strings.ToUpper(x) == strings.ToUpper("y")

	/*! consider replacing with strings.EqualFold(x, "y") */
	_ = x == strings.ToUpper("y")
}

func bytesToLower(x, y []byte) {
	/*! consider replacing with bytes.EqualFold(x, y) */
	_ = bytes.Equal(bytes.ToLower(x), y)

	/*! consider replacing with bytes.EqualFold(x, y) */
	_ = bytes.Equal(bytes.ToLower(x), bytes.ToLower(y))

	/*! consider replacing with bytes.EqualFold(x, y) */
	_ = !bytes.Equal(x, bytes.ToLower(y))

	/*! consider replacing with bytes.EqualFold(x, []byte("y")) */
	_ = !bytes.Equal(bytes.ToLower(x), []byte("y"))

	/*! consider replacing with bytes.EqualFold(x, []byte("y")) */
	_ = bytes.Equal(bytes.ToLower(x), bytes.ToLower([]byte("y")))

	/*! consider replacing with bytes.EqualFold(x, []byte("y")) */
	_ = bytes.Equal(x, bytes.ToLower([]byte("y")))

	/*! consider replacing with bytes.EqualFold(x, append(y, 'a')) */
	_ = bytes.Equal(bytes.ToLower(x), append(y, 'a'))
}

func bytesToUpper(x, y []byte) {
	/*! consider replacing with bytes.EqualFold(x, y) */
	_ = bytes.Equal(bytes.ToUpper(x), y)

	/*! consider replacing with bytes.EqualFold(x, y) */
	_ = !bytes.Equal(bytes.ToUpper(x), bytes.ToUpper(y))

	/*! consider replacing with bytes.EqualFold(x, y) */
	_ = bytes.Equal(x, bytes.ToUpper(y))

	/*! consider replacing with bytes.EqualFold(x, []byte("y")) */
	_ = bytes.Equal(bytes.ToUpper(x), []byte("y"))

	/*! consider replacing with bytes.EqualFold(x, []byte("y")) */
	_ = bytes.Equal(bytes.ToUpper(x), bytes.ToUpper([]byte("y")))

	/*! consider replacing with bytes.EqualFold(x, []byte("y")) */
	_ = bytes.Equal(x, bytes.ToUpper([]byte("y")))
}
