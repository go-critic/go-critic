package checker_test

import (
	"os"
	"path/filepath"
)

func _(a, b, c string) {
	_ = 'a' + os.PathSeparator + 'b'

	filepath.Join(a, b)

	filepath.Join(a, b, c)
}
