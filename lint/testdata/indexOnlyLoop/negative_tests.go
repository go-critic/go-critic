package checker_test

import (
	"fmt"
)

func printFilesIndexes(files []*string) {
	f := func(s int) {}

	for i := range files {
		f(i)
	}
}

func closeNonPtrFiles(files []string) {
	f := func(s string) {}

	for i := range files {
		fmt.Println(files[i])
		f(files[i])
	}
}

func indexReuse(filesA []*string, filesB []*string) {
	for i := range filesA {
		if filesA[i] == filesB[i] {
			fmt.Println("equal")
		}
	}
}
