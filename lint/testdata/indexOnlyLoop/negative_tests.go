package checker_test

func alreadyWithValue(files []*string) {
	f := func(file *string) {}

	for _, file := range files {
		f(file)
		f(file)
	}

	for _, file := range files {
		f(files[0])
		f(files[1])
		f(file)
	}
}

func blankIdent(files []*string) {
	for _ = range files {
		_ = files[0]
		_ = files[1]
	}
}

func printFilesIndexes(files []*string) {
	f := func(s int) {}

	for i := range files {
		f(i)
	}
}

func closeNonPtrFiles(files []string) {
	f := func(s string) {}

	for i := range files {
		println(files[i])
		f(files[i])
	}
}

func indexReuse(filesA []*string, filesB []*string) {
	f := func() {}

	for i := range filesA {
		if filesA[i] == filesB[i] {
			f()
		}
	}
}

func channelUse() {
	f := func(s string) {}

	c := make(chan string)
	for val := range c {
		f(val)
	}
}

func sliceUse() {
	f := func(s string) {}

	v := []string{}
	for _, val := range v[:] {
		f(val)
	}
}

func rangeOverString() {
	f := func(s rune) {}

	for _, ch := range "abcdef" {
		f(ch)
	}
}

func shadowed() {
	var xs []*string
	for k := range xs {
		var xs []int
		println(xs[k] + xs[k])
	}
}
