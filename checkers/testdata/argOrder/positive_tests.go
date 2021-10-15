package checker_test

import (
	"bytes"
	"strings"
)

func badArgOrder(s string, b []byte) {
	/*! "http://" and s arguments order looks reversed */
	_ = strings.HasPrefix("http://", s)

	/*! []byte("http://") and b arguments order looks reversed */
	_ = bytes.HasPrefix([]byte("http://"), b)
	/*! []byte{'h', 't', 't', 'p', ':', '/', '/'} and b arguments order looks reversed */
	_ = bytes.HasPrefix([]byte{'h', 't', 't', 'p', ':', '/', '/'}, b)

	/*! ":" and s arguments order looks reversed */
	_ = strings.Contains(":", s)
	/*! []byte(":") and b arguments order looks reversed */
	_ = bytes.Contains([]byte(":"), b)

	/*! ":" and s arguments order looks reversed */
	_ = strings.TrimPrefix(":", s)
	/*! []byte(":") and b arguments order looks reversed */
	_ = bytes.TrimPrefix([]byte(":"), b)

	/*! ":" and s arguments order looks reversed */
	_ = strings.TrimSuffix(":", s)
	/*! []byte(":") and b arguments order looks reversed */
	_ = bytes.TrimSuffix([]byte(":"), b)

	/*! "/" and s arguments order looks reversed */
	_ = strings.Split("/", s)
	/*! []byte("/") and b arguments order looks reversed */
	_ = bytes.Split([]byte("/"), b)
}

func argubleCases(s string) {
	// It's possible to use strings.Contains as a "check element presents in a set".
	// But that usage is somewhat rare and can be implemented via sorted []string
	// search or a map[string]bool/map[string]struct{}.

	/*! "uint uint8 uint16 uint32" and s arguments order looks reversed */
	_ = strings.Contains("uint uint8 uint16 uint32", s)

	// Code below removes "optional " prefix if s="optional ".
	// But this is not the most clear way to do it.
	// We accept this false positive as it can be fixed
	// by assigning a string literal to a variable, for example.

	/*! "optional foo bar" and s arguments order looks reversed */
	_ = strings.TrimPrefix("optional foo bar", s)
}
