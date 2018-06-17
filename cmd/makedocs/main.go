package main

import (
	"bytes"
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

type parseState uint

const (
	desciption parseState = 1 << iota
	before
	after
	note
)

func parseComment(text string, c *checker) {
	lines := strings.Split(text, "\n")
	s := desciption
	for _, l := range lines {
		if strings.HasPrefix(l, "!") {
			if c.ShortDescription != "" {
				log.Fatal("Parse error: duplicate description section")
			}
			c.ShortDescription = strings.TrimSpace(l[1:])
			c.Description += strings.TrimSpace(l[1:]) + "\n\n"
			s = desciption
			continue
		}

		// TODO: remove duplicated code
		if strings.HasPrefix(l, "Before:") {
			if c.Before != "" {
				log.Fatal("Duplicated 'Before:' section")
			}
			s = before
			continue
		}

		if strings.HasPrefix(l, "After:") {
			if c.After != "" {
				log.Fatal("Duplicated 'After:' section")
			}
			s = after
			continue
		}

		if strings.HasPrefix(l, "Note:") {
			if c.Note != "" {
				log.Fatal("Duplicated 'Note:' section")
			}
			s = note
			continue
		}

		switch s {
		case desciption:
			c.Description += l + "\n"
		case before:
			c.Before += l + "\n"
		case after:
			c.After += l + "\n"
		case note:
			c.Note += l + "\n"
		}
	}
}
