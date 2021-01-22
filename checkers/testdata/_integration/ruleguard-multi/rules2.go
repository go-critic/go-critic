// +build ignore

package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

func osFilepath(m dsl.Matcher) {
	// path/filepath package forwards path separators so if
	// the file already uses filepath-related API it might be
	// a good idea to reduce the direct os package dependency.
	// In some cases it helps to remove the "os" package import completely.

	m.Match(`os.PathSeparator`).
		Where(m.File().Imports("path/filepath")).
		Suggest(`filepath.Separator`)

	m.Match(`os.PathListSeparator`).
		Where(m.File().Imports("path/filepath")).
		Suggest(`filepath.ListSeparator`)
}
