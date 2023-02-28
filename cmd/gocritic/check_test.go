package main

import (
	"testing"
)

func TestShortenLocation(t *testing.T) {
	testGopath := "/home/queen/go/"
	testGoroot := "/usr/lib/go/"
	tests := []struct {
		wd    string
		input string
		out   string
	}{
		{"", "/home/queen/go/file.go", "$GOPATH/file.go"},
		{"", "/home/queen/go-file.go", "/home/queen/go-file.go"},

		{"", "/usr/lib/go/file.go", "$GOROOT/file.go"},
		{"", "/usr/lib/go-file.go", "/usr/lib/go-file.go"},

		{"/home/queen/go/src/", "/home/queen/go/src/file.go", "./file.go"},
		{"/home/queen/", "/home/queen-src-file.go", "/home/queen-src-file.go"},
		{"/home/", "/home/queen/go/src/file.go", "$GOPATH/src/file.go"},

		{`C:\home\queen\go\src\`, `C:\home\queen\go\src\file.go`, "./file.go"},
	}

	l := &program{
		gopath: testGopath,
		goroot: testGoroot,
	}
	for _, test := range tests {
		l.workDir = test.wd
		have := l.shortenLocation(test.input)
		want := test.out
		if have != want {
			t.Errorf("shorten(%q):\nhave: %q\nwant: %q",
				test.input, have, want)
		}
	}
}
