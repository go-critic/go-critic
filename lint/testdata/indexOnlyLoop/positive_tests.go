package checker_test

func closeFile(f *string) {}

func closeFiles(files []*string) {
	/// i occurs more than once in the loop; consider using for _, value := range files
	for i := range files {
		if files[i] != nil {
			closeFile(files[i])
		}
	}
}

func sliceLoop(files []*string) {
	/// k occurs more than once in the loop; consider using for _, value := range files
	for k := range files[:] {
		if files[k] != nil {
			closeFile(files[k])
		}
	}
}

func nestedLoop(files []*string) {
	/// k occurs more than once in the loop; consider using for _, value := range files
	for k := range files[:] {
		if files[k] != nil {
			closeFile(files[k])
		}

		var xs []*int
		/// j occurs more than once in the loop; consider using for _, value := range xs
		for j := range xs {
			println(*xs[j] + *xs[j])
		}
	}
}
