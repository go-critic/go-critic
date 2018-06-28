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
		fmt.Println(files[i].Name())
		files[i].Close()
	}
}

func indexReuse(filesA []*os.File, filesB []*os.File) {
	for i := range filesA {
		if filesA[i] == filesB[i] {
			fmt.Println("equal")
		}
	}
}
