package checker_test

import (
	"fmt"
	"os"
)

func printFilesIndexes(files []*os.File) {
	for i := range files {
		fmt.Println(i)
	}
}

func closeNonPtrFiles(files []os.File) {
	for i := range files {
		fmt.Println(files[i].Name)
		files[i].Close()
	}
}
