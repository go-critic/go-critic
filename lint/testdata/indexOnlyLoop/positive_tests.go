package checker_test

import (
	"os"
)

func closeFiles(files []*os.File) {
	/// key variable occurs more then once in the loop; consider using for _, value := range files
	for i := range files {
		if files[i] != nil {
			files[i].Close()
		}
	}
}
