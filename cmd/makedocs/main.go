package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/go-critic/go-critic/lint"
)

const (
	docsPath      = "../../docs/"
	templatesPath = docsPath + "templates/"
	checkersPath  = "../../lint/"
)

type checkerDoc struct {
	Name             string
	ShortDescription string
	Description      string
	Before           string
	After            string
	Note             string
	SyntaxOnly       bool
	Experimental     bool
	VeryOpinionated  bool
}

var checkers []checkerDoc

var docCheckers []checkerDoc

func main() {
	tmpl, err := template.ParseFiles(templatesPath + "overview.md.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	pkgs := getPkgs(checkersPath, "_checker.go")

	for _, r := range lint.RuleList() {
		log.Printf("parsing %s\n", r.Name())

		f, ok := pkgs["lint"].Files[fmt.Sprintf("%s%s_checker.go", checkersPath, r.Name())]
		if !ok {
			log.Printf("File not found: %s%s_checker.go", checkersPath, r.Name())
			continue
		}
		c := checkerDoc{
			Name:            r.Name(),
			SyntaxOnly:      r.SyntaxOnly,
			Experimental:    r.Experimental,
			VeryOpinionated: r.VeryOpinionated,
		}
		parseComments(f, &c)
		checkers = append(checkers, c)

	}
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, struct {
		Checkers []checkerDoc
	}{
		Checkers: checkers,
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(docsPath+"overview.md", buf.Bytes(), 0600); err != nil {
		log.Fatal(err)
	}
}

func parseComment(text string, c *checkerDoc) {
	lines := strings.Split(text, "\n")
	index := 0
	stages := []func(l []string, i *int, c *checkerDoc) error{
		parseShortDesc,
		parseDesc,
		parseBefore,
		parseAfter,
		parseNote,
	}
	for _, st := range stages {
		err := st(lines, &index, c)
		if err != nil {
			log.Println(err)
		}
	}
}

func parseShortDesc(lines []string, index *int, c *checkerDoc) error {
	c.ShortDescription = strings.TrimSpace(lines[0][1:])
	*index += 2 // skip empty line
	return nil
}

func parseDesc(lines []string, index *int, c *checkerDoc) error {
	if len(lines) <= *index {
		return errors.New("parseDesc: no description provided")
	}
	if strings.HasPrefix(lines[*index], "@") { // if no description
		return nil
	}
	for *index < len(lines) && len(lines[*index]) > 0 {
		c.Description += lines[*index]
		*index++
	}
	*index++ //skip empty line
	return nil
}

func parseBefore(lines []string, index *int, c *checkerDoc) error {
	if len(lines) <= *index || strings.TrimSpace(lines[*index]) != "@Before:" {
		return errors.New("parseBefore: no @Before: section found")
	}
	*index++
	for *index < len(lines) && len(lines[*index]) > 0 {
		c.Before += lines[*index]
		*index++
	}
	*index++ //skip empty line
	return nil
}

func parseAfter(lines []string, index *int, c *checkerDoc) error {
	if len(lines) <= *index || strings.TrimSpace(lines[*index]) != "@After:" {
		return errors.New("parseAfter: no @After: section found")
	}
	*index++
	for *index < len(lines) && len(lines[*index]) > 0 {
		c.After += lines[*index]
		*index++
	}
	*index++ //skip empty line
	return nil
}

func parseNote(lines []string, index *int, c *checkerDoc) error {
	if len(lines) <= *index {
		return nil // No @Note: section
	}
	if strings.TrimSpace(lines[*index]) != "@Note:" {
		return errors.New("parseNote: last section is not @Note")
	}
	*index++
	for *index < len(lines) && len(lines[*index]) > 0 {
		c.Note += lines[*index]
		*index++
	}
	return nil
}

func getPkgs(path string, suffix string) map[string]*ast.Package {
	filter := func(inf os.FileInfo) bool {
		return strings.HasSuffix(inf.Name(), suffix)
	}

	pkgs, err := parser.ParseDir(&token.FileSet{}, path,
		filter, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	return pkgs
}

func parseComments(file *ast.File, cd *checkerDoc) {
	for _, comment := range file.Comments {
		if strings.HasPrefix(comment.Text(), "!") {
			parseComment(comment.Text(), cd)
			break
		}
	}
}
