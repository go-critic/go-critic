package checker_test

import "strings"

func makeString() string {
	return strings.Repeat("abc", 3)
}

func bad() {
	/*! consider replacing []rune("abc")[0] with utf8.DecodeRuneInString("abc") */
	_ = []rune("abc")[0]
}

func badFunc() {
	/*! consider replacing []rune(makeString())[0] with utf8.DecodeRuneInString(makeString()) */
	_ = []rune(makeString())[0]
}
