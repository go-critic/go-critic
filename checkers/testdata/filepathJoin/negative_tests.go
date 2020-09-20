package checker_test

import (
	"path/filepath"
)

func goodArgs() {
	filename := "file.go"
	dir := "testdata"

	_ = filepath.Join("testdata", filename)

	_ = filepath.Join(dir, "file.go")

	_ = filepath.Join(`testdata`, `a`, `b.txt`)
}
