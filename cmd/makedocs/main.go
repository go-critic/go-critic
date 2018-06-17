package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/go-critic/go-critic/lint"
)

const (
	docsPath      = "../../docs/"
	checkersPath  = docsPath + "checkers/"
	templatesPath = docsPath + "templates/"
)

type checker struct {
	Name             string
	ShortDescription string
	Description      string
	SyntaxOnly       bool
	Experimental     bool
	VeryOpinionated  bool
}

func main() {
	tmpl, err := template.ParseFiles(templatesPath + "overview.md.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	checkers := []checker{}
	for _, r := range lint.RuleList() {
		desc, err := getDesc(r.Name())
		if err != nil {
			log.Fatal(r.Name())
		}
		checkers = append(checkers, checker{
			Name:             r.Name(),
			SyntaxOnly:       r.SyntaxOnly,
			Experimental:     r.Experimental,
			VeryOpinionated:  r.VeryOpinionated,
			ShortDescription: shortDescription(desc),
			Description:      desc,
		})
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

func getDesc(name string) (string, error) {
	b, err := ioutil.ReadFile(checkersPath + name + ".md")
	return string(b), err
}

func shortDescription(desc string) string {
	// We expect every checker to have first doc line
	// to be a short summary, complete sentence.
	summary := strings.Split(desc, "\n")[0]
	summary = strings.TrimSpace(summary)
	return summary
}
