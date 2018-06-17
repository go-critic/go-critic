package main

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

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
	SyntaxOnly       bool
	Experimental     bool
	VeryOpinionated  bool
}

var checkers []checker

func main() {
	tmpl, err := template.ParseFiles(templatesPath + "overview.md.tmpl")
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
	emptyLine parseState = 1 << iota
	desciption
	before
	after
)

func parseComment(text string, c *checker) {
	lines := strings.Split(text, "\n")
	s := emptyLine
	for _, l := range lines {
		if strings.HasPrefix(l, "!") {
			if s != emptyLine || c.ShortDescription != "" {
				log.Fatal("Parse error: duplicate description section")
			}
			c.ShortDescription = strings.TrimSpace(l[1:])
			c.Description += strings.TrimSpace(l[1:]) + "\n"
			s = desciption
			continue
		}

		if strings.HasPrefix(l, "@before") {
			if s != emptyLine {
				log.Fatal("No empty line before @before")
			}
			s = before
			continue
		}

		if strings.HasPrefix(l, "@after") {
			if s != emptyLine {
				log.Fatal("No empty line before @after")
			}
			s = after
			continue
		}

		if len(l) < 2 { // string is empty
			if s == emptyLine {
				log.Fatal("Duplicate empty line")
			}
			s = emptyLine
			continue
		}

		switch s {
		case emptyLine:
			log.Fatal("No section tag before empty line")
		case desciption:
			c.Description += l + "\n"
		case before:
			c.Before += l + "\n"
		case after:
			c.After += l + "\n"
		}
	}
}
