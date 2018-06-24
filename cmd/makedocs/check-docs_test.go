package main

import (
	"testing"
)

const (
	testdataPath = "./testdata"
)

func parseFiles(path string) []checkerDoc {
	pkgs := getPkgs(path, "_test.go")
	for name, f := range pkgs["lint"].Files {
		d := checkerDoc{Name: name}
		parseComments(f, &d)
		docCheckers = append(docCheckers, d)
	}
	return docCheckers
}

func TestOutput(t *testing.T) {
	wantDoc := checkerDoc{
		ShortDescription: "This is short desc",
		Before:           "Single line for code",
		After:            "After issue is fixed",
	}

	testFiles := parseFiles(testdataPath)

	for _, file := range testFiles {
		if file.ShortDescription != wantDoc.ShortDescription {
			t.Errorf("%s: ShortDescription mismatch:\nhave: %q\nwant: %q",
				file.Name, file.ShortDescription, wantDoc.ShortDescription)
		}
		if file.Before != wantDoc.Before {
			t.Errorf("%s: @Before mismatch:\nhave: %q\nwant: %q",
				file.Name, file.Before, wantDoc.Before)
		}
		if file.After != wantDoc.After {
			t.Errorf("%s: @After mismatch:\nhave: %q\nwant: %q",
				file.Name, file.After, wantDoc.After)
		}
	}
}
