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
	templatesPath = docsPath + "templates/"
)

func main() {
	tmpl := template.Must(template.ParseFiles(templatesPath + "overview.md.tmpl"))
	buf := bytes.Buffer{}
	err := tmpl.Execute(&buf, struct {
		Rules []*lint.Rule
	}{
		Rules: lint.RuleList(),
	})
	if err != nil {
		log.Fatalf("render template: %v", err)
	}
	if err := ioutil.WriteFile(docsPath+"overview.md", buf.Bytes(), 0600); err != nil {
		log.Fatalf("write output file: %v", err)
	}
}
