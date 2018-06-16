package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/go-critic/go-critic/lint"
)

const (
	docsPath      = "../../docs/"
	checkersPath  = docsPath + "checkers/"
	templatesPath = docsPath + "templates/"
)

type checker struct {
	Name         string
	Description  string
	Experimental bool
	Kind         string
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
			Name:         r.Name(),
			Experimental: r.Experimental(),
			Kind:         ruleKindString(r),
			Description:  desc,
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

func ruleKindString(r *lint.Rule) string {
	if r.SyntaxOnly() {
		return "syntax-only check (fast)"
	}
	return "type-aware check"
}
