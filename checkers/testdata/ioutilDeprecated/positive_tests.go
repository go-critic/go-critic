package checker_test

import (
	"io"
	"io/ioutil"
)

func _(r io.Reader) {
	/*! ioutil.ReadAll is deprecated, use io.ReadAll instead */
	ioutil.ReadAll(r)
	/*! ioutil.ReadFile is deprecated, use os.ReadFile instead */
	ioutil.ReadFile("")
	/*! ioutil.WriteFile is deprecated, use os.WriteFile instead */
	ioutil.WriteFile("", nil, 0)
	/*! ioutil.ReadDir is deprecated, use os.ReadDir instead */
	ioutil.ReadDir("")
	/*! ioutil.NopCloser is deprecated, use io.NopCloser instead */
	ioutil.NopCloser(r)
	/*! ioutil.Discard is deprecated, use io.Discard instead */
	_ = ioutil.Discard
}
