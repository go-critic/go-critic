package checker_test

import (
	"path/filepath"
)

func badArgs() {
	filename := "file.go"
	dir := "testdata"

	/*! "testdata/" contains a path separator */
	_ = filepath.Join("testdata/", filename)

	/*! "/file.go" contains a path separator */
	_ = filepath.Join(dir, "/file.go")

	/*! `\a\b.txt` contains a path separator */
	/*! `testdata\` contains a path separator */
	_ = filepath.Join(`testdata\`, `\a\b.txt`)
}
