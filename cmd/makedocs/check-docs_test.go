package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

var docCheckers []checkerDoc

func parseFiles() []checkerDoc {
	tests, err := ioutil.ReadDir("./testdata")

	if err != nil {
		log.Fatal(err)
	}

	pkgs, err := parser.ParseDir(&token.FileSet{}, "./testdata",
		func(inf os.FileInfo) bool {
			return strings.HasSuffix(inf.Name(), ".go")
		}, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for _, r := range pkgs {
		for _, f := range r.Files {
			d := checkerDoc{Name: tests[i].Name()}
			for _, comment := range f.Comments {
				if strings.HasPrefix(comment.Text(), "!") {
					parseComment(comment.Text(), &d)
					break
				}
			}
			docCheckers = append(docCheckers, d)
			i++
		}
	}
	return docCheckers
}

func TestOutput(t *testing.T) {
	wantDoc := checkerDoc{
		ShortDescription: "This is short desc",
		Before:           "Single line for code",
		After:            "After issue is fixed",
	}

	testFiles := parseFiles()

	for _, file := range testFiles {
		if file.ShortDescription == wantDoc.ShortDescription {
			continue
		} else {
			t.Errorf("%s: ShortDescription mismatch:\nhave: %q\nwant: %q", file.Name,
				file.ShortDescription, wantDoc.ShortDescription)
		}
		if file.Before == wantDoc.Before {
			continue
		} else {
			t.Errorf("%s: @Before mismatch:\nhave: %q\nwant: %q", file.Name,
				file.Before, wantDoc.Before)
		}
		if file.After == wantDoc.After {
			continue
		} else {
			t.Errorf("%s: @After mismatch:\nhave: %q\nwant: %q", file.Name,
				file.After, wantDoc.After)
		}
	}
}
