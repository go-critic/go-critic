package main

import (
	"bytes"
	"errors"
	"fmt"
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

type checker struct {
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

var checkers []checker

func main() {
	tmpl, err := template.ParseFiles(templatesPath + "overview.md.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	pkgs, err := parser.ParseDir(&token.FileSet{}, checkersPath,
		func(inf os.FileInfo) bool {
			return strings.HasSuffix(inf.Name(), "_checker.go")
		}, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range lint.RuleList() {
		log.Printf("parsing %s\n", r.Name())

		f, ok := pkgs["lint"].Files[fmt.Sprintf("%s%s_checker.go", checkersPath, r.Name())]
		if !ok {

			log.Printf("File not found: %s%s_checker.go", checkersPath, r.Name())
			continue
		}
		c := checker{
			Name:            r.Name(),
			SyntaxOnly:      r.SyntaxOnly,
			Experimental:    r.Experimental,
			VeryOpinionated: r.VeryOpinionated,
		}
		for _, comment := range f.Comments {
			if strings.HasPrefix(comment.Text(), "!") {
				parseComment(comment.Text(), &c)
				break
			}
		}
		checkers = append(checkers, c)

	}
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, struct {
		Checkers []checker
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

func parseComment(text string, c *checker) {
	lines := strings.Split(text, "\n")
	ind := 0
	stages := []func(l []string, i *int, c *checker) error{
		parseShortDesc,
		parseDesc,
		parseBefore,
		parseAfter,
		parseNote,
	}
	for _, st := range stages {
		err := st(lines, &ind, c)
		if err != nil {
			log.Println(err)
		}
	}
}

func parseShortDesc(lines []string, ind *int, c *checker) error {
	c.ShortDescription = strings.TrimSpace(lines[0][1:]) + "\n\n"
	*ind += 2 // skip empty line
	return nil
}

func parseDesc(lines []string, ind *int, c *checker) error {
	if len(lines) <= *ind {
		return errors.New("parseDesc: no description provided")
	}
	if strings.HasPrefix(lines[*ind], "@") { // if no description
		return nil
	}
	for *ind < len(lines) && len(lines[*ind]) > 0 {
		c.Description += strings.TrimSpace(lines[*ind]) + "\n"
		*ind++
	}
	*ind++ //skip empty line
	return nil
}

func parseBefore(lines []string, ind *int, c *checker) error {
	if len(lines) <= *ind || strings.TrimSpace(lines[*ind]) != "@Before:" {
		return errors.New("parseBefore: no @Before: section found")
	}
	*ind++
	for *ind < len(lines) && len(lines[*ind]) > 0 {
		c.Before += strings.TrimSpace(lines[*ind]) + "\n"
		*ind++
	}
	*ind++ //skip empty line
	return nil
}

func parseAfter(lines []string, ind *int, c *checker) error {
	if len(lines) <= *ind || strings.TrimSpace(lines[*ind]) != "@After:" {
		return errors.New("parseAfter: no @After: section found")
	}
	*ind++
	for *ind < len(lines) && len(lines[*ind]) > 0 {
		c.After += strings.TrimSpace(lines[*ind]) + "\n"
		*ind++
	}
	*ind++ //skip empty line
	return nil
}

func parseNote(lines []string, ind *int, c *checker) error {
	if len(lines) <= *ind {
		return nil // No @Note: section
	}
	if strings.TrimSpace(lines[*ind]) != "@Note:" {
		return errors.New("parseNote: last section is not @Note")
	}
	*ind++
	for *ind < len(lines) && len(lines[*ind]) > 0 {
		c.Note += strings.TrimSpace(lines[*ind]) + "\n"
		*ind++
	}
	return nil
}
