package checker_test

import (
	"bytes"
	"strings"
)

func badArgOrder(s string, b []byte) {
	/*! probably meant `strings.HasPrefix(s, "http://")` */
	_ = strings.HasPrefix("http://", s)

	/*! probably meant `bytes.HasPrefix(b, []byte("http://"))` */
	_ = bytes.HasPrefix([]byte("http://"), b)
	/*! probably meant `bytes.HasPrefix(b, []byte{'h', 't', 't', 'p', ':', '/', '/'})` */
	_ = bytes.HasPrefix([]byte{'h', 't', 't', 'p', ':', '/', '/'}, b)

	/*! probably meant `strings.Contains(s, ":")` */
	_ = strings.Contains(":", s)
	/*! probably meant `bytes.Contains(b, []byte(":"))` */
	_ = bytes.Contains([]byte(":"), b)

	/*! probably meant `strings.TrimPrefix(s, ":")` */
	_ = strings.TrimPrefix(":", s)
	/*! probably meant `bytes.TrimPrefix(b, []byte(":"))` */
	_ = bytes.TrimPrefix([]byte(":"), b)

	/*! probably meant `strings.TrimSuffix(s, ":")` */
	_ = strings.TrimSuffix(":", s)
	/*! probably meant `bytes.TrimSuffix(b, []byte(":"))` */
	_ = bytes.TrimSuffix([]byte(":"), b)

	/*! probably meant `strings.Split(s, "/")` */
	_ = strings.Split("/", s)
	/*! probably meant `bytes.Split(b, []byte("/"))` */
	_ = bytes.Split([]byte("/"), b)
}

func argubleCases(s string) {
	// It's possible to use strings.Contains as a "check element presents in a set".
	// But that usage is somewhat rare and can be implemented via sorted []string
	// search or a map[string]bool/map[string]struct{}.

	/*! probably meant `strings.Contains(s, "uint uint8 uint16 uint32")` */
	_ = strings.Contains("uint uint8 uint16 uint32", s)

	// Code below removes "optional " prefix if s="optional ".
	// But this is not the most clear way to do it.
	// We accept this false positive as it can be fixed
	// by assigning a string literal to a variable, for example.

	/*! probably meant `strings.TrimPrefix(s, "optional foo bar")` */
	_ = strings.TrimPrefix("optional foo bar", s)
}
