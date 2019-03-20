package checker_test

import (
	"bytes"
	"strings"
)

func badStringsReplace(s, from, to string) {
	/*! suspicious arg 0, probably meant -1 */
	_ = strings.Replace(s, from, to, 0)
}

func badStringsSplitN(s, sep string) {
	/*! suspicious arg 0, probably meant -1 */
	_ = strings.SplitN(s, sep, 0)
}

func badBytesReplace(s, from, to []byte) {
	/*! suspicious arg 0, probably meant -1 */
	_ = bytes.Replace(s, from, to, 0)
}

func badBytesSplitN(s, sep []byte) {
	/*! suspicious arg 0, probably meant -1 */
	_ = bytes.SplitN(s, sep, 0)
}

func badAppend(xs []int) {
	/*! no-op append call, probably missing arguments */
	_ = append(xs)
	/*! no-op append call, probably missing arguments */
	_ = append(xs[2:])
}
